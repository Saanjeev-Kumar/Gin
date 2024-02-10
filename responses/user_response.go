package responses

type UserResponse struct {
	Data    map[string]interface{} `json:"data"`
	// Name     string             `json:"name,omitempty" `
	// Email    string             `json:"location,omitempty" `
}