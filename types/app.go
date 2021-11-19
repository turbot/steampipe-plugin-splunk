package types

type AppListResponse struct {
	Entry  []App  `json:"entry"`
	Paging Paging `json:"paging"`
}

type App struct {
	Entry
	Content AppContent `json:"content"`
}

type AppContent struct {
	CheckForUpdates            bool   `json:"check_for_updates"`
	Configured                 bool   `json:"configured"`
	Core                       bool   `json:"core"`
	Disabled                   bool   `json:"disabled"`
	Description                string `json:"description"`
	Details                    string `json:"details"`
	Version                    string `json:"version"`
	EAIACL                     ACL    `json:"eai:acl"`
	Label                      string `json:"label"`
	ManagedByDeploymentClient  bool   `json:"managed_by_deployment_client"`
	ShowInNav                  bool   `json:"show_in_nav"`
	StateChangeRequiresRestart bool   `json:"state_change_requires_restart"`
	Visible                    bool   `json:"visible"`
}
