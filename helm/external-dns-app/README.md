# external-dns-app

Configure external DNS servers for Kubernetes Ingresses and Services

![Version: 2.42.0](https://img.shields.io/badge/Version-2.42.0-informational?style=flat-square)

![AppVersion: v0.11.0](https://img.shields.io/badge/AppVersion-v0.11.0-informational?style=flat-square)

**Homepage:** <https://github.com/giantswarm/external-dns-app>

## Requirements

Kubernetes: `>=1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
|  | crd | 0.1.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| annotationFilter | string | `"giantswarm.io/external-dns=managed"` |  |
| baseDomain | string | `"gigantic.io"` |  |
| ciliumNetworkPolicy.enabled | bool | `false` |  |
| clusterID | string | `"en2jo"` |  |
| commonLabels | object | `{}` |  |
| crd.install | bool | `true` |  |
| deploymentAnnotations | object | `{}` |  |
| deploymentStrategy.type | string | `"Recreate"` |  |
| dnsPolicy | string | `nil` |  |
| domainFilters | list | `[]` |  |
| e2e | bool | `false` |  |
| env | list | `[]` |  |
| extraArgs | list | `[]` |  |
| extraVolumeMounts | list | `[]` |  |
| extraVolumes | list | `[]` |  |
| fullnameOverride | string | `""` |  |
| global.image.registry | string | `"gsoci.azurecr.io"` |  |
| hostNetwork | bool | `false` |  |
| image.name | string | `"giantswarm/external-dns"` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.registry | string | `"gsoci.azurecr.io"` |  |
| image.tag | string | `"v0.11.0"` |  |
| imagePullSecrets | list | `[]` |  |
| interval | string | `"5m"` |  |
| livenessProbe.failureThreshold | int | `2` |  |
| livenessProbe.httpGet.path | string | `"/healthz"` |  |
| livenessProbe.httpGet.port | string | `"http"` |  |
| livenessProbe.initialDelaySeconds | int | `10` |  |
| livenessProbe.periodSeconds | int | `10` |  |
| livenessProbe.successThreshold | int | `1` |  |
| livenessProbe.timeoutSeconds | int | `5` |  |
| logFormat | string | `"text"` |  |
| logLevel | string | `"info"` |  |
| minEventSyncInterval | string | `"30s"` |  |
| nameOverride | string | `""` |  |
| namespaceFilter | string | `"kube-system"` |  |
| namespaced | bool | `false` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations."cluster-autoscaler.kubernetes.io/safe-to-evict" | string | `"true"` |  |
| podAnnotations."kubectl.kubernetes.io/default-container" | string | `"external-dns"` |  |
| podLabels | object | `{}` |  |
| podSecurityContext.fsGroup | int | `65534` |  |
| podSecurityContext.runAsGroup | int | `65534` |  |
| podSecurityContext.runAsUser | int | `65534` |  |
| podSecurityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| policy | string | `"sync"` |  |
| priorityClassName | string | `"giantswarm-critical"` |  |
| provider | string | `"CHANGEME"` |  |
| rbac.additionalPermissions | list | `[]` |  |
| rbac.create | bool | `true` |  |
| readinessProbe.failureThreshold | int | `6` |  |
| readinessProbe.httpGet.path | string | `"/healthz"` |  |
| readinessProbe.httpGet.port | string | `"http"` |  |
| readinessProbe.initialDelaySeconds | int | `5` |  |
| readinessProbe.periodSeconds | int | `10` |  |
| readinessProbe.successThreshold | int | `1` |  |
| readinessProbe.timeoutSeconds | int | `5` |  |
| registry | string | `"txt"` |  |
| resources.limits.memory | string | `"100Mi"` |  |
| resources.requests.cpu | string | `"50m"` |  |
| resources.requests.memory | string | `"100Mi"` |  |
| secretConfiguration.data | object | `{}` |  |
| secretConfiguration.enabled | bool | `false` |  |
| secretConfiguration.mountPath | string | `""` |  |
| secretConfiguration.subPath | string | `""` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `65534` |  |
| service.annotations | object | `{}` |  |
| service.port | int | `7979` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.labels | object | `{}` |  |
| serviceAccount.name | string | `""` |  |
| serviceMonitor.additionalLabels | object | `{}` |  |
| serviceMonitor.annotations | object | `{}` |  |
| serviceMonitor.enabled | bool | `true` |  |
| serviceMonitor.interval | string | `"60s"` |  |
| serviceMonitor.metricRelabelings | list | `[]` |  |
| serviceMonitor.relabelings[0].action | string | `"replace"` |  |
| serviceMonitor.relabelings[0].regex | string | `";(.*)"` |  |
| serviceMonitor.relabelings[0].replacement | string | `"$1"` |  |
| serviceMonitor.relabelings[0].separator | string | `";"` |  |
| serviceMonitor.relabelings[0].sourceLabels[0] | string | `"namespace"` |  |
| serviceMonitor.relabelings[0].sourceLabels[1] | string | `"__meta_kubernetes_namespace"` |  |
| serviceMonitor.relabelings[0].targetLabel | string | `"namespace"` |  |
| serviceMonitor.relabelings[1].action | string | `"replace"` |  |
| serviceMonitor.relabelings[1].sourceLabels[0] | string | `"__meta_kubernetes_pod_label_app"` |  |
| serviceMonitor.relabelings[1].targetLabel | string | `"app"` |  |
| serviceMonitor.relabelings[2].action | string | `"replace"` |  |
| serviceMonitor.relabelings[2].sourceLabels[0] | string | `"__meta_kubernetes_pod_node_name"` |  |
| serviceMonitor.relabelings[2].targetLabel | string | `"node"` |  |
| serviceMonitor.targetLabels | list | `[]` |  |
| serviceType | string | `"managed"` |  |
| shareProcessNamespace | bool | `false` |  |
| sources[0] | string | `"service"` |  |
| terminationGracePeriodSeconds | string | `nil` |  |
| tolerations | list | `[]` |  |
| topologySpreadConstraints | list | `[]` |  |
| triggerLoopOnEvent | bool | `false` |  |
| txtOwnerId | string | `""` |  |
| txtPrefix | string | `""` |  |
| txtSuffix | string | `""` |  |
