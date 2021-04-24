package auth0

import "github.com/form3tech-oss/jwt-go"

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKS struct {
	Keys []JSONWebKeys `json:"keys"`
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}
