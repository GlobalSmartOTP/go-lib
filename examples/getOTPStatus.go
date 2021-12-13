package main

import (
	"fmt"
	go_lib "github.com/GlobalSmartOTP/go-lib"
)

func main() {
	app := go_lib.New(go_lib.Config{ApiKey: "my_api_key"})
	res, err := app.GetStatus(go_lib.StatusRequest{ReferenceID: ""}) //referenceID that you got in sendOTP method
	if err != nil {
		panic(err)
	}
	fmt.Println("OTP status", res.Status)
	fmt.Println("OTP is verified?", res.OTPVerified)
	fmt.Println("full details", res)
}
