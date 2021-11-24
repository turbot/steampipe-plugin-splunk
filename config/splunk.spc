connection "splunk" {
  plugin = "splunk"
  
  # `url` - URL of the Splunk installation.
  # If `url` is not specified in a connection, it will be loaded from:
  #   - The value specified in the `SPLUNK_URL` environment variable.
  url = "localhost:8089"
  
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

  # Paths is a list of locations to search for Dockerfiles by default.
  # Wildcards are supported per https://golang.org/pkg/path/filepath/#Match
  # Exact file paths can have any name. Wildcard based matches must either
  # have a name of Dockerfile (e.g. Dockerfile, Dockerfile.example) or an
  # .splunkfile extension (e.g. nginx.splunkfile).
  # paths = [ "/path/to/dir/*", "/path/to/exact/custom-splunkfile-name" ]

  # Optional splunk engine configuration.
  # host        = "tcp://192.168.59.103:2376"
  # cert_path   = "/path/to/my-cert"
  # api_version = "1.41"
  # tls_verify  = true
}
