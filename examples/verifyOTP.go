package main

import (
	"fmt"
	go_lib "github.com/GlobalSmartOTP/go-lib"
)

func main()  {
	app := go_lib.New(go_lib.Config{ApiKey: ""})
	res, err := app.Verify(go_lib.VerifyRequest{
		OTP:         "",//otp code that we sent it to mobile
		CountryCode: 0,//optional
		Mobile:      "9123456789",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("is verified?:",res.Status)
}
