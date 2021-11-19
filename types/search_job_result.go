package types

type SearchJobResultListRequest struct {
	ListRequest
	Search *string `url:"search"`
}

type SearchJobResultListResponse struct {
	InitOffset int               `json:"init_offset"`
	Messages   []Message         `json:"messages"`
	Preview    bool              `json:"preview"`
	Results    []SearchJobResult `json:"results"`
}

type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type SearchJobResult struct {
	Time          string `json:"_time"`
	Host          string `json:"host"`
	Source        string `json:"source"`
	SourceType    string `json:"sourcetype"`
	Action        string `json:"action"`
	Info          string `json:"info"`
	User          string `json:"user"`
	IsSearches    string `json:"is_searches"`
	IsNotSearches string `json:"is_not_searches"`
	IsModify      string `json:"is_modify"`
	IsNotModify   string `json:"is_not_modify"`
	Result        string `json:"_raw"`
}

/*
  72     {
  73       "_time": "2021-07-17T11:23:03.483-04:00",
  74       "host": "e-gineer-Mac.local",
  75       "source": "audittrail",
  76       "sourcetype": "audittrail",
  77       "action": "list_health",
  78       "info": "granted",
  79       "user": "admin",
  80       "is_searches": "0",
  81       "is_not_searches": "1",
  82       "is_modify": "0",
  83       "is_not_modify": "1",
  84       "_bkt": "_audit~0~EB550810-8FB3-4E82-A07E-5C3E6860750B",
  85       "_cd": "0:181975",
  86       "_indextime": "1626535383",
  87       "_kv": "1",
  88       "_raw": "Audit:[timestamp=07-17-2021 11:23:03.483, user=admin, action=list_health, info=granted object=\"splunkd\" operation=list]",
  89       "_serial": "0",
  90       "_si": [
  91         "e-gineer-Mac.local",
  92         "_audit"
  93       ],
  94       "_sourcetype": "audittrail",
  95       "_subsecond": ".483206"
  96     },
*/
