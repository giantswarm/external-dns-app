# external-dns-app

Configure external DNS servers for Kubernetes Ingresses and Services

![Version: 3.6.0](https://img.shields.io/badge/Version-3.6.0-informational?style=flat-square)

![AppVersion: 0.22.0](https://img.shields.io/badge/AppVersion-0.22.0-informational?style=flat-square)

**Homepage:** <https://github.com/giantswarm/external-dns-app>

## Requirements

Kubernetes: `>=1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
| oci://gsoci.azurecr.io/charts/giantswarm | kubectl-apply-job | 0.13.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.imagePullSecrets | list | `[]` | Global image pull secrets. |
| image.repository | string | `"gsoci.azurecr.io/giantswarm/external-dns"` | Image repository for the `external-dns` container. |
| image.tag | string | `nil` | Image tag for the `external-dns` container, this will default to `.Chart.AppVersion` if not set. |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for the `external-dns` container. |
| imagePullSecrets | list | `[]` | Image pull secrets. |
| nameOverride | string | `"external-dns"` | Override the name of the chart. |
| fullnameOverride | string | `nil` | Override the full name of the chart. |
| namespaceOverride | string | `nil` | Override the namespace that chart resources are rendered into. Defaults to the release namespace. Useful when installing the chart as a subchart that should live in its own namespace, separate from the umbrella release namespace. |
| commonLabels | object | `{}` | Labels to add to all chart resources. |
| serviceAccount.create | bool | `true` | If `true`, create a new `ServiceAccount`. |
| serviceAccount.labels | object | `{}` | Labels to add to the service account. |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account. Templates are allowed in both the key and the value. Example: `example.com/annotation/{{ .Values.nameOverride }}: {{ .Values.nameOverride }}` |
| serviceAccount.name | string | `nil` | If this is set and `serviceAccount.create` is `true` this will be used for the created `ServiceAccount` name, if set and `serviceAccount.create` is `false` then this will define an existing `ServiceAccount` to use. |
| serviceAccount.automountServiceAccountToken | bool | `true` | Set this to `false` to [opt out of API credential automounting](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the `ServiceAccount`. |
| service.annotations | object | `{}` | Service annotations. |
| service.port | int | `7979` | Service HTTP port. |
| service.ipFamilies | list | `[]` | Service IP families (e.g. IPv4 and/or IPv6). |
| service.ipFamilyPolicy | string | `nil` | Service IP family policy. |
| service.enabled | bool | `true` | If `true`, create a `Service` Kubernetes. |
| rbac.create | bool | `true` | If `true`, create a `ClusterRole` & `ClusterRoleBinding` with access to the Kubernetes API. |
| rbac.additionalPermissions | list | `[]` | Additional rules to add to the `ClusterRole`. |
| replicaCount | int | `1` | Number of replicas of the `external-dns` `Deployment`. external-dns does not support leader election, so this must be `0` or `1` to avoid duplicate or conflicting DNS record updates. Set to `0` to scale the `Deployment` down. |
| deploymentAnnotations | object | `{}` | Annotations to add to the `Deployment`. |
| extraContainers | list | `[]` | Extra containers to add to the `Deployment`. |
| deploymentStrategy | object | `{"type":"Recreate"}` | [Deployment Strategy](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy). |
| revisionHistoryLimit | int | `nil` | Specify the number of old `ReplicaSets` to retain to allow rollback of the `Deployment``. |
| podLabels | object | `{}` | Labels to add to the `Pod`. |
| podAnnotations | object | `{"cluster-autoscaler.kubernetes.io/safe-to-evict":"true","kubectl.kubernetes.io/default-container":"external-dns"}` | Annotations to add to the `Pod`. |
| automountServiceAccountToken | bool | `true` | Set this to `false` to [opt out of API credential automounting](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#opt-out-of-api-credential-automounting) for the `Pod`. |
| shareProcessNamespace | bool | `false` | If `true`, the `Pod` will have [process namespace sharing](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/) enabled. |
| podSecurityContext | object | See _values.yaml_ | [Pod security context](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#podsecuritycontext-v1-core), this supports full customisation. |
| priorityClassName | string | `"giantswarm-critical"` | Priority class name for the `Pod`. |
| terminationGracePeriodSeconds | int | `nil` | Termination grace period for the `Pod` in seconds. |
| dnsPolicy | string | `nil` | [DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy) for the pod, if not set the default will be used. |
| dnsConfig | object | `nil` | [DNS config](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-dns-config) for the pod, if not set the default will be used. |
| hostAliases | list | `[]` | [Host aliases](https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/) to add to the `Pod` definition, injected into the pod's `/etc/hosts`. |
| initContainers | list | `[]` | [Init containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) to add to the `Pod` definition. |
| securityContext | object | See _values.yaml_ | [Security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) for the `external-dns` container. |
| env | list | `[]` | [Environment variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for the `external-dns` container. |
| livenessProbe | object | See _values.yaml_ | [Liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| readinessProbe | object | See _values.yaml_ | [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| extraVolumes | list | `[]` | Extra [volumes](https://kubernetes.io/docs/concepts/storage/volumes/) for the `Pod`. |
| extraVolumeMounts | list | `[]` | Extra [volume mounts](https://kubernetes.io/docs/concepts/storage/volumes/) for the `external-dns` container. |
| resources | object | `{"limits":{"memory":"100Mi"},"requests":{"cpu":"50m","memory":"100Mi"}}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for the `external-dns` container. |
| nodeSelector | object | `{}` | Node labels to match for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). |
| affinity | object | `{}` | Affinity settings for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). If an explicit label selector is not provided for pod affinity or pod anti-affinity one will be created from the pod selector labels. |
| topologySpreadConstraints | list | `[]` | Topology spread constraints for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). If an explicit label selector is not provided one will be created from the pod selector labels. |
| tolerations | list | `[]` | Node taints which will be tolerated for `Pod` [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/). |
| serviceMonitor.enabled | bool | `true` | If `true`, create a `ServiceMonitor` resource to support the _Prometheus Operator_. |
| serviceMonitor.additionalLabels | object | `{}` | Additional labels for the `ServiceMonitor`. |
| serviceMonitor.annotations | object | `{}` | Annotations to add to the `ServiceMonitor`. |
| serviceMonitor.namespace | string | `nil` | If set create the `ServiceMonitor` in an alternate namespace. |
| serviceMonitor.interval | string | `"60s"` | If set override the _Prometheus_ default interval. |
| serviceMonitor.scrapeTimeout | string | `nil` | If set override the _Prometheus_ default scrape timeout. |
| serviceMonitor.scheme | string | `nil` | If set overrides the _Prometheus_ default scheme. |
| serviceMonitor.tlsConfig | object | `{}` | Configure the `ServiceMonitor` [TLS config](https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig). |
| serviceMonitor.bearerTokenFile | string | `nil` | Provide a bearer token file for the `ServiceMonitor`. |
| serviceMonitor.relabelings | list | `[{"action":"replace","regex":";(.*)","replacement":"$1","separator":";","sourceLabels":["namespace","__meta_kubernetes_namespace"],"targetLabel":"namespace"},{"action":"replace","sourceLabels":["__meta_kubernetes_pod_label_app"],"targetLabel":"app"},{"action":"replace","sourceLabels":["__meta_kubernetes_pod_node_name"],"targetLabel":"node"}]` | [Relabel configs](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config) to apply to samples before ingestion. |
| serviceMonitor.metricRelabelings | list | `[]` | [Metric relabel configs](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#metric_relabel_configs) to apply to samples before ingestion. |
| serviceMonitor.targetLabels | list | `[]` | Provide target labels for the `ServiceMonitor`. |
| logLevel | string | `"info"` | Log level. |
| logFormat | string | `"text"` | Log format. |
| interval | string | `"5m"` | Interval for DNS updates. |
| minEventSyncInterval | string | `"30s"` |  |
| triggerLoopOnEvent | bool | `false` | If `true`, triggers run loop on create/update/delete events in addition of regular interval. |
| namespaceFilter | string | `""` |  |
| namespaced | bool | `false` | if `true`, _ExternalDNS_ will run in a namespaced scope (`Role`` and `Rolebinding`` will be namespaced too). |
| enableGatewayListenerSets | bool | `false` | if `true`, the Gateway API ListenerSet flag will be enabled. |
| gatewayNamespace | string | `nil` | _Gateway API_ gateway namespace to watch. When `namespaced=true`, setting this value avoids creating any cluster-scoped RBAC (no ClusterRole/ClusterRoleBinding) for Gateway sources. |
| sourceNamespace | string | `nil` | Source namespace to watch for Kubernetes resources other than Gateway API gateways. Used only when `namespaced=true`. Defaults to Release.Namespace |
| sources | list | `["service","crd"]` | _Kubernetes_ resources to monitor for DNS entries. |
| policy | REQUIRED | `"sync"` | How DNS records are synchronized between sources and providers; must be set explicitly to one of `create-only`, `sync`, or `upsert-only`. |
| registry | string | `"txt"` | Specify the registry for storing ownership and labels. Valid values are `txt`, `aws-sd`, `crd`, `dynamodb` & `noop`. |
| txtOwnerId | string | `nil` | Specify an identifier for this instance of _ExternalDNS_ when using a registry other than `noop`. |
| txtPrefix | string | `nil` | Specify a prefix for the domain names of TXT records created for the `txt` registry. Mutually exclusive with `txtSuffix`. |
| txtSuffix | string | `nil` | Specify a suffix for the domain names of TXT records created for the `txt` registry. Mutually exclusive with `txtPrefix`. |
| domainFilters | list | `[]` | Limit possible target zones by domain suffixes. |
| excludeDomains | list | `[]` | Intentionally exclude domains from being managed. |
| labelFilter | string | `nil` | Filter resources queried for endpoints by label selector. |
| annotationFilter | string | `"giantswarm.io/external-dns=managed"` | Filter resources queried for endpoints by annotation selector. |
| annotationPrefix | string | `"external-dns.alpha.kubernetes.io/"` | Annotation prefix for external-dns annotations (useful for split horizon DNS with multiple instances). |
| managedRecordTypes | list | `["A","CNAME"]` | Record types to manage (default: A, AAAA, CNAME) |
| provider.name | string | `"aws"` | _ExternalDNS_ provider name; for the available providers and how to configure them see [README](https://github.com/kubernetes-sigs/external-dns/blob/master/charts/external-dns/README.md#providers). |
| provider.webhook.image.repository | string | `nil` | Image repository for the `webhook` container. |
| provider.webhook.image.tag | string | `nil` | Image tag for the `webhook` container. |
| provider.webhook.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for the `webhook` container. |
| provider.webhook.env | list | `[]` | [Environment variables](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) for the `webhook` container. |
| provider.webhook.args | list | `[]` | Extra arguments to provide for the `webhook` container. |
| provider.webhook.extraVolumeMounts | list | `[]` | Extra [volume mounts](https://kubernetes.io/docs/concepts/storage/volumes/) for the `webhook` container. |
| provider.webhook.resources | object | `{}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for the `webhook` container. |
| provider.webhook.securityContext | object | See _values.yaml_ | [Pod security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) for the `webhook` container. |
| provider.webhook.livenessProbe | object | See _values.yaml_ | [Liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `external-dns` container. |
| provider.webhook.readinessProbe | object | See _values.yaml_ | [Readiness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) configuration for the `webhook` container. |
| provider.webhook.service.port | int | `8080` | Webhook exposed HTTP port for the service. |
| provider.webhook.serviceMonitor | object | See _values.yaml_ | Optional [Service Monitor](https://prometheus-operator.dev/docs/operator/design/#servicemonitor) configuration for the `webhook` container. |
| extraArgs | object | `{}` | Extra arguments to provide to _ExternalDNS_. An array or map can be used, with maps allowing for value overrides; maps also support slice values to use the same arg multiple times. |
| secretConfiguration.enabled | bool | `false` | If `true`, create a `Secret` to store sensitive provider configuration (**DEPRECATED**). |
| secretConfiguration.mountPath | string | `nil` | Mount path for the `Secret`, this can be templated. |
| secretConfiguration.subPath | string | `nil` | Sub-path for mounting the `Secret`, this can be templated. |
| secretConfiguration.data | object | `{}` | `Secret` data. |
| enabled | bool | `nil` | No effect - reserved for use in sub-charting. |
| crd.install | bool | `true` |  |
| hostNetwork | bool | `false` |  |
| serviceType | string | `"managed"` |  |
| e2e | bool | `false` |  |
| baseDomain | string | `"gigantic.io"` |  |
| clusterID | string | `"en2jo"` |  |
| ciliumNetworkPolicy.enabled | bool | `false` |  |
| kubectlApplyJob.enabled | bool | `true` |  |
| kubectlApplyJob.files[0] | string | `"files/dnsendpoints.externaldns.k8s.io.yaml"` |  |
| kubectlApplyJob.files[1] | string | `"files/dnsrecords.externaldns.k8s.io.yaml"` |  |
