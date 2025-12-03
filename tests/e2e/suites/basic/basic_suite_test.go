package basic

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/giantswarm/apptest-framework/v2/pkg/state"
	"github.com/giantswarm/apptest-framework/v2/pkg/suite"
	"github.com/giantswarm/clustertest/v2/pkg/application"
	"github.com/giantswarm/clustertest/v2/pkg/logger"
)

const (
	isUpgrade = false
)

func TestBasic(t *testing.T) {
	const (
		appReadyTimeout  = 10 * time.Minute
		appReadyInterval = 5 * time.Second
	)
	suite.New().
		WithInstallNamespace("kube-system").
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		AfterClusterReady(func() {
			// no
		}).
		BeforeUpgrade(func() {
			// E.g. ensure that the initial install has completed and has settled before upgrading
		}).
		Tests(func() {
			It("should create DNS records", func() {
				baseDomain := getWorkloadClusterBaseDomain()
				testDomain := fmt.Sprintf("test.%s", baseDomain)
				testTxtDomain := fmt.Sprintf("%scname-%s", state.GetCluster().Name, testDomain)
				resolver := getResolver(baseDomain)

				By("adding a LoadBalancer service")
				Eventually(func() (bool, error) {
					wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
					if err != nil {
						return false, err
					}

					service := &corev1.Service{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "external-dns-test",
							Namespace: "kube-system",
							Annotations: map[string]string{
								"external-dns.alpha.kubernetes.io/hostname": testDomain,
								"giantswarm.io/external-dns":                "managed",
							},
						},
						Spec: corev1.ServiceSpec{
							Type: corev1.ServiceTypeLoadBalancer,
							Ports: []corev1.ServicePort{
								{
									Port:       80,
									TargetPort: intstr.FromInt(80),
									Protocol:   corev1.ProtocolTCP,
								},
							},
							Selector: map[string]string{
								"app": "external-dns-test",
							},
						},
					}

					err = wcClient.Create(context.Background(), service)
					if err != nil {
						return false, err
					}

					logger.Log("created service %s", service.Name)
					return true, nil
				}).
					WithTimeout(5 * time.Minute).
					WithPolling(15 * time.Second).
					Should(BeTrue())

				By("checking the DNS record exists")

				Eventually(func() (bool, error) {
					logger.Log("checking the DNS record exists for %s", testDomain)
					ips, err := resolver.LookupHost(context.Background(), testDomain)
					if err != nil {
						return false, err
					}

					if len(ips) == 0 {
						logger.Log("no DNS records found for %s", testDomain)
						return false, fmt.Errorf("no DNS records found for %s", testDomain)

					}
					logger.Log("DNS records found for %s: %v", testDomain, ips)
					return true, nil
				}).
					WithTimeout(5 * time.Minute).
					WithPolling(15 * time.Second).
					Should(BeTrue())

				By("checking that the TXT registry record exists")

				Eventually(func() (bool, error) {
					logger.Log("checking the TXT registry record exists for %s", testTxtDomain)
					txtRecords, err := resolver.LookupTXT(context.Background(), testTxtDomain)
					if err != nil {
						return false, err
					}

					if len(txtRecords) == 0 {
						logger.Log("no TXT records found for %s", testTxtDomain)
						return false, fmt.Errorf("no TXT records found for %s", testTxtDomain)

					}

					expectedValue := "heritage=external-dns,external-dns/owner=giantswarm-io-external-dns,external-dns/resource=service/kube-system/external-dns-test"
					if txtRecords[0] != expectedValue {
						logger.Log("TXT record found for %s: %s", testTxtDomain, txtRecords[0])
						return false, fmt.Errorf("TXT record found for %s: %s", testTxtDomain, txtRecords[0])
					}

					logger.Log("TXT records found for %s: %v", testTxtDomain, txtRecords)
					return true, nil
				}).
					WithTimeout(5 * time.Minute).
					WithPolling(15 * time.Second).
					Should(BeTrue())
			})
		}).
		Run(t, "Basic Test")
}

func getWorkloadClusterBaseDomain() string {
	values := &application.ClusterValues{}
	err := state.GetFramework().MC().GetHelmValues(state.GetCluster().Name, state.GetCluster().GetNamespace(), values)
	Expect(err).NotTo(HaveOccurred())

	if values.BaseDomain == "" {
		Fail("baseDomain field missing from cluster helm values")
	}

	return fmt.Sprintf("%s.%s", state.GetCluster().Name, values.BaseDomain)
}

func getResolver(baseDomain string) *net.Resolver {
	// Get the NS record for the base domain
	nsRecords, err := net.LookupNS(baseDomain)
	Expect(err).NotTo(HaveOccurred())
	Expect(nsRecords).NotTo(BeEmpty())

	// Get the IP of the first nameserver
	nsHost := nsRecords[0].Host
	nsIPs, err := net.LookupHost(nsHost)
	Expect(err).NotTo(HaveOccurred())
	Expect(nsIPs).NotTo(BeEmpty())

	nsIP := nsIPs[0]
	logger.Log("Using nameserver %s (%s) for DNS queries", nsHost, nsIP)

	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 10 * time.Second,
			}
			return d.DialContext(ctx, "udp", nsIP+":53")
		},
	}
}
