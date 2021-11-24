![image](https://hub.steampipe.io/images/plugins/turbot/splunk-social-graphic.png)

# Splunk Plugin for Steampipe

Use SQL to query apps, indexes, logs and more from Splunk.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/splunk)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/splunk/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-splunk/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install splunk
```

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-splunk.git
cd steampipe-plugin-splunk
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```bash
make
```

Configure the plugin:

```shell
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/splunk.spc
```

Try it!

```shell
steampipe query
> .inspect splunk
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-splunk/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Splunk Plugin](https://github.com/turbot/steampipe-plugin-splunk/labels/help%20wanted)
