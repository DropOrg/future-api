package models

import "hash"

// User represents a user
type User struct {
	Username string    `json:"id,omitempty"`
	Pass     hash.Hash `json:"omitempty"`
}
