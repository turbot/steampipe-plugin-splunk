package types

type ACLPerms struct {
	Read  []string `json:"read"`
	Write []string `json:"write"`
}

type ACL struct {
	App            string   `json:"app"`
	CanChangePerms bool     `json:"can_change_perms"`
	CanShareApp    bool     `json:"can_share_app"`
	CanShareGlobal bool     `json:"can_share_global"`
	CanShareUser   bool     `json:"can_share_user"`
	CanWrite       bool     `json:"can_write"`
	Modifiable     bool     `json:"modifiable"`
	Owner          string   `json:"owner"`
	Perms          ACLPerms `json:"perms"`
	Removable      bool     `json:"removable"`
	Sharing        string   `json:"sharing"`
}

type Entry struct {
	Name    string      `json:"name"`
	ID      string      `json:"id"`
	Updated string      `json:"updated"`
	Links   interface{} `json:"links"`
	Author  string      `json:"author"`
	ACL     ACL         `json:"acl"`
}
