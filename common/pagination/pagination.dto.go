package pagination

type Pagination struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	Total    int `json:"total"`
}
