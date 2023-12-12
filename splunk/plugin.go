package splunk

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-splunk",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"splunk_app":               tableSplunkApp(ctx),
			"splunk_index":             tableSplunkIndex(ctx),
			"splunk_search_job":        tableSplunkSearchJob(ctx),
			"splunk_search_job_result": tableSplunkSearchJobResult(ctx),
			"splunk_token":             tableSplunkToken(ctx),
			"splunk_user":              tableSplunkUser(ctx),
		},
	}
	return p
}
