package models

type WebsiteList struct {
	List     []*Website `json:"list" noco:"list"`
	PageInfo *PageInfo  `json:"page_info" noco:"pageInfo"`
}
