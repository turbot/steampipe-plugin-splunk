## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#35](https://github.com/turbot/steampipe-plugin-splunk/pull/35))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#35](https://github.com/turbot/steampipe-plugin-splunk/pull/35))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#33](https://github.com/turbot/steampipe-plugin-splunk/pull/33))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#33](https://github.com/turbot/steampipe-plugin-splunk/pull/33))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#29](https://github.com/turbot/steampipe-plugin-splunk/pull/29))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#29](https://github.com/turbot/steampipe-plugin-splunk/pull/29))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-splunk/blob/main/docs/LICENSE). ([#29](https://github.com/turbot/steampipe-plugin-splunk/pull/29))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#28](https://github.com/turbot/steampipe-plugin-splunk/pull/28))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#21](https://github.com/turbot/steampipe-plugin-splunk/pull/21))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#19](https://github.com/turbot/steampipe-plugin-splunk/pull/19))
- Recompiled plugin with Go version `1.21`. ([#19](https://github.com/turbot/steampipe-plugin-splunk/pull/19))

## v0.4.0 [2023-04-07]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#15](https://github.com/turbot/steampipe-plugin-splunk/pull/15))

## v0.3.0 [2022-10-11]

_Enhancements_

- Updated `url` config argument to handle passing in URLs with the protocol, e.g., `https://prd-p-abcd1.splunkcloud.com`. If a URL without a protocol is passed in, e.g., `prd-p-abcd1.splunkcloud.com`, HTTPS will be used. ([#13](https://github.com/turbot/steampipe-plugin-splunk/pull/13))

## v0.2.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#10](https://github.com/turbot/steampipe-plugin-splunk/pull/10))
- Recompiled plugin with Go version `1.19`. ([#10](https://github.com/turbot/steampipe-plugin-splunk/pull/10))

## v0.1.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#6](https://github.com/turbot/steampipe-plugin-splunk/pull/6))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#5](https://github.com/turbot/steampipe-plugin-splunk/pull/5))

## v0.0.1 [2021-12-09]

_What's new?_

- New tables added
  - [splunk_app](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_app)
  - [splunk_index](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_index)
  - [splunk_search_job](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_search_job)
  - [splunk_search_job_result](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_search_job_result)
  - [splunk_token](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_token)
  - [splunk_user](https://hub.steampipe.io/plugins/turbot/splunk/tables/splunk_user)
