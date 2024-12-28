package models

type IPInfo struct {
	IP        string  `json:"ip"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ASN       int     `json:"asn"`
	AS        string  `json:"as"`
}
