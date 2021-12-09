# Table: splunk_app

An app is an application that runs on the Splunk platform. Apps are designed to
analyze and display knowledge around a specific data source or data set. An app
might include any or all of the following configurations: Dashboards and
supporting searches that integrate knowledge of the data source and structure.

## Examples

### Basic info

```sql
select
  name,
  version,
  author
from
  splunk_app;
```

### List apps that are not visible

```sql
select
  name,
  version,
  author,
  visible
from
  splunk_app
where
  not visible;
```

### List disabled apps

```sql
select
  name,
  version,
  author,
  disabled
from
  splunk_app
where
  disabled;
```

### List apps with auto-update check enabled

```sql
select
  name,
  version,
  author,
  check_for_updates
from
  splunk_app
where
  check_for_updates;
```
