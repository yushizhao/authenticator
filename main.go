package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
	"github.com/yushizhao/authenticator/gawrapper"
)

func main() {
	// secret := gawrapper.GenerateSecret()
	secret := "V6XBUJT5SRL57YIU"
	name := "lvjidong"
	issuer := "hubble"

	fmt.Println(secret)
	otpc := &gawrapper.OTPConfig{
		Secret:      secret,
		WindowSize:  0,
		HotpCounter: 0,
	}

	fmt.Println(gawrapper.ComputeTOTP(otpc.Secret))

	u := otpc.ProvisionURIWithIssuer(name, issuer)

	fmt.Println(u)

	err := qrcode.WriteFile(u, qrcode.Medium, 256, "qr.png")

	if err != nil {
		fmt.Println(u)
	}
}
