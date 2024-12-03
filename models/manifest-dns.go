package models

type ManifestDNS struct {
	Servers []*DNSServer `json:"servers" bson:"servers"`
}
