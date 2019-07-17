package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	// secret := gawrapper.GenerateSecret()
	// secret := "V6XBUJT5SRL57YIU"
	// name := "lvjidong"
	// issuer := "hubble"

	// fmt.Println(secret)
	// fmt.Println(gawrapper.ComputeTOTP(secret))
	// u := gawrapper.NewOTPAuth(name, secret, issuer)
	// fmt.Println(u)

	// var err error
	// err = qrcode.WriteFile(u, qrcode.Medium, 256, "qr.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var codeInput string
	// for {
	// 	fmt.Scanln(&codeInput)

	// 	ok, err := gawrapper.VerifyTOTP(secret, codeInput)
	// 	if err != nil {
	// 		fmt.Println(u)
	// 	}
	// 	fmt.Println(ok)
	// }

	// boltwrapper.InitDB()
	// userBytes := boltwrapper.UserDB.GetUser("Yushi")
	// fmt.Println(userBytes)

	qrcode.WriteFile("otpauth://totp/Hubble:Jidong?issuer=Hubble&secret=OWWC6R2S4LDDTE6P", qrcode.Medium, 256, "qr.png")
}
