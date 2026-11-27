package namespacedgateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	cr "sigs.k8s.io/controller-runtime/pkg/client"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/apptest-framework/v5/pkg/state"
	"github.com/giantswarm/apptest-framework/v5/pkg/suite"
	"github.com/giantswarm/clustertest/v5/pkg/client"
	"github.com/giantswarm/clustertest/v5/pkg/logger"

	"e2e/internal/dnstest"
)

const (
	isUpgrade = false

	namespace = "kube-system"

	// Gateway API release whose standard CRDs are installed when the cluster has no
	// ListenerSet CRD. It matches the Gateway API version ExternalDNS 0.23 is built with.
	gatewayAPIVersion = "v1.6.2"

	// No Gateway controller runs in the test cluster. The test writes the statuses a
	// controller would write itself, under this controller name.
	fakeControllerName = "giantswarm.io/external-dns-e2e"

	// ExternalDNS takes the targets of a Gateway from this annotation, so the test
	// needs no load balancer. 192.0.2.0/24 is reserved for documentation (RFC 5737).
	targetIP = "192.0.2.10"

	managedAnnotation = "giantswarm.io/external-dns"
	targetAnnotation  = "external-dns.alpha.kubernetes.io/target"

	gatewayName     = "external-dns-e2e"
	listenerSetName = "external-dns-e2e"
	listenerName    = "ls"
)

func TestNamespacedGateway(t *testing.T) {
	suite.New().
		WithInstallNamespace(namespace).
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		AfterClusterReady(func() {
			It("has the Gateway API CRDs with ListenerSet", func() {
				ensureGatewayAPICRDs()
			})
		}).
		Tests(func() {
			It("creates cluster-scoped RBAC only for nodes, namespaces, gateways and listenersets", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).NotTo(HaveOccurred())

				clusterRoles := &rbacv1.ClusterRoleList{}
				Expect(wcClient.List(state.GetContext(), clusterRoles, cr.MatchingLabels{"app.kubernetes.io/name": "external-dns"})).To(Succeed())

				rulesBySuffix := map[string][]rbacv1.PolicyRule{}
				for _, role := range clusterRoles.Items {
					for _, suffix := range []string{"-nodes", "-namespaces", "-gateways", "-listenersets"} {
						if strings.HasSuffix(role.Name, suffix) {
							rulesBySuffix[suffix] = role.Rules
						}
					}
				}
				readOnly := []string{"get", "watch", "list"}
				Expect(rulesBySuffix).To(HaveKeyWithValue("-nodes", ConsistOf(rbacv1.PolicyRule{APIGroups: []string{""}, Resources: []string{"nodes"}, Verbs: []string{"list", "watch"}})))
				Expect(rulesBySuffix).To(HaveKeyWithValue("-namespaces", ConsistOf(rbacv1.PolicyRule{APIGroups: []string{""}, Resources: []string{"namespaces"}, Verbs: readOnly})))
				Expect(rulesBySuffix).To(HaveKeyWithValue("-gateways", ConsistOf(rbacv1.PolicyRule{APIGroups: []string{gatewayv1.GroupName}, Resources: []string{"gateways"}, Verbs: readOnly})))
				Expect(rulesBySuffix).To(HaveKeyWithValue("-listenersets", ConsistOf(rbacv1.PolicyRule{APIGroups: []string{gatewayv1.GroupName}, Resources: []string{"listenersets"}, Verbs: readOnly})))
			})

			It("creates DNS records for HTTPRoutes attached to a Gateway and a ListenerSet", func() {
				ctx := state.GetContext()
				baseDomain := dnstest.WorkloadClusterBaseDomain()
				gatewayHost := fmt.Sprintf("gw.%s", baseDomain)
				listenerSetHost := fmt.Sprintf("ls.%s", baseDomain)
				resolver := dnstest.Resolver(baseDomain)

				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).NotTo(HaveOccurred())
				DeferCleanup(func() {
					if CurrentSpecReport().Failed() {
						logExternalDNSPods(wcClient)
					}
				})

				By("restarting ExternalDNS so it picks up Gateway API CRDs installed after it started")
				Expect(wcClient.DeleteAllOf(ctx, &corev1.Pod{}, cr.InNamespace(namespace), cr.MatchingLabels{"app.kubernetes.io/name": "external-dns"})).To(Succeed())

				By("creating a Gateway that allows ListenerSets from its namespace")
				fromSame := gatewayv1.NamespacesFromSame
				gateway := &gatewayv1.Gateway{
					ObjectMeta: metav1.ObjectMeta{
						Name:        gatewayName,
						Namespace:   namespace,
						Annotations: map[string]string{managedAnnotation: "managed", targetAnnotation: targetIP},
					},
					Spec: gatewayv1.GatewaySpec{
						GatewayClassName: "external-dns-e2e",
						AllowedListeners: &gatewayv1.AllowedListeners{
							Namespaces: &gatewayv1.ListenerNamespaces{From: &fromSame},
						},
						Listeners: []gatewayv1.Listener{{
							Name:     "http",
							Hostname: new(gatewayv1.Hostname(gatewayHost)),
							Port:     80,
							Protocol: gatewayv1.HTTPProtocolType,
						}},
					},
				}
				Expect(createIfMissing(ctx, wcClient, gateway)).To(Succeed())

				By("creating an accepted ListenerSet attached to the Gateway")
				listenerSet := &gatewayv1.ListenerSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:        listenerSetName,
						Namespace:   namespace,
						Annotations: map[string]string{managedAnnotation: "managed"},
					},
					Spec: gatewayv1.ListenerSetSpec{
						ParentRef: gatewayv1.ParentGatewayReference{Name: gatewayName},
						Listeners: []gatewayv1.ListenerEntry{{
							Name:     listenerName,
							Hostname: new(gatewayv1.Hostname(listenerSetHost)),
							Port:     80,
							Protocol: gatewayv1.HTTPProtocolType,
						}},
					},
				}
				Expect(createIfMissing(ctx, wcClient, listenerSet)).To(Succeed())
				Eventually(func() error {
					if err := wcClient.Get(ctx, cr.ObjectKeyFromObject(listenerSet), listenerSet); err != nil {
						return err
					}
					listenerSet.Status.Conditions = []metav1.Condition{acceptedCondition(listenerSet.Generation)}
					listenerSet.Status.Listeners = []gatewayv1.ListenerEntryStatus{{
						Name:           listenerName,
						AttachedRoutes: 1,
						SupportedKinds: []gatewayv1.RouteGroupKind{{Kind: "HTTPRoute"}},
						Conditions:     []metav1.Condition{acceptedCondition(listenerSet.Generation)},
					}}
					return wcClient.Status().Update(ctx, listenerSet)
				}).WithTimeout(2 * time.Minute).WithPolling(10 * time.Second).Should(Succeed())

				By("creating HTTPRoutes accepted by the Gateway and by the ListenerSet")
				gatewayRoute := acceptedHTTPRoute(ctx, wcClient, "external-dns-e2e-gateway", gatewayHost, gatewayv1.ParentReference{
					Name: gatewayName,
				})
				listenerSetRoute := acceptedHTTPRoute(ctx, wcClient, "external-dns-e2e-listenerset", listenerSetHost, gatewayv1.ParentReference{
					Kind:        new(gatewayv1.Kind("ListenerSet")),
					Name:        listenerSetName,
					SectionName: new(gatewayv1.SectionName(listenerName)),
				})

				for _, host := range []string{gatewayHost, listenerSetHost} {
					By(fmt.Sprintf("checking the DNS record for %s", host))
					Eventually(func() ([]string, error) {
						logger.Log("looking up %s", host)
						return resolver.LookupHost(ctx, host)
					}).
						WithTimeout(10*time.Minute).
						WithPolling(15*time.Second).
						Should(ConsistOf(targetIP), "ExternalDNS did not publish %s; check its logs for forbidden errors", host)
				}

				By("cleaning up the Gateway API objects")
				for _, obj := range []cr.Object{gatewayRoute, listenerSetRoute, listenerSet, gateway} {
					Expect(cr.IgnoreNotFound(wcClient.Delete(ctx, obj))).To(Succeed())
				}
			})
		}).
		Run(t, "Namespaced Gateway Test")
}

// logExternalDNSPods logs the ExternalDNS pods' recent output, so a failed run shows why
// no record was published (for example a forbidden error) after the cluster is gone.
func logExternalDNSPods(wcClient *client.Client) {
	pods := &corev1.PodList{}
	if err := wcClient.List(state.GetContext(), pods, cr.InNamespace(namespace), cr.MatchingLabels{"app.kubernetes.io/name": "external-dns"}); err != nil {
		logger.Log("listing ExternalDNS pods: %v", err)
		return
	}
	for i := range pods.Items {
		logs, err := wcClient.GetLogs(state.GetContext(), &pods.Items[i], new(int64(200)))
		if err != nil {
			logger.Log("getting logs of %s: %v", pods.Items[i].Name, err)
			continue
		}
		logger.Log("logs of ExternalDNS pod %s:\n%s", pods.Items[i].Name, logs)
	}
}

// ensureGatewayAPICRDs installs the standard Gateway API CRDs unless the cluster already
// serves ListenerSets at v1, and waits until they are established.
func ensureGatewayAPICRDs() {
	ctx := state.GetContext()
	wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
	Expect(err).NotTo(HaveOccurred())
	// The clustertest client's scheme has no apiextensions types.
	Expect(apiextensionsv1.AddToScheme(wcClient.Scheme())).To(Succeed())

	crdKey := cr.ObjectKey{Name: "listenersets.gateway.networking.k8s.io"}
	crd := &apiextensionsv1.CustomResourceDefinition{}
	err = wcClient.Get(ctx, crdKey, crd)
	switch {
	case err == nil:
		logger.Log("Gateway API CRDs already installed, not installing %s", gatewayAPIVersion)
	case apierrors.IsNotFound(err):
		logger.Log("installing Gateway API %s standard CRDs", gatewayAPIVersion)
		for _, obj := range gatewayAPIManifests() {
			Expect(createIfMissing(ctx, wcClient, obj)).To(Succeed(), "creating %s %s", obj.GetKind(), obj.GetName())
		}
	default:
		Expect(err).NotTo(HaveOccurred())
	}

	Eventually(func() error {
		if err := wcClient.Get(ctx, crdKey, crd); err != nil {
			return err
		}
		served := false
		for _, v := range crd.Spec.Versions {
			served = served || (v.Name == "v1" && v.Served)
		}
		if !served {
			return fmt.Errorf("%s does not serve v1, which ExternalDNS 0.23 needs", crd.Name)
		}
		for _, c := range crd.Status.Conditions {
			if c.Type == apiextensionsv1.Established && c.Status == apiextensionsv1.ConditionTrue {
				return nil
			}
		}
		return fmt.Errorf("%s is not established yet", crd.Name)
	}).WithTimeout(2 * time.Minute).WithPolling(5 * time.Second).Should(Succeed())
}

func gatewayAPIManifests() []*unstructured.Unstructured {
	url := fmt.Sprintf("https://github.com/kubernetes-sigs/gateway-api/releases/download/%s/standard-install.yaml", gatewayAPIVersion)
	resp, err := http.Get(url) //nolint:gosec,noctx // fixed release URL
	Expect(err).NotTo(HaveOccurred())
	defer resp.Body.Close() //nolint:errcheck
	Expect(resp.StatusCode).To(Equal(http.StatusOK), "downloading %s", url)
	body, err := io.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())

	var objs []*unstructured.Unstructured
	for doc := range strings.SplitSeq(string(body), "\n---") {
		obj := &unstructured.Unstructured{}
		Expect(yaml.Unmarshal([]byte(doc), &obj.Object)).To(Succeed())
		if len(obj.Object) == 0 {
			continue
		}
		objs = append(objs, obj)
	}
	Expect(objs).NotTo(BeEmpty())
	return objs
}

// acceptedHTTPRoute creates an HTTPRoute for host with the given parent and marks it
// accepted by that parent, as a Gateway controller would.
func acceptedHTTPRoute(ctx context.Context, c cr.Client, name, host string, parent gatewayv1.ParentReference) *gatewayv1.HTTPRoute {
	route := &gatewayv1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: map[string]string{managedAnnotation: "managed"},
		},
		Spec: gatewayv1.HTTPRouteSpec{
			CommonRouteSpec: gatewayv1.CommonRouteSpec{ParentRefs: []gatewayv1.ParentReference{parent}},
			Hostnames:       []gatewayv1.Hostname{gatewayv1.Hostname(host)},
		},
	}
	Expect(createIfMissing(ctx, c, route)).To(Succeed())

	Eventually(func() error {
		if err := c.Get(ctx, cr.ObjectKeyFromObject(route), route); err != nil {
			return err
		}
		route.Status.Parents = []gatewayv1.RouteParentStatus{{
			ParentRef:      parent,
			ControllerName: fakeControllerName,
			Conditions:     []metav1.Condition{acceptedCondition(route.Generation)},
		}}
		return c.Status().Update(ctx, route)
	}).WithTimeout(2 * time.Minute).WithPolling(10 * time.Second).Should(Succeed())

	return route
}

func acceptedCondition(generation int64) metav1.Condition {
	return metav1.Condition{
		Type:               "Accepted",
		Status:             metav1.ConditionTrue,
		Reason:             "Accepted",
		Message:            "Accepted by the ExternalDNS e2e test",
		ObservedGeneration: generation,
		LastTransitionTime: metav1.Now(),
	}
}

func createIfMissing(ctx context.Context, c cr.Client, obj cr.Object) error {
	err := c.Create(ctx, obj)
	if apierrors.IsAlreadyExists(err) {
		return c.Get(ctx, cr.ObjectKeyFromObject(obj), obj)
	}
	return err
}
