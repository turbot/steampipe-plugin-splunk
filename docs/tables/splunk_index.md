# Table: splunk_index

Retrieves all indexes installed locally. An index in Splunk is simply a repository for the data. It is stored on an indexer, which is a Splunk instance configured to index local and remote data. The indexed data can then be searched through a search app.

## Examples

### Basic info

```sql
select
  name,
  max_time,
  total_event_count
from
  splunk_index;
```

### List disabled indexes

```sql
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  disabled;
```

### List all internal indexes

```sql
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  is_internal;
```

### Get index count by type

```sql
select
  data_type,
  count(*)
from
  splunk_index
group by
  data_type;
```
