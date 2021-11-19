package types

type SearchJobListRequest struct {
	ListRequest
}

type SearchJobListResponse struct {
	Entry  []SearchJob `json:"entry"`
	Paging Paging      `json:"paging"`
}

type SearchJob struct {
	Entry
	Content SearchJobContent `json:"content"`
}

type SearchJobContent struct {
	CursorTime            string      `json:"cursorTime"`
	Custom                interface{} `json:"custom"`
	Delegate              string      `json:"delegate"`
	DiskUsage             int         `json:"diskUsage"`
	DispatchState         string      `json:"dispatchState"`
	DoneProgress          float64     `json:"doneProgress"`
	DropCount             int         `json:"dropCount"`
	EventAvailableCount   int         `json:"eventAvailableCount"`
	EventCount            int         `json:"eventCount"`
	EventFieldCount       int         `json:"eventFieldCount"`
	EventIsStreaming      bool        `json:"eventIsStreaming"`
	EventIsTruncated      bool        `json:"eventIsTruncated"`
	EventPreviewableCount int         `json:"eventPreviewableCount"`
	EventSearch           string      `json:"eventSearch"`
	EventSorting          string      `json:"eventSorting"`
	IsDone                bool        `json:"isDone"`
	IsEventPreviewEnabled bool        `json:"isEventPreviewEnabled"`
	IsFailed              bool        `json:"isFailed"`
	IsFinalized           bool        `json:"isFinalized"`
	IsPaused              bool        `json:"isPaused"`
	IsPreviewEnabled      bool        `json:"isPreviewEnabled"`
	IsRealTimeSearch      bool        `json:"isRealTimeSearch"`
	IsRemoteTimeline      bool        `json:"isRemoteTimeline"`
	IsSaved               bool        `json:"isSaved"`
	IsSavedSearch         bool        `json:"isSavedSearch"`
	IsZombie              bool        `json:"isZombie"`
	Keywords              string      `json:"keywords"`
	Label                 string      `json:"label"`
	LatestTime            string      `json:"latestTime"`
	Messages              interface{} `json:"messages"`
	NumPreviews           int         `json:"numPreviews"`
	Performance           interface{} `json:"performance"`
	Priority              int         `json:"priority"`
	RemoteSearch          string      `json:"remoteSearch"`
	ReportSearch          string      `json:"reportSearch"`
	Request               interface{} `json:"request"`
	ResultCount           int         `json:"resultCount"`
	ResultIsStreaming     bool        `json:"resultIsStreaming"`
	ResultPreviewCount    int         `json:"resultPreviewCount"`
	RunDuration           float64     `json:"runDuration"`
	ScanCount             int         `json:"scanCount"`
	SearchEarliestTime    float64     `json:"searchEarliestTime"`
	SearchLatestTime      float64     `json:"searchLatestTime"`
	SearchProviders       []string    `json:"searchProviders"`
	SID                   string      `json:"sid"`
	StatusBuckets         int         `json:"statusBuckets"`
	TTL                   int         `json:"ttl"`
}
