package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-splunk/splunk"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: splunk.Plugin})
}
