package entity

type Pagination struct{
	PageSizes []int `json:"page_sizes"`
	PageSize int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	Total int64 `json:"total"`
}
