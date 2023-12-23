package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type UserClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

func getSecret() []byte {
	return []byte(viper.Get("TOKEN_SECRET").(string))
}

func newUserClaims(username string) *UserClaims {
	return &UserClaims{Username: username}
}

func GenerateJwtToken(username string) (string, error) {
	claims := newUserClaims(username)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := getSecret()

	return token.SignedString(secret)
}

func ParseToken(token string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := getSecret()
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}
