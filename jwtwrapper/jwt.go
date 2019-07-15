package jwtwrapper

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetMapClaims(tokenStr, secretStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretStr), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if token.Valid {
			err := claims.Valid()
			return claims, err
		}
		return nil, fmt.Errorf("Invalid token signature")
	}
	return nil, fmt.Errorf("Unexpected claims format: %T", claims)
}

func IssueTokenStr(claims map[string]interface{}, secretStr string) (string, error) {
	var claimsMC jwt.MapClaims
	claimsMC = claims
	claimsMC["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMC)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secretStr))
}

func IssueTokenStrWithExp(claims map[string]interface{}, secretStr string, second int64) (string, error) {
	var claimsMC jwt.MapClaims
	claimsMC = claims
	now := time.Now().Unix()
	claimsMC["iat"] = now
	claimsMC["exp"] = now + second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMC)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secretStr))
}
