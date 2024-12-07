package models

type Brand struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Logo string `json:"logo" bson:"logo"`
}
