package jwt

import (

)

type Claims interface {
	Valid() error
}

type StandardClaims struct {
	Audience string    `json:"aud,omitempty"`
	ExpiresAt string    `json:"exp,omitempty"`
	Id       string    `json:"jti,omitempty"`
	IssuedAt int64     `json:"iat,omitempty"`
	Issuer   string    `json:"iss,omitempty"`
	NoBefore  int64    `json:"nbf,omitempty"`
	Subject  string    `json:"sub,omitempty"`
}

