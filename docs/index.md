---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/splunk.svg"
brand_color: "#65a637"
display_name: "Splunk"
short_name: "splunk"
description: "Steampipe plugin to query apps, indexes, logs and more from Splunk."
og_description: "Query Splunk with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/splunk-social-graphic.png"
---

# Splunk + Steampipe

[Splunk](https://splunk.com) software is used for searching, monitoring and analyzing log data.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List indexes in your Splunk account:

```sql
select
  name,
  max_time,
  total_event_count
from
  splunk_index
```

```
+----------------+---------------------+-------------------+
| name           | max_time            | total_event_count |
+----------------+---------------------+-------------------+
| my_event_index | 2021-11-18T01:29:21 |              2345 |
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
  plugin   = "splunk"
  url      = "localhost:8089"
  username = "admin"
  password = "password"
}
```

- `url` - URL of the Splunk installation.
- `username` - Username for authentication.
- `password` - Password for authentication.


## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-splunk
- Community: [Slack Channel](https://steampipe.io/community/join)
