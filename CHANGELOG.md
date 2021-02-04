# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/external-dns-app/compare/v2.0.2...HEAD
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
