package types

type UserListResponse struct {
	Entry  []User `json:"entry"`
	Paging Paging `json:"paging"`
}

type User struct {
	Entry
	Content UserContent `json:"content"`
}

type UserContent struct {
	Capabilities             []string `json:"capabilities"`
	DefaultApp               string   `json:"defaultApp"`
	DefaultAppIsUserOverride bool     `json:"defaultAppIsUserOverride"`
	DefaultAppSourceRole     string   `json:"defaultAppSourceRole"`
	EAIACL                   ACL      `json:"eai:acl"`
	Email                    string   `json:"email"`
	Lang                     string   `json:"lang"`
	LockedOut                bool     `json:"locked-out"`
	RealName                 string   `json:"realname"`
	RestartBackgroundJobs    bool     `json:"restart_background_jobs"`
	Roles                    []string `json:"roles"`
	SearchAssistant          string   `json:"search_assistant"`
	SearchAutoFormat         bool     `json:"search_auto_format"`
	SearchLineNumbers        bool     `json:"search_line_numbers"`
	SearchSyntaxHighlighting string   `json:"search_syntax_highlighting"`
	SearchUseAdvancedEditor  bool     `json:"search_use_advanced_editor"`
	Type                     string   `json:"type"`
	TZ                       string   `json:"tz"`
}
