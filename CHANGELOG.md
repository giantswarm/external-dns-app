# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [3.2.0] - 2025-02-20

### Changed

- Update architect-orb and ATS.
- Add DNSEndpoints as a source for DNS records.

## [3.1.0] - 2024-02-13

Release [v3.0.0](https://github.com/giantswarm/external-dns-app/releases/tag/v3.0.0) contains some breaking changes! In case you're skipping v3.0.0, consult the [migration guide](https://github.com/giantswarm/external-dns-app/blob/main/docs/upgrading.md)

### Changed

- Remove default namespaceFilter configuration. ([#324](https://github.com/giantswarm/external-dns-app/pull/324)).

## [3.0.0] - 2023-12-13

This release contains some breaking changes! Please also consult the [migration guide](https://github.com/giantswarm/external-dns-app/blob/main/docs/upgrading.md)

### Added

- Add vendir for upstream sync.
- Add namespaced feature to scope permissions to one namespace.
- Add support for Gateway API ([#305](https://github.com/giantswarm/external-dns-app/pull/305)).

### Changed

- Deployment: Align to upstream ([#255](https://github.com/giantswarm/external-dns-app/pull/255)).
  - Use `crd.podSecurityContext` for crd job.
  - Rename `global.resources` as `resources`.
  - Rename `externalDNS.extraArgs` as `extraArgs`.
  - Rename `externalDNS.policy` as `policy`.
  - Rename `externalDNS.sources` as `sources` and adjust default value.
  - Rename `externalDNS.interval` as `interval`.
  - Rename `global.image` as `image` using helper for name composition.
  - Move `global.securityContext` to `podSecurityContext` and align names.
- Service: Align to upstream ([#243](https://github.com/giantswarm/external-dns-app/pull/243)).
  - Replace `global.metrics.port` value with `service.port`.
  - Add service annotations with GS defaults.
  - Set readinessProbe and livenessProbe from values.
  - Move podAnnotations to values.
- Update README and config docs ([#290](https://github.com/giantswarm/external-dns-app/pull/290)).
- Switch Registry to ACR ([#318](https://github.com/giantswarm/external-dns-app/pull/318)).

### Removed

- Deployment: Align to upstream ([#255](https://github.com/giantswarm/external-dns-app/pull/255)).
  - Remove dedicated option for `min-event-sync-interval` and set it in extraArgs.
  - Remove `externalDNS.dryRun` option.
- Secrets: Remove deprecated values for AWS Route53 external authentication [#266](https://github.com/giantswarm/external-dns-app/pull/266).
- Remove support for KIAM ([#278](https://github.com/giantswarm/external-dns-app/pull/278)).
- Remove `aws.iam.customRoleName` value ([#278](https://github.com/giantswarm/external-dns-app/pull/278)).
- Remove `aws`, `gcpProject` and `externalDNS` values ([#284](https://github.com/giantswarm/external-dns-app/pull/284)).
- Remove Azure volume configuration ([#284](https://github.com/giantswarm/external-dns-app/pull/284)).
- Remove unused helpers ([#290](https://github.com/giantswarm/external-dns-app/pull/290)).
- Remove PSP ([#305](https://github.com/giantswarm/external-dns-app/pull/305)).

## [2.42.0] - 2023-09-28

### Changed

- Make CRD install job compliant with PSS ([#309](https://github.com/giantswarm/external-dns-app/pull/309)).

## [2.41.0] - 2023-09-26

### Changed

- Make deployment compliant with PSS ([#307](https://github.com/giantswarm/external-dns-app/pull/307)).

## [2.40.0] - 2023-09-21

### Changed

- Replace condition for PSP CR installation.

## [2.39.0] - 2023-08-24

### Changed

- Replace monitoring labels with ServiceMonitor ([#296](https://github.com/giantswarm/external-dns-app/pull/296)).
- Update ATS to 0.4.1 and python deps ([#297](https://github.com/giantswarm/external-dns-app/pull/297)).

## [2.38.1] - 2023-08-04

### Added

- Add minAllowed in VPA to minimize OOM cycle.

### Changed

- Increase Memory limit.

## [2.38.0] - 2023-07-13

### Changed

- Move CRD jobs into a separated subchart ([#275](https://github.com/giantswarm/external-dns-app/pull/275)).
- Prepare new values for alignment ([]()).
  - Add domainFilter and extraArgs values.
  - Add interval, namepsaceFilter and minEventSyncInterval values.
  - Add txtPrefix value with higher priority.
  - Add txtOwnerId value with higher priority.
  - Add annotationFilter value with higher priority.
- Allow projected volumes across all providers ([#282](https://github.com/giantswarm/external-dns-app/pull/282)).

### Removed

- Hardcoded references to `provider==vmware` ([#277](https://github.com/giantswarm/external-dns-app/pull/277)).

## [2.37.1] - 2023-06-15

### Changed

- Remove deprecated annotation from Pod.[#265](https://github.com/giantswarm/external-dns-app/pull/265).

### Fixed

- Fix indentation in environment variables for secret injection [#272](https://github.com/giantswarm/external-dns-app/pull/272).

## [2.37.0] - 2023-05-04

### Changed

- Disable PSPs for k8s 1.25 and newer.

## [2.36.0] - 2023-05-02

### Changed

- Move initContainers and env into a helper function [#259](https://github.com/giantswarm/external-dns-app/pull/259).
- Add `safe-to-evict` annotations to allow eviction [#261](https://github.com/giantswarm/external-dns-app/pull/261).

## [2.35.1] - 2023-04-14

### Changed

- Create secret from `secretConfiguration.data` value without breaking AWS Credentials values compatibility.

## [2.35.0] - 2023-04-04

### Changed

- Make CiliumNetworkPolicy CR creation be deployed or not with a flag in the Values.

## [2.34.2] - 2023-03-23

### Added

- Configure request and limits for CRD job.

## [2.34.1] - 2023-03-22

### Added

- Add `node-role.kubernetes.io/control-plane` to crd install jobs toleration.

## [2.34.0] - 2023-03-21

### Added

- Add ServiceMonitor and default values ([#245](https://github.com/giantswarm/external-dns-app/pull/245)).

## [2.33.0] - 2023-03-07

### Added

- Add support to run in `hostNetwork` (primary used in `CAPZ` based management clusters)

## [2.23.2] - 2023-01-17

### Fixed

- Hardcode `external-dns.name` default name dropping the `-app` suffix ([#235](https://github.com/giantswarm/external-dns-app/pull/235))

## [2.23.1] - 2023-01-13

### Fixed

- Restore missing pod annotations in deployment ([#232](https://github.com/giantswarm/external-dns-app/pull/232)).

## [2.23.0] - 2023-01-13

### Changed

- Service account `irsa` annotation for `aws` and `capa` to align with `aws-pod-identity-webhook-app` changes
- Deployment: Align to upstream ([#227](https://github.com/giantswarm/external-dns-app/pull/227) [#229](https://github.com/giantswarm/external-dns-app/pull/229) [#224](https://github.com/giantswarm/external-dns-app/pull/224)).
  - Template deployment strategy from values
  - Align indentation
  - Move blocks to match upstream structure
  - Add annotations for secret reload
  - Take imagePullPolicy from values
  - Add secret's mount subpath
  - Take securityContext from values
  - Add new arguments for logging and events

## [2.22.0] - 2023-01-02

### Added

- Add projected volumes for `capa` ([#219](https://github.com/giantswarm/external-dns-app/pull/219)).
- Add nodeSelector, affinity, topologySpreadContraints and tolerations values to align to upstream ([223](https://github.com/giantswarm/external-dns-app/pull/223))

### Changed 

- ServiceAccount: Align to upstream ([#222](https://github.com/giantswarm/external-dns-app/pull/222)).
  - Labels: Add labels from values.
- Allow overrides of service account annotations ([#221](https://github.com/giantswarm/external-dns-app/pull/221)).

## [2.21.0] - 2022-12-08

### Added

- Deployment: Align to upstream ([#214](https://github.com/giantswarm/external-dns-app/pull/214)).
  - Add extraVolumes and extraVolumeMounts from values.
  - Add environment variables from values.
  - Add secretConfiguration for injecting secrets to deployment.

## [2.20.0] - 2022-12-05

### Changed

- ServiceAccount: Align to upstream ([#207](https://github.com/giantswarm/external-dns-app/pull/207)).
  - Helper: Add upstream helpers.
  - ServiceAccount: Add annotations from values.
- RBAC: Align to upstream ([#209](https://github.com/giantswarm/external-dns-app/pull/209))
  - Split rbac.yaml into clusterrole.yaml and clusterrolebinding.yaml.
  - Compose role rules based on values.
  - Rename ClusterRoleBinding.
  - Enable RBAC creation based on values.
- Deployment: Align to upstream ([#210](https://github.com/giantswarm/external-dns-app/pull/210) [#211](https://github.com/giantswarm/external-dns-app/pull/211)).
  - Add annotations from values.
  - Add labels in pods from values.
  - Add annotations in pods from values.
  - Add deployment specs.

## [2.19.0] - 2022-11-18

### Added

- CAPA provider for service account `irsa` annotation

## [2.18.0] - 2022-11-17

### Added

- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in the deployment if defined in `values.yaml`.
- Support for using `cluster-apps-operator` specific `cluster.proxy` values.

## [2.17.1] - 2022-11-15

### Changed

- Allow using AWS Route53 from any provider [#200](https://github.com/giantswarm/external-dns-app/pull/200)

## [2.17.0] - 2022-11-10

### Added

- Added `CiliumNetworkPolicy` for the CRD install job.

### Changed

- The helm job that installs CRDs is not removed if the job fails.

## [2.16.0] - 2022-11-09

### Added

- Add support for GCP workload identity for authentication.

## [2.15.4] - 2022-09-14

### Changed

- Adjust manifest to behave properly with `capa` provider.

## [2.15.3] - 2022-09-14

### Added

- Add support for `capa` provider.

## [2.15.2] - 2022-08-22

### Changed

- Update init container image to v3.16.2([#182](https://github.com/giantswarm/external-dns-app/pull/182))

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

[Unreleased]: https://github.com/giantswarm/external-dns-app/compare/v3.2.0...HEAD
[3.2.0]: https://github.com/giantswarm/external-dns-app/compare/v3.1.0...v3.2.0
[3.1.0]: https://github.com/giantswarm/external-dns-app/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/giantswarm/external-dns-app/compare/v2.42.0...v3.0.0
[2.42.0]: https://github.com/giantswarm/external-dns-app/compare/v2.41.0...v2.42.0
[2.41.0]: https://github.com/giantswarm/external-dns-app/compare/v2.40.0...v2.41.0
[2.40.0]: https://github.com/giantswarm/external-dns-app/compare/v2.39.0...v2.40.0
[2.39.0]: https://github.com/giantswarm/external-dns-app/compare/v2.38.1...v2.39.0
[2.38.1]: https://github.com/giantswarm/external-dns-app/compare/v2.38.0...v2.38.1
[2.38.0]: https://github.com/giantswarm/external-dns-app/compare/v2.37.1...v2.38.0
[2.37.1]: https://github.com/giantswarm/external-dns-app/compare/v2.37.0...v2.37.1
[2.37.0]: https://github.com/giantswarm/external-dns-app/compare/v2.36.0...v2.37.0
[2.36.0]: https://github.com/giantswarm/external-dns-app/compare/v2.35.1...v2.36.0
[2.35.1]: https://github.com/giantswarm/external-dns-app/compare/v2.35.0...v2.35.1
[2.35.0]: https://github.com/giantswarm/external-dns-app/compare/v2.34.2...v2.35.0
[2.34.2]: https://github.com/giantswarm/external-dns-app/compare/v2.34.1...v2.34.2
[2.34.1]: https://github.com/giantswarm/external-dns-app/compare/v2.34.0...v2.34.1
[2.34.0]: https://github.com/giantswarm/external-dns-app/compare/v2.33.0...v2.34.0
[2.33.0]: https://github.com/giantswarm/external-dns-app/compare/v2.23.2...v2.33.0
[2.23.2]: https://github.com/giantswarm/external-dns-app/compare/v2.23.1...v2.23.2
[2.23.1]: https://github.com/giantswarm/external-dns-app/compare/v2.23.0...v2.23.1
[2.23.0]: https://github.com/giantswarm/external-dns-app/compare/v2.22.0...v2.23.0
[2.22.0]: https://github.com/giantswarm/external-dns-app/compare/v2.21.0...v2.22.0
[2.21.0]: https://github.com/giantswarm/external-dns-app/compare/v2.20.0...v2.21.0
[2.20.0]: https://github.com/giantswarm/external-dns-app/compare/v2.19.0...v2.20.0
[2.19.0]: https://github.com/giantswarm/external-dns-app/compare/v2.18.0...v2.19.0
[2.18.0]: https://github.com/giantswarm/external-dns-app/compare/v2.17.1...v2.18.0
[2.17.1]: https://github.com/giantswarm/external-dns-app/compare/v2.17.0...v2.17.1
[2.17.0]: https://github.com/giantswarm/external-dns-app/compare/v2.16.0...v2.17.0
[2.16.0]: https://github.com/giantswarm/external-dns-app/compare/v2.15.4...v2.16.0
[2.15.4]: https://github.com/giantswarm/external-dns-app/compare/v2.15.3...v2.15.4
[2.15.3]: https://github.com/giantswarm/external-dns-app/compare/v2.15.2...v2.15.3
[2.15.2]: https://github.com/giantswarm/external-dns-app/compare/v2.15.1...v2.15.2
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
