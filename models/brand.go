package models

type Brand struct {
	ID   int    `json:"id" noco:"Id"`
	Name string `json:"name" noco:"name"`
	Logo []struct {
		SignedURL string `json:"signed_url" noco:"signedUrl"`
	} `json:"logo" noco:"logo"`
}
