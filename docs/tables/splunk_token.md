# Table: splunk_token

Authentication tokens let users of Splunk platform environments access REST
endpoint resources or use the Splunk CLI in Splunk Enterprise environments.

**Note:** You must enable token authentication. Please see [Enable or disable token authentication](https://docs.splunk.com/Documentation/Splunk/8.2.3/Security/EnableTokenAuth).

## Examples

### Basic info

```sql
select
  name,
  status,
  issuer,
  expiration_time
from
  splunk_token;
```

### List disabled tokens

```sql
select
  name,
  status,
  issuer,
  expiration_time
from
  splunk_token
where
  status = 'disabled';
```

### List unused tokens

```sql
select
  name,
  status,
  issuer,
  expiration_time
from
  splunk_token
where
  last_used is null;
```

### List expired tokens

```sql
select
  name,
  status,
  issuer,
  expiration_time
from
  splunk_token
where
  now()::timestamp > expiration_time;
```
