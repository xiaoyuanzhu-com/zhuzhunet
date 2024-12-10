package models

type BrandList struct {
	List     []*Brand  `json:"list" noco:"list"`
	PageInfo *PageInfo `json:"page_info" noco:"pageInfo"`
}
