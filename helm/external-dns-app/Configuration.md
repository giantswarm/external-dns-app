# external-dns

This chart installs `external-dns` as managed application. ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.


## Configuration

The following table lists the configurable parameters of the external-dns chart, its dependencies and default values.

Parameter | Description | Default
--- | --- | ---
`baseDomain` | Cluster base domain. `external-dns` applies only to `aws` privder | 'aws'
`clusterID` | Cluster identifier. Applies only to Giant Swarm managed clusters | 'testid'
`provider` | Provider identifier (`aws`/`azure`/`kvm`). `external-dns` applies only to `aws`/`azure` providers | 'azure'
`route53.access` | Access type (`internal`/`external`). `internal` uses IAM role, `external` credentials for AWS sdk configuration | `internal`
