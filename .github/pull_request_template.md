<!--
@team-cabbage will be automatically requested for review once
this PR has been submitted.
-->

This PR:



---

## Checklist

- [ ] Added a CHANGELOG entry

## Testing

The instance of external-dns installed as part of Giant Swarm platform releases watches services in the `kube-system` namespace with annotations `giantswarm.io/external-dns=managed` and `external-dns.alpha.kubernetes.io/hostname` matching the clusters base domain. (You can find this in the deployments args `--domain-filter` value)

You can take this example `Service`, apply it to your cluster. Change the `external-dns.alpha.kubernetes.io/hostname` annotation to match your clusters base domain.

then:

- Check external-dns logs for lines like `Desired change: CREATE test.your.configured.domain.gigantic.io CNAME`
- Try to resolve the domain (`https://www.dnstester.net/`)

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    external-dns.alpha.kubernetes.io/hostname: test.your.configured.domain.gigantic.io
    external-dns.alpha.kubernetes.io/ttl: "60"
    giantswarm.io/external-dns: managed
  name: test-external-dns
  namespace: kube-system
spec:
  type: ExternalName
  externalName: www.giantswarm.io
```

For testing upgrades:

- Create the service and check for creation
- Upgrade
- Delete the service and check for deletion

### Default app on AWS releases

- [ ] Fresh install works
- [ ] Upgrade works

### Default app on Azure releases

- [ ] Fresh install works
- [ ] Upgrade works

### Optional app (KVM)

- [ ] Fresh install works
- [ ] Upgrade works

