package models

type AuthResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

var ValidProviders = map[string]bool{
	"local":  true,
	"google": true,
}
