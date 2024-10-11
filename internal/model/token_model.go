package model

import (
	"time"
)

type TokenClaims struct {
	Subject   string     `json:"sub,omitempty"`
	Audience  string     `json:"aud,omitempty"`
	ExpiresAt *time.Time `json:"exp,omitempty"`
}
