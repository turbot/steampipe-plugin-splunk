---
title: "Steampipe Table: splunk_app - Query Splunk Apps using SQL"
description: "Allows users to query Splunk Apps, specifically to gather information about the applications available in the Splunk environment, including their version, description, and visibility."
---

# Table: splunk_app - Query Splunk Apps using SQL

Splunk Apps are applications that are designed to enhance and extend the functionality of the Splunk platform. These apps provide features like dashboards, reports, alerts, data inputs, and workflows, tailored for specific use-cases or data sources. They are a crucial part of the Splunk ecosystem, enabling users to customize and optimize their data analysis and visualization.

## Table Usage Guide

The `splunk_app` table provides insights into the applications available in the Splunk environment. As a data analyst or a security professional, you can explore app-specific details through this table, including the app version, description, and visibility. Utilize it to uncover information about the apps, such as their configuration, status, and the data sources they are designed to handle.

## Examples

### Basic info
Explore the basic details of your Splunk applications such as their names, versions, and authors. This can help in tracking the app updates and understanding the source of each application.

```sql+postgres
select
  name,
  version,
  author
from
  splunk_app;
```

```sql+sqlite
select
  name,
  version,
  author
from
  splunk_app;
```

### List apps that are not visible
Determine the apps that are hidden or not readily visible to users, useful for understanding which applications may not be fully accessible or utilized within your system.

```sql+postgres
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

```sql+sqlite
select
  name,
  version,
  author,
  visible
from
  splunk_app
where
  visible = 0;
```

### List disabled apps
Determine the areas in which apps are disabled to understand their impact on the system's functionality. This can be useful in identifying potential issues or areas for improvement in the system's performance.

```sql+postgres
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

```sql+sqlite
select
  name,
  version,
  author,
  disabled
from
  splunk_app
where
  disabled = 1;
```

### List apps with auto-update check enabled
Discover the segments that have enabled the auto-update feature in their apps. This can assist in maintaining up-to-date applications and ensuring the latest features and security measures are in place.

```sql+postgres
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

```sql+sqlite
select
  name,
  version,
  author,
  check_for_updates
from
  splunk_app
where
  check_for_updates = 1;
```