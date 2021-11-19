package types

type IndexListRequest struct {
	ListRequest
	DataType *string `url:"datatype"`
}

type IndexListResponse struct {
	Entry  []Index `json:"entry"`
	Paging Paging  `json:"paging"`
}

type Index struct {
	Entry
	Content IndexContent `json:"content"`
}

type IndexContent struct {
	AssureUTF8               bool        `json:"assureUTF8"`
	BlockSignSize            int         `json:"blockSignSize"`
	BlockSignatureDatabase   string      `json:"blockSignatureDatabase"`
	ColdPath                 string      `json:"coldPath"`
	ColdPathExpanded         string      `json:"coldPath_expanded"`
	ColdToFrozenDir          string      `json:"coldToFrozenDir"`
	ColdToFrozenScript       string      `json:"coldToFrozenScript"`
	CurrentDBSizeMB          string      `json:"currentDBSizeMB"`
	DataType                 string      `json:"dataType"`
	DefaultDatabase          string      `json:"defaultDatabase"`
	Disabled                 bool        `json:"disabled"`
	EnableRealtimeSearch     bool        `json:"enableRealtimeSearch"`
	FrozenTimePeriodInSecs   int         `json:"frozenTimePeriodInSecs"`
	HomePath                 string      `json:"homePath"`
	HomePathExpanded         string      `json:"homePath_expanded"`
	IndexThreads             string      `json:"indexThreads"`
	IsInternal               bool        `json:"isInternal"`
	IsReady                  bool        `json:"isReady"`
	LastInitTime             int         `json:"lastInitTime"`
	MaxConcurrentOptimizes   int         `json:"maxConcurrentOptimizes"`
	MaxDataSize              string      `json:"maxDataSize"`
	MaxHotBuckets            string      `json:"maxHotBuckets"`
	MaxHotIdleSecs           int         `json:"maxHotIdleSecs"`
	MaxHotSpanSecs           int         `json:"maxHotSpanSecs"`
	MaxMemMB                 int         `json:"maxMemMB"`
	MaxMetaEntries           int         `json:"maxMetaEntries"`
	MaxRunningProcessGroups  int         `json:"maxRunningProcessGroups"`
	MaxTime                  string      `json:"maxTime"`
	MaxTotalDataSizeMB       int         `json:"maxTotalDataSizeMB"`
	MaxWarmDBCount           int         `json:"maxWarmDBCount"`
	MemPoolMB                string      `json:"memPoolMB"`
	MinRawFileSyncSecs       string      `json:"minRawFileSyncSecs"`
	MinTime                  string      `json:"minTime"`
	PartialServiceMetaPeriod int         `json:"partialServiceMetaPeriod"`
	QuarantineFutureSecs     int         `json:"quarantineFutureSecs"`
	QuarantinePastSecs       int         `json:"quarantinePastSecs"`
	RawChunkSizeBytes        int         `json:"rawChunkSizeBytes"`
	RotatePeriodInSecs       int         `json:"rotatePeriodInSecs"`
	ServiceMetaPeriod        int         `json:"serviceMetaPeriod"`
	Summarize                bool        `json:"summarize"`
	SuppressBannerList       interface{} `json:"suppressBannerList"`
	Sync                     int         `json:"sync"`
	SyncMeta                 bool        `json:"syncMeta"`
	ThawedPath               string      `json:"thawedPath"`
	ThawedPathExpanded       string      `json:"thawedPath_expanded"`
	ThrottleCheckPeriod      int         `json:"throttleCheckPeriod"`
	TotalEventCount          int         `json:"totalEventCount"`
}
