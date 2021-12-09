# Table: splunk_search_job_result

Retrieves the results of a given search job. A job is a process that tracks
information about the ad hoc search or saved search. The information that is
tracked includes the owner of the job, the app that the job was run on, how
many events were returned, and how long the job took to run.

To query `splunk_search_job_result` table, **you must specify the search ID** in the where or join clause (`where sid=`, `join splunk_search_job_result on sid=`).

**Notes:**

- You can specify `time` in a `where` clause in order to query results for a specific time period.
- You can also use `query` in a `where` clause to specify a [search expression](https://docs.splunk.com/Documentation/SCS/current/SearchReference/SearchCommandOverview) that allows you to filter results. Please see [search command syntax details](https://docs.splunk.com/Documentation/SCS/current/SearchReference/SearchCommandSyntaxDetails) and related documentation for more information.

## Examples

### Basic info

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685';
```

## Time Examples

### List search results for the last two days

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685'
  and time >= (current_date - interval '2' day);
```

### List search results in a specific time range

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685'
  and time <= '2021-11-24T14:05:50Z'
  and time >= '2021-11-25T14:05:50Z';
```

## Query Examples

### List search results from the 'main' index for the past 24 hours

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685'
  and time >= (current_date - interval '1' day)
  and query = 'search index=main';
```

### List completed search results performed by splunk-system-user in the last week

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685'
  and time >= (current_date - interval '7' day)
  and query = 'search user=splunk-system-user AND info=completed';
```

### List search results related to authentication changes for the past 2 days

```sql
select
  sid,
  user,
  source_type,
  time,
  result
from
  splunk_search_job_result
where
  sid = '1637923584.685'
  and time >= (current_date - interval '2' day)
  and query = 'search action=change_authentication';
```
