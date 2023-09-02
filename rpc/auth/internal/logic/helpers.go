package logic

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type ParseClaim struct {
	UserId int64
	jwt.Claims
}

type CreateClaim struct {
	UserId int64
	jwt.RegisteredClaims
}

const (
	Salt       = "mju71qazmju72wsx5tgbasdavzxcqwwfcsdc23rqwdsacnhy6ki87ujm0p"
	SigningKey = "qwhkasdlkbwjdoasbjbkwkafcsdc23rqwdsacnhy6ki87ujm0p"
)

func generatePasswordHash(password *string) {
	// convert password string to slice of bytes
	passwordBytes := []byte(*password)

	// creating sha-512 header
	sha512Header := sha512.New()

	// convert salt string to slice of bytes
	saltBytes := []byte(Salt)

	// Append salt to password
	passwordBytes = append(passwordBytes, saltBytes...)

	// write password to sha-512 header
	if _, err := sha512Header.Write(passwordBytes); err != nil {
		log.Println(err)
		return
	}

	// get sha-512 hashed password
	hashedPassword := sha512Header.Sum(nil)

	// convert hashed password to HEX string
	*password = hex.EncodeToString(hashedPassword)

	return
}

func generateToken(id int64) string {
	claims := CreateClaim{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultToken, err := token.SignedString([]byte(SigningKey))
	if err != nil {
		log.Println(err)
	}
	return resultToken
}

func parseToken(accessToken string) (int64, error) {
	var claim ParseClaim
	token, err := jwt.ParseWithClaims(accessToken, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	switch {
	case !token.Valid:
		return 0, fmt.Errorf("invalid token")
	case claim.UserId <= 0:
		return 0, fmt.Errorf("user not found")
	}

	return claim.UserId, nil
}
