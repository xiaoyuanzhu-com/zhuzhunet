package models

type PageInfo struct {
	IsFirstPage bool `json:"is_first_page" noco:"isFirstPage"`
	IsLastPage  bool `json:"is_last_page" noco:"isLastPage"`
	Page        int  `json:"page" noco:"page"`
	PageSize    int  `json:"page_size" noco:"pageSize"`
	TotalRows   int  `json:"total_rows" noco:"totalRows"`
}
