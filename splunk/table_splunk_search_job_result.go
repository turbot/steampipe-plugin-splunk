package splunk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-splunk/types"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableSplunkSearchJobResult(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_search_job_result",
		Description: "List results for a given search job.",
		List: &plugin.ListConfig{
			Hydrate: listSearchJobResult,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "sid"},
				{Name: "query", Require: plugin.Optional},
				{Name: "time", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the result."},
			{Name: "result", Type: proto.ColumnType_STRING, Description: "The search result."},
			// Other columns
			{Name: "action", Type: proto.ColumnType_STRING, Description: "Action of the result."},
			{Name: "host", Type: proto.ColumnType_STRING, Description: "Host the result came from."},
			{Name: "info", Type: proto.ColumnType_STRING, Description: "Info of the result."},
			{Name: "is_modify", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_not_modify", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_not_searches", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_searches", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Query to further refine search job results."},
			{Name: "sid", Type: proto.ColumnType_STRING, Transform: transform.FromQual("sid"), Description: "The search job ID."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "Source of the result."},
			{Name: "source_type", Type: proto.ColumnType_STRING, Description: "Source type of the result."},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "User of the result."},
		},
	}
}

func listSearchJobResult(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_search_job.listSearchJobResult", "connection_error", err)
		return nil, err
	}

	params := types.SearchJobResultListRequest{}
	exactQuals := d.KeyColumnQuals
	searchURL := fmt.Sprintf("services/search/jobs/%s/results", exactQuals["sid"].GetStringValue())
	queryParts := []string{}

	if exactQuals["query"] != nil {
		queryParts = append(queryParts, exactQuals["query"].GetStringValue())
	}

	if d.QueryContext.Limit != nil {
		queryParts = append(queryParts, fmt.Sprintf("head %d", *d.QueryContext.Limit))
	}

	// Comparison values
	quals := d.Quals

	if quals["time"] != nil {
		for _, q := range quals["time"].Quals {
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch q.Operator {
			case ">":
				queryParts = append(queryParts, fmt.Sprintf("where _time>%d", tsSecs))
			case ">=":
				queryParts = append(queryParts, fmt.Sprintf("where _time>=%d", tsSecs))
			case "<":
				queryParts = append(queryParts, fmt.Sprintf("where _time<%d", tsSecs))
			case "<=":
				queryParts = append(queryParts, fmt.Sprintf("where _time<=%d", tsSecs))
			case "=":
				queryParts = append(queryParts, fmt.Sprintf("where _time=%d", tsSecs))
			}
		}
	}

	query := strings.Join(queryParts, " | ")
	params.Search = types.String(query)

	endpoint := conn.BuildSplunkURL(searchURL, params)
	plugin.Logger(ctx).Warn("splunk_search_job.listSearchJobResult", "endpoint", endpoint)

	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}
	obj := types.SearchJobResultListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_search_job.listSearchJobResult", "query_error", err)
		return nil, err
	}
	for _, i := range obj.Results {
		// Format time
		formatTime, _ := time.Parse(time.RFC3339, i.Time)
		utcTime := formatTime.UTC()
		i.Time = utcTime.Format("2006-01-02T15:04:05Z")

		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
