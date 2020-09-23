[![CircleCI](https://circleci.com/gh/giantswarm/external-dns-app.svg?style=svg)](https://circleci.com/gh/giantswarm/external-dns-app)

# external-dns-app

Helm chart for the [external-dns](https://github.com/kubernetes-sigs/external-dns) service running in guest clusters

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/external-dns-app.git
$ cd external-dns-app
$ helm install helm/external-dns-app
```

Provide a custom `values.yaml`:

```bash
$ helm install external-dns-app -f values.yaml
```

Deployment to Tenant Clusters is handled by [app-operator](https://github.com/giantswarm/app-operator).

## Configuration

Configuration options are documented in [Configuration.md](helm/external-dns-app/Configuration.md) document.

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release with the version e.g. `v0.1.0` and link the
changelog entry.
* This will push a new git tag and trigger a new tarball to be pushed to the
[default-catalog].  
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/cluster-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[external-dns]: https://github.com/kubernetes-incubator/external-dns
