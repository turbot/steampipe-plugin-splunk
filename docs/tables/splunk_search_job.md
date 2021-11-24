# Table: splunk_search_job

A job is a process that tracks information about the ad hoc search or saved search. The information that is tracked includes the owner of the job, the app that the job was run on, how many events were returned, and how long the job took to run.
Each time you run a search, create a pivot, open a report, or load a dashboard panel, the Splunk software creates a job in the system. When you run a search, you are creating an ad hoc search. Pivots, reports, and panels are powered by saved searches.

## Examples

### Basic info

```sql
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

```sql
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

### List failed search jobs

```sql
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
