package splunk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turbot/steampipe-plugin-splunk/types"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableSplunkApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_app",
		Description: "List all apps installed locally.",
		List: &plugin.ListConfig{
			Hydrate: listApp,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getApp,
			KeyColumns: plugin.SingleColumn("name"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the app."},
			{Name: "label", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Label"), Description: "App name."},
			{Name: "version", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Version"), Description: "App version."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Description"), Description: "App description."},
			// Other columns
			{Name: "acl", Type: proto.ColumnType_JSON, Description: "Access Control List for the app."},
			{Name: "author", Type: proto.ColumnType_STRING, Description: "Author of this object in the system."},
			{Name: "check_for_updates", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.CheckForUpdates"), Description: "If true, then check Splunkbase for app updates."},
			{Name: "configured", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Configured"), Description: "If true, then Custom app setup is complete."},
			{Name: "core", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Core"), Description: ""},
			{Name: "details", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Details"), Description: "URL to use for detailed information about the app."},
			{Name: "disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Disabled"), Description: "If true, the app is disabled."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the app."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links for the app resource."},
			{Name: "managed_by_deployment_client", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.ManagedByDeploymentClient"), Description: ""},
			{Name: "show_in_nav", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.ShowInNav"), Description: ""},
			{Name: "state_change_requires_restart", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.StateChangeRequiresRestart"), Description: ""},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the app was last updated."},
			{Name: "visible", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Visible"), Description: "If true, app is visible and navigable from Splunk Web."},
		},
	}
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_app.listApp", "connection_error", err)
		return nil, err
	}

	params := types.ListRequest{}

	if d.QueryContext.Limit != nil {
		params.Count = d.QueryContext.Limit
	}

	count := int64(0)
	for {
		params.Offset = types.Int64(count)
		endpoint := conn.BuildSplunkURL("services/apps/local", params)
		data, err := conn.Get(endpoint)
		if err != nil {
			plugin.Logger(ctx).Error("splunk_app.listApp", "query_error", err)
			return nil, err
		}
		obj := types.AppListResponse{}
		err = json.Unmarshal(data, &obj)
		if err != nil {
			plugin.Logger(ctx).Error("splunk_app.listApp", "parse_error", err)
			return nil, err
		}
		for _, i := range obj.Entry {
			d.StreamListItem(ctx, i)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}

			count++
		}
		if count >= *obj.Paging.Total {
			break
		}
	}

	return nil, nil
}

func getApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_app.getApp", "connection_error", err)
		return nil, err
	}

	var name string
	equalQuals := d.KeyColumnQuals
	if equalQuals["name"] != nil {
		name = equalQuals["name"].GetStringValue()
	}

	endpoint := conn.BuildSplunkURL(fmt.Sprintf("services/apps/local/%s", name), nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}
	obj := types.AppListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_app.getApp", "query_error", err)
		return nil, err
	}

	if len(obj.Entry) > 0 {
		return obj.Entry[0], nil
	}

	return nil, nil
}
