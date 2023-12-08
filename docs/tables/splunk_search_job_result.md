---
title: "Steampipe Table: splunk_search_job_result - Query Splunk Search Job Results using SQL"
description: "Allows users to query Splunk Search Job Results, specifically the details of each search job result, providing insights into search patterns and potential anomalies."
---

# Table: splunk_search_job_result - Query Splunk Search Job Results using SQL

Splunk is a software platform that allows you to search, analyze, and visualize the data gathered from the components of your IT infrastructure or business. It collects logs from different sources and stores them in an index from where it can be searched. Splunk gives you the ability to leverage machine data to gain insights into opportunities, risks, and operational efficiency.

## Table Usage Guide

The `splunk_search_job_result` table provides insights into search job results within Splunk. As a data analyst or security specialist, explore search job result-specific details through this table, including search patterns, associated metadata, and potential anomalies. Utilize it to uncover information about search jobs, such as those with unusual patterns, the relationships between different search jobs, and the verification of search results.

## Examples

### Basic info
Analyze the details of specific search jobs to understand their source type, the time they were conducted, and their results. This could be useful for auditing purposes, tracking user activity, or troubleshooting issues with certain search jobs.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which search results from the last two days can be analyzed. This is useful for gaining insights into recent user activity or data changes.

```sql+postgres
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

```sql+sqlite
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
  and time >= date('now','-2 day');
```

### List search results in a specific time range
Explore the search results within a specific time frame to gain insights into user activity and source types. This can help in identifying patterns or anomalies in the data during that particular period.

```sql+postgres
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

```sql+sqlite
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
This example demonstrates how to identify search results from a specific index within the last 24 hours. This can be useful for reviewing recent activity or events in a specific area of interest.

```sql+postgres
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

```sql+sqlite
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
  and time >= date('now','-1 day')
  and query = 'search index=main';
```

### List completed search results performed by splunk-system-user in the last week
Explore the completed search activities performed by a system user over the past week. This is useful for auditing system usage and identifying patterns or anomalies in system-user interactions.

```sql+postgres
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

```sql+sqlite
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
  and time >= date('now','-7 day')
  and query = 'search user=splunk-system-user AND info=completed';
```

### List search results related to authentication changes for the past 2 days
Explore recent changes to authentication settings over the past two days. This is useful to monitor and track any alterations made, ensuring the integrity and security of the system.

```sql+postgres
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

```sql+sqlite
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
  and time >= date('now','-2 day')
  and query = 'search action=change_authentication';
```