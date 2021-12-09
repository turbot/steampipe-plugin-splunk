# Table: splunk_user

Users can be assigned different privileges to configure the overall deployment,
manage other users' accounts and permissions, search for events, and build
reports.

## Examples

### Basic info

```sql
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

```sql
select
  name,
  email,
  authentication_type,
  roles
from
  splunk_user;
```

### List users with admin privileges

```sql
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

### List locked-out users

```sql
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
