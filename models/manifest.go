package models

type Manifest struct {
	DNS *ManifestDNS `json:"dns" bson:"dns"`
}
