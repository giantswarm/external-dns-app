[![CircleCI](https://circleci.com/gh/giantswarm/external-dns-app.svg?style=svg)](https://circleci.com/gh/giantswarm/external-dns-app)

# external-dns-app chart

Helm chart for the [external-dns](https://github.com/kubernetes-sigs/external-dns) service running in Workload
Clusters. This chart is used to deploy both as a default app and as a Managed App.
It can be installed multiple times in the same Workload Cluster.

**What is this App?**

`external-dns` makes Kubernetes resources discoverable via external DNS servers. It dynamically configures external DNS providers ([AWS Route 53](https://aws.amazon.com/route53/), [Azure DNS](https://learn.microsoft.com/en-us/azure/dns/)) for Kubernetes Ingresses, Services etc.

**Why did we add it?**

The App is already used as a default App in most clusters (except on-prem) to provide DNS records for [Ingress NGINX Controller](https://github.com/giantswarm/ingress-nginx-app).

**Who can use it?**

Customers using Giant Swarm clusters on AWS or Azure.

---

## Index
- [Installing](#installing)
- [Configuring](#configuring)
- [Compatibility](#compatibility)
- [Limitations](#limitations)
- [Release Process](#release-process)
- [Contributing & Reporting Bugs](#contributing--reporting-bugs)

## Installing

There are 3 ways to install this app onto a workload cluster:

1. [Using our web interface](https://docs.giantswarm.io/reference/web-interface/app-catalog/)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the [App custom resource](https://docs.giantswarm.io/use-the-api/management-api/crd/apps.application.giantswarm.io/) on the Management Cluster

## Configuring

Configuration options are documented in the [README](https://github.com/giantswarm/external-dns-app/blob/main/helm/external-dns-app/README.md)
document. See also the [default `values.yaml`](https://github.com/giantswarm/external-dns-app/blob/main/helm/external-dns-app/values.yaml)

### values.yaml

This is an example of a values file you could upload using our web interface. It assumes:

- The cloud provider is AWS.
- API access is internal and therefore authentication is provided by IRSA.
- Only Ingress resources in the namespace `web-app` should be reconciled.
- Only Hosted Zone `Z262CGXUQ3M97` will be modified.

```yaml
# values.yaml
serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: <CLUSTER_ID>-Route53Manager-Role

domainFilters:
  - web-app.mydomain.com
namespaced: 'web-app'
annotationFilter: "mydomain.com/external-dns=owned"

txtPrefix: 'webapp'
sources:
  - ingress

provider: aws
extraArgs:
  - "--zone-id-filter=Z262CGXUQ3M97"
```

Additionally to the above example, `external-dns` can also be configured to synchronize `DNSEndpoint` custom resources:

```yaml
# values.yaml
...
sources:
  - crd
...
```

Here is an example `DNSEndpoint` resource:

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: my-record
  namespace: web-app
  annotations:
    mydomain.com/external-dns: owned
spec:
  endpoints:
  - dnsName: www.mydomain.com
    recordTTL: 60
    recordType: A
    targets:
    - 1.2.3.4
```

See our [full reference page on how to configure applications](https://docs.giantswarm.io/reference/app-configuration/) for more details.

## Upgrade to v3

This version includes many breaking changings from previous versions. Please read the defailed information in the [Upgrade guide](https://github.com/giantswarm/external-dns-app/blob/main/docs/upgrading.md)

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new branch to trigger the release workflow e.g. to release `v0.1.0`,
create a branch from main called `main#release#v0.1.0` and push it.
* This will push a new git tag and trigger a new tarball to be pushed to the
`default-catalog` and the `giantswarm-catalog`

## Contributing & Reporting Bugs
If you have suggestions for how `external-dns` could be improved, or want to report a bug, open an issue! We'd love all and any contributions. 

Check out the [Contributing Guide](CONTRIBUTING.md) for details on the contribution workflow, submitting patches, and reporting bugs.

[app-operator]: https://github.com/giantswarm/app-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[external-dns]: https://github.com/kubernetes-incubator/external-dns
[giantswarm-catalog]: https://github.com/giantswarm/giantswarm-catalog
[giantswarm-test-catalog]: https://github.com/giantswarm/giantswarm-test-catalog
