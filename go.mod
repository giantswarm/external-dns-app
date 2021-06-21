module github.com/giantswarm/external-dns-app

go 1.15

require (
	github.com/giantswarm/apptest v0.11.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	k8s.io/apimachinery v0.20.4
)

replace k8s.io/client-go => k8s.io/client-go v0.20.4
