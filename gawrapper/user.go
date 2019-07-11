package gawrapper

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateSecret() string {
	b := make([]byte, 10)
	rand.Read(b)
	return base32.StdEncoding.EncodeToString(b)
}

func NewOTPAuth(name, secret, issuer string) string {
	otpc := &OTPConfig{
		Secret: secret,
	}
	return otpc.ProvisionURIWithIssuer(name, issuer)
}
