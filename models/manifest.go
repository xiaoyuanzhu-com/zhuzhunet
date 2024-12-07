package models

type Manifest struct {
	SessionID string       `json:"session_id" bson:"session_id"`
	DNS       *ManifestDNS `json:"dns" bson:"dns"`
}
