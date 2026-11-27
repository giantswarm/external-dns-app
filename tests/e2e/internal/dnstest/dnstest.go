// Package dnstest holds the helpers the e2e suites share to look up the DNS
// records ExternalDNS writes for a workload cluster.
package dnstest

import (
	"context"
	"fmt"
	"net"
	"time"

	. "github.com/onsi/gomega" //nolint:staticcheck

	"github.com/giantswarm/apptest-framework/v5/pkg/state"
	"github.com/giantswarm/clustertest/v5/pkg/application"
	"github.com/giantswarm/clustertest/v5/pkg/logger"
)

// WorkloadClusterBaseDomain returns the DNS zone of the workload cluster under test,
// `<cluster name>.<base domain>`.
func WorkloadClusterBaseDomain() string {
	values := &application.ClusterValues{}
	err := state.GetFramework().MC().GetHelmValues(state.GetCluster().Name, state.GetCluster().GetNamespace(), values)
	Expect(err).NotTo(HaveOccurred())
	Expect(values.BaseDomain).NotTo(BeEmpty(), "baseDomain field missing from cluster helm values")

	return fmt.Sprintf("%s.%s", state.GetCluster().Name, values.BaseDomain)
}

// Resolver returns a resolver that queries the first authoritative nameserver of
// baseDomain directly, so lookups are not delayed by caching resolvers.
func Resolver(baseDomain string) *net.Resolver {
	nsRecords, err := net.LookupNS(baseDomain)
	Expect(err).NotTo(HaveOccurred())
	Expect(nsRecords).NotTo(BeEmpty())

	nsHost := nsRecords[0].Host
	nsIPs, err := net.LookupHost(nsHost)
	Expect(err).NotTo(HaveOccurred())
	Expect(nsIPs).NotTo(BeEmpty())

	nsIP := nsIPs[0]
	logger.Log("Using nameserver %s (%s) for DNS queries", nsHost, nsIP)

	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 10 * time.Second,
			}
			return d.DialContext(ctx, "udp", nsIP+":53")
		},
	}
}
