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

	// JWTSecret := "xihaijie"
	// c := map[string]interface{}{
	// 	"foo": "bar",
	// }

	// // Sign and get the complete encoded token as a string using the secret
	// tokenString, err := jwtwrapper.IssueTokenStr(c, JWTSecret)
	// tokenString2, err := jwtwrapper.IssueTokenStrWithExp(c, JWTSecret, 5)

	// fmt.Println(tokenString, err)

	// cc, err := jwtwrapper.GetMapClaims(tokenString, JWTSecret)
	// fmt.Printf("%v,%v\n", cc, err)

	// for {
	// 	cc2, err := jwtwrapper.GetMapClaims(tokenString2, JWTSecret)
	// 	fmt.Printf("%v,%v\n", cc2, err)

	// 	time.Sleep(time.Second)
	// }

	// boltwrapper.InitDB()
	// userBytes := boltwrapper.UserDB.GetUser("Yushi")
	// fmt.Println(userBytes)

	qrcode.WriteFile("otpauth://totp/Hubble:Yushi?issuer=Hubble&secret=WTGQ3OEBTQI4O2U5", qrcode.Medium, 256, "qr.png")
}
