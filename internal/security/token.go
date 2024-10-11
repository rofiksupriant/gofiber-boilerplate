package security

import (
	"boilerplate/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewToken(username string, role string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   username,
		Audience:  jwt.ClaimStrings{role},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	secretKey := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47")
	signedToken, errSignToken := token.SignedString(secretKey)
	if errSignToken != nil {
		log.Errorf("error sign token %v", errSignToken)
		return nil, errSignToken
	}

	return &signedToken, nil
}

func VerifyToken(tokenPlain string) (*model.TokenClaims, bool) {
	secretKey := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47")

	// parse token
	token, err := jwt.ParseWithClaims(tokenPlain, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		log.Errorf("error parse token")
		return nil, false
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); !ok {
		log.Errorf("error get claims from token")
		return nil, false
	} else {
		subject, _ := claims.GetSubject()
		audience, _ := claims.GetAudience()
		expirationTime, _ := claims.GetExpirationTime()

		var tokenClaims = model.TokenClaims{
			Subject:   subject,
			Audience:  audience[0],
			ExpiresAt: &expirationTime.Time,
		}
		return &tokenClaims, true
	}
}
