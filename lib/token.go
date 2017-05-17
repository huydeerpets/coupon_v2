package lib

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateToken from profile of user
func GenerateToken(email string) (string, error) {
	// Create JWT token

	claims := jwt.StandardClaims{
		Audience:  email,
		ExpiresAt: time.Now().Unix() + int64(time.Duration(time.Second*24*3600)),
		Issuer:    "kowon",
		Subject:   "kowon API",
	}
	// Expire in 1d
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := tokenJwt.SignedString([]byte(beego.AppConfig.String("token_secret")))
	if err != nil {
		return "", err
	}
	beego.Info("Token: %s", tokenString)
	return tokenString, nil

}

//ParseToken token from header of client
func ParseToken(tokenstring string) (email string, err error) {
	claim := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenstring, &claim, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(beego.AppConfig.String("token_secret")), nil
	})
	if token.Valid {
		return claim.Audience, nil
	}
	return "", err
}
