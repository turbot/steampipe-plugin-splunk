---
title: "Steampipe Table: splunk_index - Query Splunk Indexes using SQL"
description: "Allows users to query Splunk Indexes, providing insights into indexed data, its storage, retention policies, and other related information."
---

# Table: splunk_index - Query Splunk Indexes using SQL

Splunk is a software platform widely used for monitoring, searching, analyzing and visualizing the machine-generated data. In Splunk, an Index is a place where Splunk stores the data. Indexes are used in Splunk to maintain the data, with its own set of configuration options that control its behavior.

## Table Usage Guide

The `splunk_index` table provides insights into indexed data within Splunk. As a Splunk administrator or a security analyst, explore index-specific details through this table, including data storage, retention policies, and associated metadata. Utilize it to uncover information about indexes, such as their current size, maximum size, home path, and more.

## Examples

### Basic info
Discover the segments that have the most events based on their maximum time, useful for identifying trends and high-activity periods within your data. This can be beneficial for optimizing resource allocation and planning future data management strategies.

```sql+postgres
select
  name,
  max_time,
  total_event_count
from
  splunk_index;
```

```sql+sqlite
select
  name,
  max_time,
  total_event_count
from
  splunk_index;
```

### List disabled indexes
Assess the elements within your system that have been disabled to better manage your resources and ensure optimal performance. This allows you to identify areas of inefficiency and take corrective measures.

```sql+postgres
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  disabled;
```

```sql+sqlite
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  disabled = 1;
```

### List all internal indexes
Discover the segments that encompass all internal indexes, enabling you to analyze event frequency and duration. This aids in efficient data management and optimal resource allocation.

```sql+postgres
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  is_internal;
```

```sql+sqlite
select
  name,
  max_time,
  total_event_count
from
  splunk_index
where
  is_internal = 1;
```

### Get index count by type
Determine the distribution of different types of data across indexes to optimize data management and enhance system performance.

```sql+postgres
select
  data_type,
  count(*)
from
  splunk_index
group by
  data_type;
```

```sql+sqlite
select
  data_type,
  count(*)
from
  splunk_index
group by
  data_type;
```