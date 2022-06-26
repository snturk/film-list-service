package models

type Film struct {
	Name string 					`bson:"name" json:"name"`
	Description string 				`bson:"description" json:"description"`
	Date string 					`bson:"date" json:"date"`
}
