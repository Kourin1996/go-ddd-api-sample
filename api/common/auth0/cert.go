package auth0

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
)

func FetchJWKS(auth0Domain string) (*JWKS, error) {
	url := fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jwks := &JWKS{}
	err = json.NewDecoder(resp.Body).Decode(jwks)

	return jwks, err
}

func GetPemCert(jwks *JWKS, token *jwt.Token) (string, error) {
	cert := ""

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
