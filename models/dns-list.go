package models

type DNSList struct {
	List     []*DNS    `json:"list" noco:"list"`
	PageInfo *PageInfo `json:"page_info" noco:"pageInfo"`
}
