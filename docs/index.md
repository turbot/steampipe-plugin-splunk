---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/splunk.svg"
brand_color: "#CC2B7E"
display_name: "Splunk"
short_name: "splunk"
description: "Steampipe plugin to query apps, indexes, logs and more from Splunk."
og_description: "Query Splunk with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/splunk-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Splunk + Steampipe

[Splunk](https://splunk.com) software is used for searching, monitoring and analyzing log data.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List indexes in your Splunk account:

```sql
select
  name,
  max_time,
  total_event_count
from
  splunk_index;
```

```sh
+----------------+---------------------+-------------------+
| name           | max_time            | total_event_count |
+----------------+---------------------+-------------------+
| my_event_index | 2021-11-18T01:29:21 | 2345              |
+----------------+---------------------+-------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/splunk/tables)**

## Get started

### Install

Download and install the latest Splunk plugin:

```bash
steampipe plugin install splunk
```

### Configuration

Installing the latest splunk plugin will create a config file (`~/.steampipe/config/splunk.spc`) with a single connection named `splunk`:

```hcl
connection "splunk" {
  plugin = "splunk"

  # Splunk base URL.
  # Can also be set with the SPLUNK_URL environment variable.
  # Defaults to "https://localhost:8089".
  url = "https://localhost:8089"

  # You can connect to Splunk using one of the options below:

  # 1. Authenticate using username and password
  # If `username` or `password` are not specified credentials will be loaded
  # from the `SPLUNK_USERNAME` and `SPLUNK_PASSWORD` environment variables
  # respectively.
  # username = "admin"
  # password = "password"

  # 2. Authenticate using Splunk authentication token
  # If `auth_token` is not specified, it will be loaded from the
  # `SPLUNK_AUTH_TOKEN` environment variable.
  # auth_token = "<YOUR_AUTH_TOKEN>"
}
```


