---
title: "Steampipe Table: splunk_search_job - Query Splunk Search Jobs using SQL"
description: "Allows users to query Splunk Search Jobs, providing details about the search jobs executed in the Splunk environment."
---

# Table: splunk_search_job - Query Splunk Search Jobs using SQL

Splunk Search Jobs is a feature within Splunk that enables users to execute searches and store the results. These search jobs can be saved for future reference or used to create alerts, visualizations, and reports. It provides a way to manage and track the progress of all searches executed in the Splunk environment.

## Table Usage Guide

The `splunk_search_job` table provides insights into the search jobs executed within the Splunk environment. As a security analyst or a Splunk administrator, you can explore job-specific details through this table, including search parameters, status, and associated metadata. Utilize it to uncover information about search jobs, such as those running for an extended period, the progress of specific search jobs, and the verification of search parameters.

## Examples

### Basic info
Analyze the settings to understand the performance and resource usage of search jobs in Splunk. This can help in identifying any jobs that are consuming excessive resources or taking too long to run, thereby aiding in efficient resource management and performance optimization.

```sql+postgres
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job;
```

```sql+sqlite
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job;
```

### List incomplete search jobs
Determine the areas in which search jobs are still in progress to manage resources effectively by identifying jobs that are consuming significant resources and time.

```sql+postgres
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job
where
  not is_done;
```

```sql+sqlite
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job
where
  is_done is not 1;
```

### List failed search jobs
Identify instances where search jobs have failed to gain insights into potential issues that may be impacting system performance or data accuracy.

```sql+postgres
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job
where
  is_failed;
```

```sql+sqlite
select
  sid,
  event_count,
  run_duration,
  earliest_time,
  disk_usage
from
  splunk_search_job
where
  is_failed = 1;
```