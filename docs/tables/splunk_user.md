---
title: "Steampipe Table: splunk_user - Query Splunk Users using SQL"
description: "Allows users to query Splunk Users, specifically providing details about each user including their role, real name, and email. This table offers insights into user management and role assignment in Splunk."
---

# Table: splunk_user - Query Splunk Users using SQL

Splunk is a software platform widely used for monitoring, searching, analyzing and visualizing the machine-generated data. In Splunk, a user is an entity with a set of capabilities that determine the actions that the user can take and the resources the user can access. Users are assigned roles that define the actions they can take and the resources they can access.

## Table Usage Guide

The `splunk_user` table provides insights into user management within Splunk. As a system administrator or security analyst, explore user-specific details through this table, including their assigned roles, real names, and emails. Utilize it to manage user access, understand role assignment, and ensure the right privileges are assigned to the right users.

## Examples

### Basic info
Determine the areas in which user data, such as name, email, and timezone, can be analyzed from the user base. This can be useful for gaining insights into user behavior and preferences, enabling more targeted and effective communication strategies.

```sql+postgres
select
  name,
  real_name,
  email,
  authentication_type,
  tz as user_timezone
from
  splunk_user;
```

```sql+sqlite
select
  name,
  real_name,
  email,
  authentication_type,
  tz as user_timezone
from
  splunk_user;
```

### List users using role-based user authentication
Determine the areas in which role-based user authentication is being used by examining the user list. This can assist in understanding the security measures in place and identifying any potential vulnerabilities or inconsistencies.

```sql+postgres
select
  name,
  email,
  authentication_type,
  roles
from
  splunk_user;
```

```sql+sqlite
select
  name,
  email,
  authentication_type,
  roles
from
  splunk_user;
```

### List users with admin privileges
Determine the areas in which certain users have been granted administrative privileges. This can be essential for managing access control and ensuring system security.

```sql+postgres
select
  name,
  email,
  authentication_type,
  roles
from
  splunk_user
where
  roles ?| array['admin'];
```

```sql+sqlite
Error: SQLite does not support array operations.
```

### List locked-out users
Identify instances where users are locked out of their accounts, a crucial step for maintaining security and ensuring user accessibility. This can be particularly useful for administrators to quickly address and resolve such issues.

```sql+postgres
select
  name,
  email,
  authentication_type,
  locked_out
from
  splunk_user
where
  locked_out;
```

```sql+sqlite
select
  name,
  email,
  authentication_type,
  locked_out
from
  splunk_user
where
  locked_out = 1;
```