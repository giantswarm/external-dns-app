[![CircleCI](https://circleci.com/gh/giantswarm/external-dns-app.svg?style=svg)](https://circleci.com/gh/giantswarm/external-dns-app)

# external-dns-app chart

Helm chart for the [external-dns](https://github.com/kubernetes-sigs/external-dns) service running in Workload
Clusters. This chart is used to deploy both as a default app (only required on AWS), and as a Managed App.
It can be installed multiple times in the same Workload Cluster.

**What is this App?**

`external-dns` configures external DNS servers (AWS Route53, Azure DNS) for Kubernetes Ingresses and Service.

**Why did we add it?**

The App is already used as a default App in AWS clusters to provide DNS records for [nginx-ingress-controller-app](https://github.com/giantswarm/nginx-ingress-controller-app).

**Who can use it?**

Customers using Giant Swarm clusters on AWS or Azure.

## Installing

There are 3 ways to install this app onto a tenant cluster:

1. [Using our web interface](https://docs.giantswarm.io/reference/web-interface/app-catalog/)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the App custom resource on the Control Plane

## Configuring

Configuration options are documented in the [Configuration.md](helm/external-dns-app/Configuration.md)
document. See also the [default `values.yaml`](helm/external-dns-app/values.yaml)

### values.yaml

This is an example of a values file you could upload using our web interface. It assumes:

- The cloudprovider is AWS.
- API access is internal and therefore authentication is provided by KIAM.
- Only public Hosted Zones should be updated.
- Only Ingress resources in the namespace `web-app` should be reconciled.

```yaml
# values.yaml
aws:
  iam:
    customRoleName: 'my-precreated-route53-role'
  zoneType: private

externalDNS:
  annotationFilter: "mydomain.com/external-dns=owned"
  domainFilterList:
    - web-app.mydomain.com
  namespaceFilter: 'web-app'
  registry:
    txtPrefix: 'webapp'
  sources:
    - ingress

provider: aws
```

See our [full reference page on how to configure applications](https://docs.giantswarm.io/reference/app-configuration/) for more details.

## Compatibility

This app has been tested to work with the following tenant cluster release versions:

* AWS `v13.0.0`

## Limitations

Some apps have restrictions on how they can be deployed.
Not following these limitations will most likely result in a broken deployment.

* Requires [nginx-ingress-controller-app v1.14.0](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#1140---2021-02-23) or greater to work (due to the need for the filtering annotation).
   * If you do not (or cannot) upgrade `nginx-ingress-controller-app` to `v1.14.0`,
you can work around this by running the following command to ensure the default
`external-dns` continues to reconcile the relevant Service:

```bash
kubectl -n kube-system annotate service nginx-ingress-controller-app "giantswarm.io/external-dns=managed"
```

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new branch to trigger the release workflow e.g. to release `v0.1.0`,
create a branch from master called `master#release#v0.1.0` and push it.
* This will push a new git tag and trigger a new tarball to be pushed to the
`default-catalog` and the `giantswarm-catalog`

[app-operator]: https://github.com/giantswarm/app-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[external-dns]: https://github.com/kubernetes-incubator/external-dns
[giantswarm-catalog]: https://github.com/giantswarm/giantswarm-catalog
[giantswarm-test-catalog]: https://github.com/giantswarm/giantswarm-test-catalog
