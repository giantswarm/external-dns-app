# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Update all container image versions ([#182](https://github.com/giantswarm/external-dns-app/pull/182))

## [2.15.1] - 2022-08-02

### Changed

- Update alpine image to v3.15.5 ([#178](https://github.com/giantswarm/external-dns-app/pull/178))

## [2.15.0] - 2022-07-01

### Changed

- Update test dependencies and py-helm-charts version to [0.7.0](https://github.com/giantswarm/pytest-helm-charts/blob/master/CHANGELOG.md) ([#173](https://github.com/giantswarm/external-dns-app/pull/173))
- Ignore IRSA annotation for service account when using AWS `external` access.

## [2.14.0] - 2022-05-31

### Added

- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.

## [2.14.0] - 2022-05-31

### Added

- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.

## [2.13.0] - 2022-05-25

### Added

- Support IRSA (IAM Roles for Service Accounts) for `aws` provider.

## [2.12.0] - 2022-05-13

### Added

- Add support for Google Cloud `gcp` provider.

## [2.11.0] - 2022-05-10

### Added

- Add support for `DNSEndpoint` CRs. See README for further information.

## [2.10.0] - 2022-04-20

### Added

- Add team ownership label.

### Changed

- Upgrade upstream external-dns image from [v0.10.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.10.2) to [v0.11.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.11.0).

## [2.9.1] - 2022-02-03

### Changed

- Allow setting the AWS default region (`aws.region`) indepentent from any other value.

## [2.9.0] - 2022-01-24

This release contains some changes to mitigate rate limiting on AWS clusters. Please take note of the defaults
for values `aws.batchChangeInterval`, `aws.zonesCacheDuration`, `externalDNS.interval`
and `externalDNS.minEventSyncInterval`.

If you already specify `--aws-batch-change-interval` or `--aws-zones-cache-duration`, please migrate to the new values `aws.batchChangeInterval` and `aws.zonesCacheDuration`.

### Added

- Allow to set `--aws-batch-change-interval` through `aws.batchChangeInterval` value. Default `10s`.
- Allow to set `--aws-zones-cache-duration` through `aws.zonesCacheDuration` value. Default `3h`.

### Changed

- Set default `externalDNS.interval` to `5m`.
- Set default `externalDNS.minEventSyncInterval` to `30s`.
- Allow setting Route53 credentials (`externalDNS.aws_access_key_id` and `externalDNS.aws_secret_access_key`) indepentent from `aws.access` value.
- Allow setting the AWS default region (`aws.region`) indepentent from `aws.access` value.
- Allow to omit the `--domain-filter` flag completely by setting `externalDNS.domainFilterList` to `null`.

## [2.8.0] - 2022-01-12

### Changed

- Add ability to specify extra arguments to the external-dns deployment through `externalDNS.extraArgs`.

## [2.7.0] - 2021-12-16

### Changed

- Upgrade upstream external-dns from v0.9.0 to [v0.10.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.10.2). The new release brings a lot of smaller improvements and bug fixes.
- Remove support for Kubernetes <= 1.18.

### Fixed

- Fix dry-run option.

## [2.6.1] - 2021-10-15

### Changed

- Updated icon

## [2.6.0] - 2021-08-31

### Added

- Add support for CAPZ clusters by detecting the Azure configuration file location.

## [2.5.0] - 2021-08-18

### Changed

- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.

## [2.4.0] - 2021-06-16

### Changed

- Upgrade upstream external-dns from v0.7.6 to [v0.8.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.8.0).
- Allow to configure the minimum interval between two consecutive synchronizations triggered from kubernetes events through `externalDNS.minEventSyncInterval`.

## [2.3.1] - 2021-05-20

### Changed

- Increase memory limit to 100Mi since we ran into out of memory problems. This will make the app more stable.

## [2.3.0] - 2021-04-06

### Changed

- Change default annotation filter to match the one we use for the nginx ingress controller.

### Added

- Add sidecar container for `provider: aws` to periodically validate IAM credential acessibility ([#76](https://github.com/giantswarm/external-dns-app/pull/76))

## [2.2.2] - 2021-03-26

### Changed

- Set docker.io as the default registry

## [2.2.1] - 2021-03-25

### Fixed

- Adds additional options required for vmware installations. ([#74](https://github.com/giantswarm/external-dns-app/pull/74))

## [2.2.0] - 2021-03-24

### Added

- Add crd source if the provider is vmware. ([#72](https://github.com/giantswarm/external-dns-app/pull/72))

## [2.1.1] - 2021-02-08

### Fixed

- Ensure CNAMEs are always used when AWS access is external. ([#62](https://github.com/giantswarm/external-dns-app/pull/62))

## [2.1.0] - 2021-02-05

### Added

- Allow the sync policy to be configured. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports customisation of the txt-owner-id (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Supports dry-run mode and warns the user if enabled. ([#60](https://github.com/giantswarm/external-dns-app/pull/60))

### Changed

- Rework the way the txt prefix is generated (whilst still defaulting for default apps). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))
- Rework how the annotation filter value is generated (whilst still defaulting for default app). ([#60](https://github.com/giantswarm/external-dns-app/pull/60))

## [2.0.2] - 2021-02-04

### Fixed

- Revert location of AWS API credentials in `values.yaml`. ([#57](https://github.com/giantswarm/external-dns-app/pull/57))

## [2.0.1] - 2021-02-03

### Changed

- Only template Secret if both required values are present in `values.yaml`. ([#53](https://github.com/giantswarm/external-dns-app/pull/53))

## [2.0.0] - 2021-02-03

### Changed

- Reworked the App to prepare it for customer use. ([#49](https://github.com/giantswarm/external-dns-app/pull/49))
  - General:
    - Pushes the app to the giantswarm app catalog.
    - Uses Helm release namespace.
    - Uses the release name for resource naming to avoid conflicts.
    - Added a values schema to catch incorrect values.
    - Generally makes the chart easier to use (fully documented values file).
  - external-dns options:
    - Allows customisation of the txt registry prefix.
    - Allows configuration of synchronisation interval.
    - Filter resources to reconcile via annotations.
  - AWS-specifc:
    - Allows the user to provide an IAM role to use.
    - Allows the user to provide the list of domains for external-dns to manage.
    - Allows configuration of batch size.
    - Allows configuration of CNAME instead of ALIAS records.
    - Allows configuration of the AWS zone type to update.

## [1.6.0] - 2021-01-27

### Changed

- Upgrade upstream external-dns from v0.7.4 to [v0.7.6](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.6).

## [1.5.0] - 2020-10-07

### Changed

- Upgrade upstream external-dns from v0.7.3 to [v0.7.4](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.4).

## [1.4.0] - 2020-08-21

### Added

- Monitoring headless Service.
- More Giant Swarm custom monitoring annotations.
- Explicitly expose metrics container port.

### Changed

- Use default external-dns metrics port 7979.

## [1.3.0] - 2020-08-18

### Added

- Added monitoring and common labels.

### Changed

- Upgrade upstream external-dns from v0.7.2 to [v0.7.3](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.3).
- Upgrade architect-orb to 0.10.1

## [1.2.2] 2020-06-30

### Changed

- Upgrade upstream external-dns from v0.5.18 to v0.7.2.

## [1.2.1] 2020-05-29

### Changed

- Prefer CNAMEs record sets for AWS SDK configuration with explicit credentials.

## [1.2.0] 2020-02-04

### Changed

- Updated external-dns app version to v0.5.18.

## [1.1.0]

### Added

- Add support AWS SDK configuration with explicit credentials.

## [1.0.1]

### Changed

- Remove CPU limits.

## [1.0.0]

### Changed

- Added support for AWS provider.

## [0.4.1]

### Added

- Add AWS support

## [0.4.0]

### Changed

- Migrated to be deployed via an app CR not a chartconfig CR

## 0.3.1

### Changed

- Change priority class to `system-cluster-critical`.

## 0.3.0

### Added

- Network policy that allows all egress traffic.
- Network policy that allows accessing metrics on port `10254`.

[Unreleased]: https://github.com/giantswarm/external-dns-app/compare/v2.15.1...HEAD
[2.15.1]: https://github.com/giantswarm/external-dns-app/compare/v2.15.0...v2.15.1
[2.15.0]: https://github.com/giantswarm/external-dns-app/compare/v2.14.0...v2.15.0
[2.14.0]: https://github.com/giantswarm/external-dns-app/compare/v2.13.0...v2.14.0
[2.13.0]: https://github.com/giantswarm/external-dns-app/compare/v2.12.0...v2.13.0
[2.12.0]: https://github.com/giantswarm/external-dns-app/compare/v2.11.0...v2.12.0
[2.11.0]: https://github.com/giantswarm/external-dns-app/compare/v2.10.0...v2.11.0
[2.10.0]: https://github.com/giantswarm/external-dns-app/compare/v2.9.1...v2.10.0
[2.9.1]: https://github.com/giantswarm/external-dns-app/compare/v2.9.0...v2.9.1
[2.9.0]: https://github.com/giantswarm/external-dns-app/compare/v2.8.0...v2.9.0
[2.8.0]: https://github.com/giantswarm/external-dns-app/compare/v2.7.0...v2.8.0
[2.7.0]: https://github.com/giantswarm/external-dns-app/compare/v2.6.1...v2.7.0
[2.6.1]: https://github.com/giantswarm/external-dns-app/compare/v2.6.0...v2.6.1
[2.6.0]: https://github.com/giantswarm/external-dns-app/compare/v2.5.0...v2.6.0
[2.5.0]: https://github.com/giantswarm/external-dns-app/compare/v2.4.0...v2.5.0
[2.4.0]: https://github.com/giantswarm/external-dns-app/compare/v2.3.1...v2.4.0
[2.3.1]: https://github.com/giantswarm/external-dns-app/compare/v2.3.0...v2.3.1
[2.3.0]: https://github.com/giantswarm/external-dns-app/compare/v2.2.2...v2.3.0
[2.2.2]: https://github.com/giantswarm/external-dns-app/compare/v2.2.1...v2.2.2
[2.2.1]: https://github.com/giantswarm/external-dns-app/compare/v2.2.0...v2.2.1
[2.2.0]: https://github.com/giantswarm/external-dns-app/compare/v2.1.1...v2.2.0
[2.1.1]: https://github.com/giantswarm/external-dns-app/compare/v2.1.0...v2.1.1
[2.1.0]: https://github.com/giantswarm/external-dns-app/compare/v2.0.2...v2.1.0
[2.0.2]: https://github.com/giantswarm/external-dns-app/compare/v2.0.1...v2.0.2
[2.0.1]: https://github.com/giantswarm/external-dns-app/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/giantswarm/external-dns-app/compare/v1.6.0...v2.0.0
[1.6.0]: https://github.com/giantswarm/external-dns-app/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/giantswarm/external-dns-app/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/giantswarm/external-dns-app/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/external-dns-app/compare/v1.2.2...v1.3.0
[1.2.2]: https://github.com/giantswarm/external-dns-app/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/giantswarm/external-dns-app/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/giantswarm/external-dns-app/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/giantswarm/external-dns-app/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/giantswarm/external-dns-app/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/external-dns-app/compare/v0.4.1...v1.0.0
[0.4.1]: https://github.com/giantswarm/external-dns-app/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/external-dns-app/releases/tag/v0.4.0
