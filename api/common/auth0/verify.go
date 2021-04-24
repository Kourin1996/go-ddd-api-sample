package auth0

import "crypto/subtle"

func VerifyAuthorizedParties(azp string, clientIDs []string, required bool) bool {
	if len(azp) == 0 {
		return !required
	}
	for _, b := range clientIDs {
		if subtle.ConstantTimeCompare([]byte(azp), []byte(b)) != 0 {
			return true
		}
	}
	return false
}
