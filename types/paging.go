package types

type Paging struct {
	Total   *int64 `json:"total,omitempty"`
	PerPage *int64 `json:"perPage,omitempty"`
	Offset  *int64 `json:"offset,omitempty"`
}

type ListRequest struct {
	Count  *int64 `json:"count,omitempty" url:"count"`
	Offset *int64 `json:"offset,omitempty" url:"offset"`
}
