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

func tableSplunkIndex(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "splunk_index",
		Description: "List all indexes installed locally.",
		List: &plugin.ListConfig{
			Hydrate: listIndex,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "data_type", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getIndex,
			KeyColumns: plugin.SingleColumn("name"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the index."},
			{Name: "current_db_size_mb", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.CurrentDBSizeMB"), Description: "Total size, in MB, of data stored in the index. The total incudes data in the home, cold and thawed paths."},
			{Name: "data_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.DataType"), Description: "The type of index: event, metric."},
			{Name: "total_event_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.TotalEventCount"), Description: "Total number of events in the index."},
			// Other columns
			{Name: "assure_utf8", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.AssureUTF8"), Description: "Indicates whether all data retreived from the index is proper UTF8. If enabled (set to True), degrades indexing performance. This is a global setting, not a per index setting."},
			{Name: "block_sign_size", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.BlockSignSize"), Description: "Controls How many events make up a block for block signatures. If this is set to 0, block signing is disabled for this index."},
			{Name: "block_signature_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.BlockSignatureDatabase"), Description: "The index that stores block signatures of events. This is a global setting, not a per index setting."},
			{Name: "cold_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ColdPath"), Description: "Filepath to the cold databases for the index."},
			{Name: "cold_path_expanded", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ColdPathExpanded"), Description: "Absoute filepath to the cold databases."},
			{Name: "cold_to_frozen_dir", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ColdToFrozenDir"), Description: "Destination path for the frozen archive. Used as an alternative to a coldToFrozenScript. Splunk software automatically puts frozen buckets in this directory."},
			{Name: "cold_to_frozen_script", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ColdToFrozenScript"), Description: "Path to the archiving script."},
			{Name: "default_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.DefaultDatabase"), Description: "If no index destination information is available in the input data, the index shown here is the destination of such data."},
			{Name: "disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Disabled"), Description: "If no index destination information is available in the input data, the index shown here is the destination of such data."},
			{Name: "enable_realtime_search", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.EnableRealtimeSearch"), Description: "Indicates if this is a real-time search. This is a global setting, not a per index setting."},
			{Name: "frozen_time_period_in_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.FrozenTimePeriodInSecs"), Description: "Number of seconds after which indexed data rolls to frozen. Defaults to 188697600 (6 years). Freezing data means it is removed from the index. If you need to archive your data, refer to coldToFrozenDir and coldToFrozenScript parameter documentation."},
			{Name: "home_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.HomePath"), Description: "An absolute path that contains the hot and warm buckets for the index."},
			{Name: "home_path_expanded", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.HomePathExpanded"), Description: "An absolute filepath to the hot and warm buckets for the index."},
			{Name: "index_threads", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.IndexThreads"), Description: "Number of threads used for indexing. This is a global setting, not a per index setting."},
			{Name: "is_internal", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsInternal"), Description: "True if this is an internal index (for example, _internal, _audit)."},
			{Name: "is_ready", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.IsReady"), Description: "True if the index is properly initialized."},
			{Name: "last_init_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.LastInitTime").Transform(transform.UnixToTimestamp), Description: "Last time the index processor was successfully initialized. This is a global setting, not a per index setting."},
			{Name: "max_concurrent_optimizes", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxConcurrentOptimizes"), Description: "The number of concurrent optimize processes that can run against a hot bucket. This number should be increased if instructed by Splunk Support. Typically the default value should suffice."},
			{Name: "max_data_size", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.MaxDataSize"), Description: "The maximum size in MB for a hot DB to reach before a roll to warm is triggered. Specifying 'auto' or 'auto_high_volume' causes Splunk software to autotune this parameter (recommended). Use 'auto_high_volume' for high volume indexes (such as the main index); otherwise, use 'auto'. A 'high volume index' is typically one that gets over 10GB of data per day. 'auto' sets the size to 750MB. 'auto_high_volume' sets the size to 10GB on 64-bit, and 1GB on 32-bit systems. Although the maximum value you can set this is 1048576 MB, which corresponds to 1 TB, a reasonable number ranges anywhere from 100 - 50000. Any number outside this range should be approved by Splunk Support before proceeding."},
			{Name: "max_hot_buckets", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.MaxHotBuckets"), Description: "Maximum hot buckets that can exist per index. Defaults to 3. When maxHotBuckets is exceeded, Splunk software rolls the least recently used (LRU) hot bucket to warm. Both normal hot buckets and quarantined hot buckets count towards this total. This setting operates independently of maxHotIdleSecs, which can also cause hot buckets to roll."},
			{Name: "max_hot_idle_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxHotIdleSecs"), Description: "Maximum life, in seconds, of a hot bucket. Defaults to 0. A value of 0 turns off the idle check (equivalent to INFINITE idle time). If a hot bucket exceeds maxHotIdleSecs, Splunk software rolls it to warm. This setting operates independently of maxHotBuckets, which can also cause hot buckets to roll."},
			{Name: "max_hot_span_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxHotSpanSecs"), Description: "Upper bound of target maximum timespan of hot/warm buckets in seconds. Defaults to 7776000 seconds (90 days)."},
			{Name: "max_mem_db", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxMemMB"), Description: "The amount of memory, in MB, allocated for indexing. This is a global setting, not a per index setting."},
			{Name: "max_meta_entries", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxMemMB"), Description: "Sets the maximum number of unique lines in .data files in a bucket, which may help to reduce memory consumption. If set to 0, this setting is ignored (it is treated as infinite). If exceeded, a hot bucket is rolled to prevent further increase. If your buckets are rolling due to Strings.data hitting this limit, the culprit may be the punct field in your data. If you do not use punct, it may be best to simply disable this (see props.conf.spec in $SPLUNK_HOME/etc/system/README)."},
			{Name: "max_running_process_groups", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxRunningProcessGroups"), Description: "Maximum number of processes that the indexer fires off at a time. This is a global setting, not a per index setting."},
			{Name: "max_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.MaxTime").Transform(transform.NullIfZeroValue), Description: "ISO8601 timestamp of the newest event time in the index."},
			{Name: "max_total_data_size_mb", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxTotalDataSizeMB"), Description: "The maximum size of an index, in MB."},
			{Name: "max_warm_db_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.MaxWarmDBCount"), Description: "The maximum number of warm buckets. If this number is exceeded, the warm bucket/s with the lowest value for their latest times are moved to cold."},
			{Name: "mem_pool_mb", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.MemPoolMB"), Description: "Determines how much memory is given to the indexer memory pool. This is a global setting, not a per-index setting."},
			{Name: "min_raw_file_sync_secs", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.MinRawFileSyncSecs"), Description: "Can be either an integer (or 'disable'). Some filesystems are very inefficient at performing sync operations, so only enable this if you are sure it is needed. The integer sets how frequently splunkd forces a filesystem sync while compressing journal slices. During this period, uncompressed slices are left on disk even after they are compressed. Then splunkd forces a filesystem sync of the compressed journal and removes the accumulated uncompressed files. If 0 is specified, splunkd forces a filesystem sync after every slice completes compressing. Specifying 'disable' disables syncing entirely: uncompressed slices are removed as soon as compression is complete."},
			{Name: "min_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Content.MinTime").Transform(transform.NullIfZeroValue), Description: "ISO8601 timestamp of the oldest event time in the index."},
			{Name: "partial_service_meta_period", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.PartialServiceMetaPeriod"), Description: "Related to serviceMetaPeriod. By default it is turned off (zero). If set, it enables metadata sync every <integer> seconds, but only for records where the sync can be done efficiently in-place, without requiring a full re-write of the metadata file. Records that require full re-write are be sync'ed at serviceMetaPeriod. partialServiceMetaPeriod specifies, in seconds, how frequently it should sync. Zero means that this feature is turned off and serviceMetaPeriod is the only time when metadata sync happens. If the value of partialServiceMetaPeriod is greater than serviceMetaPeriod, this setting has no effect."},
			{Name: "quarantine_future_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.QuarantineFutureSecs"), Description: "Events with timestamp of quarantineFutureSecs newer than 'now' that are dropped into quarantine bucket. Defaults to 2592000 (30 days). This is a mechanism to prevent main hot buckets from being polluted with fringe events."},
			{Name: "quarantine_past_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.QuarantinePastSecs"), Description: "Events with timestamp of quarantinePastSecs older than 'now' are dropped into quarantine bucket. Defaults to 77760000 (900 days). This is a mechanism to prevent the main hot buckets from being polluted with fringe events."},
			{Name: "raw_chunk_size_bytes", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.RawChunkSizeBytes"), Description: "Target uncompressed size in bytes for individual raw slice in the rawdata journal of the index. Defaults to 131072 (128KB). 0 is not a valid value. If 0 is specified, rawChunkSizeBytes is set to the default value."},
			{Name: "rotate_period_in_secs", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.RotatePeriodInSecs"), Description: "Rotation period, in seconds, that specifies how frequently to check: If a new hot bucket needs to be created. If there are any cold buckets that should be frozen. If there are any buckets that need to be moved out hot and cold DBs, due to size constraints."},
			{Name: "service_meta_period", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.ServiceMetaPeriod"), Description: "Defines how frequently metadata is synced to disk, in seconds. Defaults to 25 (seconds). You may want to set this to a higher value if the sum of your metadata file sizes is larger than many tens of megabytes, to avoid the hit on I/O in the indexing fast path."},
			{Name: "summarize", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.Summarize"), Description: "If true, leaves out certain index details, which provides a faster response."},
			// TODO - should be JSON?
			{Name: "suppress_banner_list", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.SuppressBannerList"), Description: "List of indexes for which we suppress 'index missing' warning banner messages. This is a global setting, not a per index setting."},
			{Name: "sync", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.Sync"), Description: "Specifies the number of events that trigger the indexer to sync events. This is a global setting, not a per index setting."},
			{Name: "sync_meta", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Content.SyncMeta"), Description: "When true, a sync operation is called before file descriptor is closed on metadata file updates. This functionality improves integrity of metadata files, especially in regards to operating system crashes/machine failures."},
			{Name: "thawed_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ThawedPath"), Description: "An absolute path that contains the thawed (resurrected) databases for the index."},
			{Name: "thawed_path_expanded", Type: proto.ColumnType_STRING, Transform: transform.FromField("Content.ThawedPathExpanded"), Description: "Absolute filepath to the thawed (resurrected) databases."},
			{Name: "throttle_check_period", Type: proto.ColumnType_INT, Transform: transform.FromField("Content.ThrottleCheckPeriod"), Description: "Defines how frequently Splunk software checks for index throttling condition, in seconds. Defaults to 15 (seconds)."},
		},
	}
}

func listIndex(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_index.listIndex", "connection_error", err)
		return nil, err
	}

	params := types.IndexListRequest{DataType: types.String("all")}

	equalQuals := d.KeyColumnQuals
	if equalQuals["data_type"] != nil {
		params.DataType = types.String(equalQuals["data_type"].GetStringValue())
	}

	if d.QueryContext.Limit != nil {
		params.Count = d.QueryContext.Limit
	}

	count := int64(0)
	for {
		params.Offset = types.Int64(count)
		endpoint := conn.BuildSplunkURL("services/data/indexes", params)
		data, err := conn.Get(endpoint)
		if err != nil {
			return nil, err
		}
		obj := types.IndexListResponse{}
		err = json.Unmarshal(data, &obj)
		if err != nil {
			plugin.Logger(ctx).Error("splunk_index.listIndex", "query_error", err)
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

func getIndex(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_index.getIndex", "connection_error", err)
		return nil, err
	}

	var name string
	equalQuals := d.KeyColumnQuals
	if equalQuals["name"] != nil {
		name = equalQuals["name"].GetStringValue()
	}

	endpoint := conn.BuildSplunkURL(fmt.Sprintf("services/data/indexes/%s", name), nil)
	data, err := conn.Get(endpoint)
	if err != nil {
		return nil, err
	}
	obj := types.IndexListResponse{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		plugin.Logger(ctx).Error("splunk_index.getIndex", "query_error", err)
		return nil, err
	}

	if len(obj.Entry) > 0 {
		return obj.Entry[0], nil
	}

	return nil, nil
}
