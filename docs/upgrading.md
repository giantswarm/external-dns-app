# Upgrading

## Upgrading from v2 to v3.0.0

Starting from `external-dns` version `3.0.0` and later, there have been modifications to the values.yaml schema. These changes were introduced in version `2.28.0`, and the previous format has been completely removed in version 3. Therefore, it's important to take note of these updates while upgrading.

If you are utilizing a custom setup with this app, certain adjustments might be necessary before upgrading."

### Values

- `aws.access` has been removed. Check the [External DNS with AWS Route 53 and static credentials](https://docs.giantswarm.io/advanced/external-dns/aws-route53-static-creds) guide for more details.
- `aws.accountID` has been removed. Use `serviceAccount.annotations` to inject the `eks.amazonaws.com/role-arn` annotation containing the proper IAM role ARN into the Service Account.
- `aws.baseDomain` has been removed. Use `domainFilters` instead.
- `aws.batchChangeSize`, `aws.batchChangeInterval`, `aws.preferCNAME` and `aws.zonesCacheDuration` have been removed. Use the arguments `--aws-batch-change-size=N`, `--aws-batch-change-interval=SECONDS`, `--aws-prefer-cname` and `--aws-zones-cache-duration=SECONDS` in `extraArgs` if desired.
- `aws.irsa` is no longer needed, as all our clusters now support IRSA.
- `aws.region` has been removed. Set the region if needed injecting the AWS_DEFAULT_REGION variable using the `env` value.
- Additional arguments must be configured by setting `extraArgs` instead of `externalDNS.extraArgs`
- `externalDNS.annotationFilter` has been removed in favor of `annotationFilter`.
- Use `secretConfiguration` in replacement of `externalDNS.aws_access_key_id` and `externalDNS.aws_secret_access_key`. Both values have been removed.
- `externalDNS.domainFilterList` has been replaced with `domainFilters`
- `externalDNS.dryRun` has been removed. Add `--dry-run` to `extraArgs` instead.
- `externalDNS.interval` has been replaced by `interval`.
- `externalDNS.namespaceFilter` as been removed. Use `namespaceFilter` instead.
- `externalDNS.minEventSyncInterval` has been replaced by `minEventSyncInterval`.
- `externalDNS.policy` has been renamed as `policy`.
- `externalDNS.registry.txtOwnerID` and `externalDNS.registry.txtPrefix` have been renamed as `txtOwnerId` and `txtPrefix` respectivelly.
- `externalDNS.sources` has been removed. Use `sources` instead.
- `baseDomain` has been replaced by `domainFilters`.
- `proxy` and `cluster.proxy` values have been removed in favor of the `env` value.

