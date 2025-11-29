[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/external-dns-app/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/external-dns-app/tree/main)

# external-dns-app chart

Giant Swarm offers an `external-dns-app-bundle` Managed App which can be installed in tenant clusters.

Here we define the `external-dns-app-bundle`, `external-dns-app` charts with their templates and default configuration.

## Architecture

This repository contains two Helm charts:

- `helm/external-dns-app-bundle/`: Main chart installed on the management cluster, contains the workload cluster chart and the required AWS IAM role.
- `helm/external-dns-app/`: Workload cluster chart that contains the actual external-dns setup.

Users only need to install the bundle chart on the management cluster, which in turn will deploy the workload cluster chart.

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

Install the chart on the management cluster using an App CR:

```yaml
apiVersion: application.giantswarm.io/v1alpha1
kind: App
metadata:
  name: coyote-external-dns-app-bundle
  namespace: org-acme
spec:
  catalog: default
  config:
    configMap:
      name: coyote-cluster-values
      namespace: org-acme
  kubeConfig:
    inCluster: true
  name: external-dns-app-bundle
  namespace: org-acme
  version: 3.2.0
```

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
domainFilters:
  - web-app.mydomain.com
namespaced: 'web-app'
annotationFilter: "mydomain.com/external-dns=owned"

txtOwnerId: 'webapp'
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

Starting from this version, `external-dns-app` includes significant changes that may affect its functionality if not configured properly. For the full list of modifications, we strongly recommend referring to the [Upgrade guide](https://github.com/giantswarm/external-dns-app/blob/main/docs/upgrading.md).

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
