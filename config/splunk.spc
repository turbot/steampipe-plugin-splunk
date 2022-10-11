connection "splunk" {
  plugin = "splunk"

  # Splunk base URL.
  # Can also be set with the SPLUNK_URL environment variable.
  # Defaults to "https://localhost:8089".
  url = "https://localhost:8089"

  # You can connect to Splunk using one of the options below:

  # 1. Authenticate using username and password
  # If `username` or `password` are not specified credentials will be loaded
  # from the `SPLUNK_USERNAME` and `SPLUNK_PASSWORD` environment variables
  # respectively.
  # username = "admin"
  # password = "password"

  # 2. Authenticate using Splunk authentication token
  # If `auth_token` is not specified, it will be loaded from the
  # `SPLUNK_AUTH_TOKEN` environment variable.
  # auth_token = "<YOUR_AUTH_TOKEN>"
}
