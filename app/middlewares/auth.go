package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
    ID          uint       
    Email     	string        
	Name      	string     
    jwt.StandardClaims
}

type ConfigJWT struct {
    SecretJWT string
    ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig{
	return middleware.JWTConfig{
		Claims: &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}



func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint, email string, name string) (string, error) {
    claims := JWTCustomClaims{
        id,
        email,
        name,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour*3).Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
        },
    }
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := t.SignedString([]byte(jwtConf.SecretJWT))
    return token, err
}
