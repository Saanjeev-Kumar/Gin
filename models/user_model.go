package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string             `json:"name,omitempty" `
	Email    string             `json:"location,omitempty" `
}
