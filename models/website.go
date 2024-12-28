package models

type Website struct {
	ID      int    `json:"id" noco:"Id"`
	Address string `json:"address" noco:"address"`
	Brands  struct {
		ID   int    `json:"id" noco:"Id"`
		Name string `json:"name" noco:"name"`
	} `json:"brands" noco:"brands"`
}
