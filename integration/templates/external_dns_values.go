//go:build k8srequired
// +build k8srequired

package templates

// ExternalDNSValues values used for external-dns-app in integration test
const ExternalDNSValues = `
provider: inmemory
image:
  registry: quay.io`
