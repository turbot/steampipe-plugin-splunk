connection "splunk" {
  plugin = "splunk"
  
  # `url` - URL of the Splunk installation.
  # If `url` is not specified in a connection, it will be loaded from:
  #   - The value specified in the `SPLUNK_URL` environment variable.
  url = "localhost:8089"

  # `insecure_skip_verify` - InsecureSkipVerify controls whether a client verifies the server’s certificate chain and host name. If InsecureSkipVerify is true, crypto/tls accepts any certificate presented by the server and any host name in that certificate.
  # If `insecure_skip_verify` is not specified in a connection, it will be loaded from:
  #   - The value specified in the `SPLUNK_INSECURE_SKIP_VERIFY` environment variable.
  # insecure_skip_verify = true
  
  # 1. To authenticate using username and password
  # `username` - Username for authentication.
  # `password` - Password for authentication.
  # If `username` or `password` are not specified in a connection, credentials will be loaded from:
  #   - The value specified in the `SPLUNK_USERNAME` and `SPLUNK_PASSWORD` environment variables respectively.
  # username = "admin"
  # password = "password"

  # 2. To authenticate using Splunk authentication token
  # `auth_token` - Splunk authentication token.
  # If `auth_token` is not specified in a connection, it will be loaded from:
  #   - The value specified in the `SPLUNK_AUTH_TOKEN` environment variable.
  # auth_token = "<YOUR_AUTH_TOKEN>"
}
