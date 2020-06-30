# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

## [v1.2.2] 2020-06-30

### Changed

- Upgrade upstream external-dns from v0.5.18 to v0.7.2.

## [v1.2.1] 2020-05-29

### Changed

- Prefer CNAMEs record sets for AWS SDK configuration with explicit credentials.

## [v1.2.0] 2020-02-04

### Changed

- Updated external-dns app version to v0.5.18.

## [v1.1.0]

### Added

- Add support AWS SDK configuration with explicit credentials.

### Changed

- Remove CPU limits.

## [v1.0.0]

### Changed

- Added support for AWS provider.

## [v0.4.1]

### Added

- Add AWS support

## [v0.4.0]

### Changed

- Migrated to be deployed via an app CR not a chartconfig CR

## v0.3.1

### Changed

- Change priority class to `system-cluster-critical`.

## v0.3.0

### Added

- Network policy that allows all egress traffic.
- Network policy that allows accessing metrics on port `10254`.

[Unreleased]: https://github.com/giantswarm/external-dns-app/compare/v1.2.2...master
[v1.2.2]: https://github.com/giantswarm/external-dns-app/compare/v1.2.1...v1.2.2
[v1.2.1]: https://github.com/giantswarm/external-dns-app/compare/v1.2.0...v1.2.1
[v1.2.0]: https://github.com/giantswarm/external-dns-app/compare/v1.1.0...v1.2.0
[v1.1.0]: https://github.com/giantswarm/external-dns-app/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/giantswarm/external-dns-app/compare/v0.4.1...v1.0.0
[v0.4.1]: https://github.com/giantswarm/external-dns-app/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/giantswarm/external-dns-app/releases/tag/v0.4.0
