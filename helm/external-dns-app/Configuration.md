## Configuration

Configuration values below are relevant when external-dns is running as a second
instance in a Workload Cluster. All values are documented further in `values.yaml`,
and are are optional unless otherwise stated.

| Key                                | Type   | Default                                | Description                                                                          |
|------------------------------------|--------|----------------------------------------|--------------------------------------------------------------------------------------|
| aws.access                         | string | `"internal"`                           | Authenticate via KIAM or credentials **(required)**                                  |
| aws.baseDomain                     | string | `nil`                                  | Base domain of the cluster                                                           |
| aws.batchChangeSize                | int    | `nil`                                  | How many records to synchronise in a batch                                           |
| aws.iam.customRoleName             | string | `nil`                                  | Custom IAM role for KIAM to assume                                                   |
| aws.preferCNAME                    | bool   | `false`                                | Prefer CNAME records over ALIAS                                                      |
| aws.region                         | string | `nil`                                  | Region (required when aws.access is 'external')                                      |
| aws.zoneType                       | string | `nil`                                  | Hosted zone types to update                                                          |
| externalDNS.annotationFilter       | string | `"giantswarm.io/external-dns=managed"` | Only reconcile Service/Ingress with this annotation                                  |
| externalDNS.aws_access_key_id      | string | `nil`                                  | Access key (required when aws.access is 'external')                                  |
| externalDNS.aws_secret_access_key  | string | `nil`                                  | Secret key (required when aws.access is 'external')                                  |
| externalDNS.domainFilterList       | list   | `[]`                                   | List of domains to update                                                            |
| externalDNS.interval               | string | `nil`                                  | Synchronisation interval                                                             |
| externalDNS.namespaceFilter        | string | `"kube-system"`                        | Filter namespace to watch endpoints in                                               |
| externalDNS.registry.txtPrefix     | string | `nil`                                  | Custom prefix for TXT records (used to indicate ownership of records) **(required)** |
| externalDNS.sources[0]             | string | `"service"`                            | Filter resources to reconcile **(required)**                                         |
| global.image.name                  | string | `"giantswarm/external-dns"`            | Docker image name **(required)**                                                     |
| global.image.registry              | string | `"quay.io"`                            | Docker image registry **(required)**                                                 |
| global.image.tag                   | string | `"v0.7.6"`                             | Docker image tag **(required)**                                                      |
| global.metrics.port                | int    | `7979`                                 | Port to serve metrics on **(required)**                                              |
| global.metrics.scrape              | bool   | `true`                                 | Toggle Prometheus scrape via annotation **(required)**                               |
| global.resources                   | dict   | `{}`                                   | Yaml dict for resource limits and requests                                           |
| global.securityContext.fsGroupID   | int    | `65534`                                | fsGroup ID to run as **(required)**                                                  |
| global.securityContext.groupID     | int    | `65534`                                | Group ID to run as **(required)**                                                    |
| global.securityContext.userID      | int    | `65534`                                | User ID to run as **(required)**                                                     |
| provider                           | string | `nil`                                  | Platform provider **(required)**                                                     |

## Giant Swarm specific configuration

The following values are only relevant when external-dns is running as a default
App in Giant Swarm clusters.

| Key         | Type   | Default         | Description                                       |
|-------------|--------|-----------------|---------------------------------------------------|
| baseDomain  | string | `"gigantic.io"` | Base domain of the cluster                        |
| clusterID   | string | `"en2jo"`       | Workload cluster ID                               |
| e2e         | bool   | `false`         | If the app is running in an e2e test              |
| serviceType | string | `"managed"`     | Annotation which denotes if this is a managed App |
