package splunk

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-splunk/types"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSplunkToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_token",
		Description: "List all tokens.",
		List: &plugin.ListConfig{
			Hydrate: listToken,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the token."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the token."},
			{Name: "subject", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Claims.Subject"), Description: "Subject is the user that the token represents."},
			// Other columns
			{Name: "acl", Type: proto.ColumnType_JSON, Description: "Access Control List for the token."},
			{Name: "audience", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Claims.Audience"), Description: "Audience description given to the token on creation."},
			{Name: "author", Type: proto.ColumnType_STRING, Description: "Author of this object in the system."},
			{Name: "expiration_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.Claims.ExpirationTime").Transform(transform.UnixToTimestamp), Description: "Time when the token expires."},
			{Name: "headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Headers"), Description: "Token headers including algorithm and version."},
			{Name: "identity_provider", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Claims.IdentityProvider"), Description: "Identity provider for the token, e.g. Splunk."},
			{Name: "issued_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.Claims.IssuedAt").Transform(transform.UnixToTimestamp), Description: "Time when the token was issued."},
			{Name: "issuer", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Claims.Issuer"), Description: "Issuer of the token."},
			{Name: "last_used", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.LastUsed").Transform(transform.UnixToTimestamp), Description: "Time when the token was last used."},
			{Name: "last_used_ip", Type: proto.ColumnType_IPADDR, Transform: transform.FromField("Content.LastUsedIP"), Description: "IP address the token was last used from."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links for the token resource."},
			{Name: "not_before", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.Claims.NotBefore").Transform(transform.UnixToTimestamp), Description: "Time when the token becomes valid for use."},
			{Name: "roles", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Claims.Roles"), Description: "Roles assigned to the token."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Status"), Description: "Status of the token: enabled, disabled."},
			{Name: "updated", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the token was last updated."},
		},
	}
}

func listToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_token.listToken", "connection_error", err)
		return nil, err
	}

	endpoint := conn.BuildSplunkURL("services/authorization/tokens", nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}

	obj := types.TokenListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_token.listToken", "query_error", err)
		return nil, err
	}
	for _, i := range obj.Entry {
		d.StreamListItem(ctx, i)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			break
		}
	}
	return nil, nil
}
