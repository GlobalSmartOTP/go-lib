package main

import (
	"fmt"
	go_lib "github.com/GlobalSmartOTP/go-lib"
)

func main()  {

	app := go_lib.New(go_lib.Config{ApiKey: ""})
	res, err := app.Send(&go_lib.SendAutoOTPCode{
		CountryCode: 0,
		Mobile:      "9123456789",
		ExpireTime:  0,//optional: otp expire time in second
		Param1:      "",//optional
		Param2:      "",//optional
		Param3:      "",//optional
		Length:      6,//optional: length of your otp code-> example with length 6: 123456
		TemplateID:  3,//template that you want to send otp with that scheme
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("otp referenceID:",res.ReferenceID)

}





