# go-lib
<p dir=rtl>
 توسعه سامان فعالیت خود را بصورت تخصصی از سال 1381 در زمینه تجارت الکترونیک و فناوری اطلاعات آغاز کرد. فعالیت‌های شرکت در سه مسیر کلی تولید بازی‌های رایانه‌ای، تولید سیستم‌های تحت وب و ارائه زیرساخت‌های تخصصی تحت وب و سرویس‌های کاربردی اینترنت می‌باشد.</p>
<p dir=rtl>
در حوزه تولید نرم افزار اس ام اس، شرکت با توجه به رویکرد بین المللی و با تکیه بر تجربیات خود در بازارهای جهانی از سال 1384 وارد عرصه تولید نرم افزار smspanel شد و هم اکنون به بیش از 10 هزار کاربر خدمات پیامکی ارائه می‌دهد. </p>

## Usage and documentation
Please see https://doc.gsotp.com/#/gsOTP/post_otp_send for detailed usage docs.

## Installation
Use go get.
```
go get github.com/GlobalSmartOTP/go-lib
```
Then import the go-lib package into your own code.
```
import "github.com/GlobalSmartOTP/go-lib"
```
## Sample Usage
```go
package main

import (
	"fmt"
	go_lib "github.com/GlobalSmartOTP/go-lib"
	"time"
)

func main() {
	app := go_lib.New(go_lib.Config{ApiKey: "my_api_key"})
	res, err := app.Send(&go_lib.SendAutoSMSCode{
		CountryCode: 0,
		Mobile:      "9123456789",
		ExpireTime:  0,
		Param1:      "",
		Param2:      "",
		Param3:      "",
		Length:      6,
		TemplateID:  3,
	})
	if err != nil {
		return
	}
	fmt.Println(res.ReferenceID)
	time.Sleep(2 * time.Second)
	r, err := app.GetStatus(go_lib.StatusRequest{ReferenceID: res.ReferenceID})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Status)
}



```
### ApiKey
See https://gsotp.com for getting apiKey


<div dir=rtl>
دفتر مشهد:

آزادشهر، بلوار استقلال، استقلال 9، نبش چهار راه چهارم، پلاک 28

تماس:
36161 (051)
  
دفتر تهران:

خیابان شریعتی، خیابان خواجه عبدالله انصاری، کوچه چهاردهم (شهید کرمی ابرده)، پلاک۲۴، طبقه۵، واحد۱۳

تماس:
   22849646 (021)
   26713051 (021)

رایانامه: info@gsotp.com
</div> 
