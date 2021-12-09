package splunk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turbot/steampipe-plugin-splunk/types"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableSplunkUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_user",
		Description: "List all users.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getUser,
			KeyColumns: plugin.SingleColumn("name"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			// Other columns
			{Name: "acl", Type: proto.ColumnType_JSON, Description: "Access Control List for the user."},
			{Name: "authentication_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Type"), Description: "Authentication type: LDAP, Scripted, Splunk, System (reserved for system user)."},
			{Name: "capabilities", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Capabilities"), Description: "List of capabilities assigned to the user."},
			{Name: "default_app", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.DefaultApp"), Description: "Default app for the user, which is invoked at login."},
			{Name: "default_app_is_user_override", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.DefaultAppIsUserOverride"), Description: "True if the default app overrides the user role default app."},
			{Name: "default_app_source_role", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.DefaultAppSourceRole"), Description: "The role that determines the default app for the user, if the user has multiple roles."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Email"), Description: "User email address."},
			{Name: "locked_out", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.LockedOut"), Description: "True if the user has been locked out."},
			{Name: "real_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.RealName"), Description: "User full name."},
			{Name: "restart_background_jobs", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.RestartBackgroundJobs"), Description: "True if incomplete background search jobs should be restarted when Splunk restarts."},
			{Name: "roles", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Roles"), Description: "Roles assigned to the user."},
			{Name: "search_assistant", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.SearchAssistant"), Description: "Full search assistant is useful when first learning to create searches. Compact provides more succinct assistance."},
			{Name: "search_auto_format", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.SearchAutoFormat"), Description: "If true, automatically format search syntax to improve readability."},
			{Name: "search_line_numbers", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.SearchLineNumbers"), Description: "If true, shows numbers next to each line in the search syntax."},
			{Name: "search_syntax_highlighting", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.SearchSyntaxHighlighting"), Description: "Theme for search query syntax highlighting."},
			{Name: "search_use_advanced_editor", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.SearchUseAdvancedEditor"), Description: "The advanced editor can provide auto-formatting, line numbers, and highlight search syntax for increased readability. You can also turn off the advanced editor to use the basic search format."},
			{Name: "tz", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.TZ"), Description: "User timezone."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the user was last updated."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_user.listUser", "connection_error", err)
		return nil, err
	}

	endpoint := conn.BuildSplunkURL("services/authentication/users", nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}

	obj := types.UserListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_user.listUser", "query_error", err)
		return nil, err
	}
	for _, i := range obj.Entry {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_user.getUser", "connection_error", err)
		return nil, err
	}

	var name string
	equalQuals := d.KeyColumnQuals
	if equalQuals["name"] != nil {
		name = equalQuals["name"].GetStringValue()
	}

	endpoint := conn.BuildSplunkURL(fmt.Sprintf("services/authentication/users/%s", name), nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}
	obj := types.UserListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_user.getUser", "query_error", err)
		return nil, err
	}

	if len(obj.Entry) > 0 {
		return obj.Entry[0], nil
	}

	return nil, nil
}
