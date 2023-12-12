package splunk

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type splunkConfig struct {
	URL                *string `hcl:"url"`
	Username           *string `hcl:"username"`
	Password           *string `hcl:"password"`
	AuthToken          *string `hcl:"auth_token"`
	InsecureSkipVerify *bool   `hcl:"insecure_skip_verify"`
}

func ConfigInstance() interface{} {
	return &splunkConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) splunkConfig {
	if connection == nil || connection.Config == nil {
		return splunkConfig{}
	}
	config, _ := connection.Config.(splunkConfig)
	return config
}
