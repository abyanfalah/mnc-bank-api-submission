package authenticator

import (
	"fmt"
	"mnc-bank-api/config"
	"mnc-bank-api/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type accessToken struct {
	config config.TokenConfig
}

type AccessToken interface {
	GenerateAccessToken(customer *model.Customer) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

func (at *accessToken) GenerateAccessToken(customer *model.Customer) (string, error) {
	now := time.Now().UTC()
	end := now.Add(at.config.AccessTokenLifetime)

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    at.config.ApplicationName,
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
		Id:       customer.Id,
		Username: customer.Username,
	}

	token := jwt.NewWithClaims(
		at.config.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(at.config.JwtSignatureKey))
}

func (at *accessToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("method not ok")
		} else if method != at.config.JwtSigningMethod {
			return nil, fmt.Errorf("signing method is different from config")
		}

		return []byte(at.config.JwtSignatureKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		config: config,
	}
}
