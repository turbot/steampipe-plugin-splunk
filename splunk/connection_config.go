package splunk

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type splunkConfig struct {
	URL                *string `cty:"url"`
	Username           *string `cty:"username"`
	Password           *string `cty:"password"`
	AuthToken          *string `cty:"auth_token"`
	InsecureSkipVerify *bool   `cty:"insecure_skip_verify"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"auth_token": {
		Type: schema.TypeString,
	},
	"insecure_skip_verify": {
		Type: schema.TypeBool,
	},
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
