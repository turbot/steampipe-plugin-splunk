package types

type TokenListResponse struct {
	Entry  []Token `json:"entry"`
	Paging Paging  `json:"paging"`
}

type Token struct {
	Entry
	Content TokenContent `json:"content"`
}

type TokenContent struct {
	Claims     TokenClaims `json:"claims"`
	EAIACL     ACL         `json:"eai:acl"`
	Headers    interface{} `json:"headers"`
	LastUsed   int64       `json:"lastUsed"`
	LastUsedIP string      `json:"lastUsedIp"`
	Status     string      `json:"status"`
}

type TokenClaims struct {
	Audience         string   `json:"aud,omitempty"`
	Subject          string   `json:"sub,omitempty"`
	IssuedAt         int64    `json:"iat,omitempty"`
	NotBefore        int64    `json:"nbr,omitempty"`
	ExpirationTime   int64    `json:"exp,omitempty"`
	Issuer           string   `json:"iss,omitempty"`
	IdentityProvider string   `json:"idp,omitempty"`
	Roles            []string `json:"roles,omitempty"`
}
