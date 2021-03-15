package jwt

import (
	"fmt"
	"time"

	"github.com/Kourin1996/go-crud-api-sample/api/constants"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
	"github.com/dgrijalva/jwt-go"
)

type TokenData struct {
	HashId string
	Exp    int64
}

func EncodeJWTToken(u *user.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["hash_id"] = u.HashId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(constants.JWT_SECRET))
}

func DecodeJWTToken(token *jwt.Token) *TokenData {
	fmt.Printf("DecodeJWTToken %+v\n", token)
	claims := token.Claims.(jwt.MapClaims)
	hashId := claims["hash_id"].(string)

	return &TokenData{HashId: hashId}
}
