package splunk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/turbot/steampipe-plugin-splunk/types"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableSplunkSearchJob(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_search_job",
		Description: "List all search jobs.",
		List: &plugin.ListConfig{
			Hydrate: listSearchJob,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSearchJob,
			KeyColumns: plugin.SingleColumn("sid"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "sid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.SID"), Description: "The search ID number."},
			{Name: "event_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.EventCount"), Description: "The number of events returned by the search."},
			{Name: "run_duration", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Content.RunDuration"), Description: "Time in seconds that the search took to complete."},
			// Other columns
			{Name: "cursor_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.CursorTime"), Description: "The earliest time from which no events are later scanned. Can be used to indicate progress. See description for doneProgress."},
			{Name: "custom", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Custom"), Description: "Custom job property."},
			{Name: "delegate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Delegate"), Description: "For saved searches, specifies jobs that were started by the user. Defaults to scheduler."},
			{Name: "disk_usage", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.DiskUsage"), Description: "The total amount of disk space used, in bytes."},
			{Name: "dispatch_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.DispatchState"), Description: "The state of the search. Can be any of QUEUED, PARSING, RUNNING, PAUSED, FINALIZING, FAILED, DONE."},
			{Name: "done_progress", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Content.DoneProgress"), Description: "A number between 0 and 1.0 that indicates the approximate progress of the search. doneProgress = (latestTime – cursorTime) / (latestTime – earliestTime)."},
			{Name: "drop_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.DropCount"), Description: "For real-time searches only, the number of possible events that were dropped due to the rt_queue_size (default to 100000)."},
			{Name: "earliest_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.EarliestTime"), Description: "The earliest time a search job is configured to start. Can be used to indicate progress."},
			{Name: "event_available_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.EventAvailableCount"), Description: "The number of events that are available for export."},
			{Name: "event_field_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.EventFieldCount"), Description: "The number of fields found in the search results."},
			{Name: "event_is_streaming", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.EventIsStreaming"), Description: "Indicates if the events of this search are being streamed."},
			{Name: "event_is_truncated", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.EventIsTruncated"), Description: "Indicates if events of the search are not stored, making them unavailable from the events endpoint for the search."},
			{Name: "event_search", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.EventSearch"), Description: "Subset of the entire search that is before any transforming commands. The timeline and events endpoint represents the result of this part of the search."},
			{Name: "event_sorting", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.EventSorting"), Description: "Indicates if the events of this search are sorted, and in which order: asc, desc, none."},
			{Name: "is_done", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsDone"), Description: "Indicates if the search has completed."},
			{Name: "is_event_preview_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsEventPreviewEnabled"), Description: "Indicates if the timeline_events_preview setting is enabled in limits.conf."},
			{Name: "is_failed", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsFailed"), Description: "Indicates if there was a fatal error executing the search. For example, invalid search string syntax."},
			{Name: "is_finalized", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsFinalized"), Description: "Indicates if the search was finalized (stopped before completion)."},
			{Name: "is_paused", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsPaused"), Description: "Indicates if the search is paused."},
			{Name: "is_preview_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsPreviewEnabled"), Description: "Indicates if previews are enabled."},
			{Name: "is_real_time_search", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsRealTimeSearch"), Description: "Indicates if the search is a real time search."},
			{Name: "is_remote_timeline", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsRemoteTimeline"), Description: "Indicates if the remote timeline feature is enabled."},
			{Name: "is_saved", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsSaved"), Description: "Indicates that the search job is saved, storing search artifacts on disk for 7 days from the last time that the job was viewed or touched. Add or edit the default_save_ttl value in limits.conf to override the default value of 7 days."},
			{Name: "is_saved_search", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsSavedSearch"), Description: "Indicates if this is a saved search run using the scheduler."},
			{Name: "is_zombie", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsZombie"), Description: "Indicates if the process running the search is dead, but with the search not finished."},
			{Name: "keywords", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Keywords"), Description: "All positive keywords used by this search. A positive keyword is a keyword that is not in a NOT clause."},
			{Name: "label", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.Label"), Description: "Custom name created for this search."},
			{Name: "latest_time", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.LatestTime"), Description: "The latest time a search job is configured to start. Can be used to indicate progress."},
			{Name: "messages", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Messages"), Description: "Errors and debug messages."},
			{Name: "num_previews", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.NumPreviews"), Description: "Number of previews generated so far for this search job."},
			{Name: "performance", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Performance"), Description: "A representation of the execution costs."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.Priority"), Description: "An integer between 0-10 that indicates the search priority."},
			{Name: "remote_search", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.RemoteSearch"), Description: "The search string that is sent to every search peer."},
			{Name: "report_search", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ReportSearch"), Description: "If reporting commands are used, the reporting search."},
			{Name: "request", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.Request"), Description: "GET arguments that the search sends to splunkd."},
			{Name: "result_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.ResultCount"), Description: "The total number of results returned by the search. In other words, this is the subset of scanned events (represented by the scanCount) that actually matches the search terms."},
			{Name: "result_is_streaming", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.ResultIsStreaming"), Description: "Indicates if the final results of the search are available using streaming (for example, no transforming operations)."},
			{Name: "result_preview_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.ResultPreviewCount"), Description: "The number of result rows in the latest preview results."},
			{Name: "scan_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.ScanCount"), Description: "The number of events that are scanned or read off disk."},
			{Name: "search_earliest_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.SearchEarliestTime").Transform(transform.UnixToTimestamp), Description: "Specifies the earliest time for a search, as specified in the search command rather than the earliestTime parameter. It does not snap to the indexed data time bounds for all-time searches (something that earliestTime/latestTime does)."},
			{Name: "search_latest_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.SearchLatestTime").Transform(transform.UnixToTimestamp), Description: "Specifies the latest time for a search, as specified in the search command rather than the latestTime parameter. It does not snap to the indexed data time bounds for all-time searches (something that earliestTime/latestTime does)."},
			{Name: "search_providers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Content.SearchProviders"), Description: "A list of all the search peers that were contacted."},
			{Name: "status_buckets", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.StatusBuckets"), Description: "Maximum number of timeline buckets."},
			{Name: "ttl", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.TTL"), Description: "The time to live, or time before the search job expires after it completes."},
		},
	}
}

func listSearchJob(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_search_job.listSearchJob", "connection_error", err)
		return nil, err
	}

	params := types.SearchJobListRequest{}

	if d.QueryContext.Limit != nil {
		params.Count = d.QueryContext.Limit
	}

	count := int64(0)
	for {
		params.Offset = types.Int64(count)
		endpoint := conn.BuildSplunkURL("services/search/jobs", params)
		data, err := conn.Get(endpoint)
		if err != nil {
			return nil, err
		}
		obj := types.SearchJobListResponse{}
		err = json.Unmarshal(data, &obj)
		if err != nil {
			plugin.Logger(ctx).Error("splunk_search_job.listSearchJob", "query_error", err)
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

func getSearchJob(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_search_job.getSearchJob", "connection_error", err)
		return nil, err
	}

	var searchID string
	equalQuals := d.KeyColumnQuals
	if equalQuals["sid"] != nil {
		searchID = equalQuals["sid"].GetStringValue()
	}

	endpoint := conn.BuildSplunkURL(fmt.Sprintf("services/search/jobs/%s", searchID), nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}
	obj := types.SearchJobListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_search_job.getSearchJob", "query_error", err)
		return nil, err
	}

	if len(obj.Entry) > 0 {
		return obj.Entry[0], nil
	}

	return nil, nil
}
