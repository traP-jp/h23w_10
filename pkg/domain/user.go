package domain

import "net/url"

type User struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	IconURL     url.URL  `json:"icon_url,omitempty"`
	UserType    UserType `json:"user_type,omitempty"`
}

type UserType string

const (
	UserTypeTrap     UserType = "trap"
	UserTypeExternal UserType = "external"
)
