package auth0

import (
	"errors"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func NewMiddleware(clientIDs []string, domain string, jwks *JWKS) (*jwtmiddleware.JWTMiddleware, error) {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: NewValidationKeyGetter(clientIDs, domain, jwks),
		SigningMethod:       jwt.SigningMethodRS256,
		// to prevent validator from writing to response
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {},
	}), nil
}

func NewValidationKeyGetter(clientIDs []string, domain string, jwks *JWKS) func(*jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return token, errors.New("invalid claims type")
		}

		// Verify 'azp' claim
		azp, ok := claims["azp"].(string)
		if !ok {
			return token, errors.New("authorized parties are required")
		}
		checkAzp := VerifyAuthorizedParties(azp, clientIDs, true)
		if !checkAzp {
			return token, errors.New("invalid authorized parties")
		}

		// Verify 'iss' claim
		iss := fmt.Sprintf("https://%s/", domain)
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := GetPemCert(jwks, token)
		if err != nil {
			return token, err
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	}
}
