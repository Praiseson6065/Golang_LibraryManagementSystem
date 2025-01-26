package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"fmt"
)

var (
	jwtExpiration int
	signingKey    []byte
)

type JWTClaims struct {
	UserId string `json:"userId"`
	RoleId uint `json:"roleId"`
	jwt.StandardClaims
}

func init() {
	jwtExpiration = viper.GetInt("JWT.EXPIRE")
	signingKey = []byte(viper.GetString("JWT.PRIVATE_KEY"))
}
func GenerateToken(userId string, roleId uint) (string, error) {

	claims := JWTClaims{
		userId,
		roleId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(jwtExpiration) * time.Minute).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Issuer:    "Librarian-PS",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err

}


func ValidateToken(encodedToken string) (string,uint,error){
	claims:=&JWTClaims{}

	_,	err := jwt.ParseWithClaims(encodedToken,claims,func(t *jwt.Token) (interface{}, error) {
		if _ , isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s",t.Header["alg"])
		}
		return []byte(signingKey),nil
	})
	if err!=nil{
		return "",999,err
	}
	return claims.UserId,claims.RoleId,nil
}