package go_lib_test

import (
	go_lib "github.com/GlobalSmartOTP/go-lib"
	"os"
	"testing"
)

var Token string

func init() {
	err := os.Setenv("TOKEN", "bb576c7cf9aeea304eaaa6c5be9dfc0c")
	if err != nil {
		return
	}
	Token = os.Getenv("TOKEN")

}

func TestApp_Send1(t *testing.T) {
	app := go_lib.New(go_lib.Config{ApiKey: Token})
	res, err := app.Send(&go_lib.SendAutoSMSCode{
		CountryCode: 0,
		Mobile:      "9016574449",
		ExpireTime:  0,
		Param1:      "",
		Param2:      "",
		Param3:      "",
		Length:      6,
		TemplateID:  3,
	})
	if err != nil {
		t.Errorf("sending sms failed %e", err)
		return
	}
	if res.ReferenceID == 0 {
		t.Errorf("sending sms failed:" + res.ToString())
	}

}

func TestApp_Send2(t *testing.T) {
	app := go_lib.New(go_lib.Config{ApiKey: Token})
	_, err := app.Send(&go_lib.SendAutoSMSCode{
		CountryCode: 0,
		Mobile:      "9016574449",
		ExpireTime:  0,
		Param1:      "",
		Param2:      "",
		Param3:      "",
		Length:      18,
		TemplateID:  3,
	})
	if err == nil {
		t.Errorf("length check ignored%e", err)
		return
	}

}

func TestApp_Verify(t *testing.T) {
	app := go_lib.New(go_lib.Config{ApiKey: Token})
	res, err := app.Verify(go_lib.VerifyRequest{
		CountryCode: 0,
		Mobile:      "9016574449",
		OTP:         "xxx",
	})
	if err != nil {
		t.Errorf("verify error accurred%e", err)
		return
	}
	if res.Status != "error" {
		t.Errorf("verify not failed")
		t.Log(res)
		return
	}

}

func TestApp_GetStatus1(t *testing.T) {
	app := go_lib.New(go_lib.Config{ApiKey: Token})
	res, err := app.GetStatus(go_lib.StatusRequest{
		ReferenceID: 234,
	})
	if err != nil {
		t.Errorf("verify error accurred%e", err)
		return
	}
	if res.Status != "error" {
		t.Errorf("verify not failed")
		t.Log(res)
		return
	}

}

func TestApp_GetStatus2(t *testing.T) {
	app := go_lib.New(go_lib.Config{ApiKey: Token})
	res, err := app.GetStatus(go_lib.StatusRequest{
		ReferenceID: 1625640786533816718,
	})
	if err != nil {
		t.Errorf("verify error accurred%e", err)
		return
	}
	if res.Status != "success" {
		t.Errorf("verify not failed")
		t.Log(res)
		return
	}
	t.Log(res.ToString())
}
