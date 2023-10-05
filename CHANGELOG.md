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
