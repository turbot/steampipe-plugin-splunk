# Table: splunk_search_job_result

Retrieves the results of a given search job. A job is a process that tracks information about the ad hoc search or saved search. The information that is tracked includes the owner of the job, the app that the job was run on, how many events were returned, and how long the job took to run.

To query `splunk_search_job_result` table, **you must specify the search ID** in the where or join clause (`where sid=`, `join splunk_search_job_result on sid=`).

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

### Filter search result using time

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
  and time <= '2021-11-24T14:05:50Z';
```

### List all search results related to 'main' index

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
  and query = 'search index=main';
```

### List all completed search results performed by splunk-system-user

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
  and query = 'search user=splunk-system-user AND info=completed';
```

### List all search results related to authentication changes

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
  and query = 'search action=change_authentication';
```
