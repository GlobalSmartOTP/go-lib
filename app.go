package go_lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	SMS          = "sms"
	IVR          = "ivr"
	GapMessenger = "gap"
	Email        = "email"
)

var (
	CodeIsRequiredErr     = errors.New("code not found")
	MobileIsRequiredErr   = errors.New("mobile not found")
	InvalidExpireTimeErr  = errors.New("expire time is invalid")
	InvalidCountryCodeErr = errors.New("countryCode is invalid")
	InvalidLengthErr      = errors.New("length must >4 and <10")
	InvalidRequestErr     = errors.New("invalid request")
	InvalidReferenceID    = errors.New("referenceID not found")
)

//SendOTP use when you want to send your code to end user
type SendOTP struct {
	//Code is required.
	Code string

	//CountryCode is optional.
	CountryCode int

	//Mobile is required.
	Mobile string

	//Param1 is optional.
	Param1 string

	//Param2 is optional.
	Param2 string

	//Param3 is optional.
	Param3 string

	//if you set TemplateID to 0 then default templateID use.
	TemplateID int
}

//SendAutoOTPCode use when you don't want to generate otp code and our system generate otp code and verify it.
type SendAutoOTPCode struct {
	//CountryCode is optional.
	CountryCode int

	//Mobile is required.
	Mobile string

	//if you want your otp could be expired, set ExpireTime as second .
	//
	//for example 60 second
	ExpireTime int64

	//Param1 is optional.
	Param1 string

	//Param2 is optional.
	Param2 string

	//Param3 is optional.
	Param3 string
	//Length use when you don't want to send specific Code and our system generate code with Length you set.
	//muse be > 4 and < 10
	Length int
	//if you set TemplateID to 0 then default templateID use.
	TemplateID int
}

type SendIVRCode struct {
	//Code is required.
	Code int

	//Mobile is required.
	Mobile string

	//if you set TemplateID to 0 then default templateID use.
	TemplateID int
}

type BasicRequest struct {
	//Code is optional : send specific code
	Code string
	//CountryCode is optional
	CountryCode int
	//ExpireTime is optional

	//use when you want to send OTPCode without sending Code
	ExpireTime int64
	//Length use when you don't want to send specific Code and our system generate code with Length you set.
	//muse be > 4 and < 10
	Length int
	//Method required.
	//one of sms ivr gap email
	Method string
	//Mobile required.
	//example 9012341234 or 09012341234
	Mobile string
	//Param1 is optional.
	Param1 string
	//Param2 is optional.
	Param2 string
	//Param3 is optional.
	Param3 string
	Smart  bool
	//if you set TemplateID to 0 then default templateID use.
	TemplateID int
}

func (b *VerifyRequest) validate() error {
	if b.Mobile == "" {
		return MobileIsRequiredErr
	}
	if b.OTP == "" {
		return CodeIsRequiredErr
	}
	if b.CountryCode < 0 {
		return InvalidCountryCodeErr
	}
	return nil
}
func (b *StatusRequest) validate() error {
	if b.ReferenceID == "" {
		return InvalidReferenceID
	}
	return nil
}
func (b *SendIVRCode) validate() error {
	if b.Mobile == "" {
		return MobileIsRequiredErr
	}
	if b.Code != 0 {
		return CodeIsRequiredErr
	}
	return nil
}

func (b *SendAutoOTPCode) validate() error {
	if b.Mobile == "" {
		return MobileIsRequiredErr
	}
	if b.ExpireTime != 0 && (b.ExpireTime < 60 || b.ExpireTime > 86400) {
		return InvalidExpireTimeErr
	}
	if b.CountryCode < 0 {
		return InvalidCountryCodeErr
	}
	if b.Length < 4 || b.Length > 10 {
		return InvalidLengthErr
	}
	return nil
}
func (b *SendOTP) validate() error {
	if b.Mobile == "" {
		return MobileIsRequiredErr
	}
	if b.CountryCode < 0 {
		return InvalidCountryCodeErr
	}
	return nil
}
func (b *BasicRequest) validate() error {
	if b.Mobile == "" {
		return MobileIsRequiredErr
	}
	if b.ExpireTime < 0 {
		return InvalidExpireTimeErr
	}
	if b.ExpireTime != 0 && (b.ExpireTime < 60 || b.ExpireTime > 86400) {
		return InvalidCountryCodeErr
	}
	if b.Code == "" {
		return CodeIsRequiredErr
	}
	if b.Length < 4 || b.Length > 10 {
		return InvalidLengthErr
	}
	return nil
}

type Config struct {
	ApiKey string
}
type App struct {
	config Config
	client *http.Client
}
type Request interface {
	validate() error
}
type Response interface {
	ToString() string
}

type SendResponse struct {
	Status      string                 `json:"status"`
	Error       map[string]interface{} `json:"error"`
	ReferenceID string                 `json:"referenceID"`
}
type StatusResponse struct {
	Status      string                 `json:"status"`
	Error       map[string]interface{} `json:"error"`
	OTPStatus   string                 `json:"OTPStatus"`
	OTPVerified bool                   `json:"OTPVerified"`
	OTPMethod   string                 `json:"OTPMethod"`
}
type StatusRequest struct {
	ReferenceID string `json:"OTPReferenceID"`
}
type VerifyRequest struct {
	//OTP is code
	OTP         string `json:"otp"`
	CountryCode int    `json:"countryCode"`
	Mobile      string `json:"mobile"`
}

type VerifyResponse struct {
	Status string                 `json:"status"`
	Error  map[string]interface{} `json:"error"`
}

func (r *VerifyResponse) ToString() string {
	if r.Status == "success" {
		return "verify status: success\nsent : true"
	} else {
		return "verify status: failed\nmessage : " + r.Error["message"].(string) + "\nerror code : " + strconv.FormatFloat(r.Error["code"].(float64), 'E', -1, 64)
	}
}
func (r *StatusResponse) ToString() string {
	if r.Status == "success" {
		return "sms status: success\nOTP Status : " + r.OTPStatus + "\nOTP Verified : " + strconv.FormatBool(r.OTPVerified) + "\nOTP Method : " + r.OTPMethod
	} else {
		return "sms status: failed\nmessage : " + r.Error["message"].(string) + "\nerror code : " + strconv.FormatFloat(r.Error["code"].(float64), 'E', -1, 64)
	}
}
func (r *SendResponse) ToString() string {
	if r.Status == "success" {
		return "sending status: success\nsent : true\nreferenceID : " + r.ReferenceID
	} else {
		return "sending status: failed\nmessage : " + r.Error["message"].(string) + "\nerror code : " + strconv.FormatFloat(r.Error["code"].(float64), 'E', -1, 64)
	}
}
//Verify otp code that clients sent it to you
func (app *App) Verify(req VerifyRequest) (*VerifyResponse, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	request, err := http.NewRequest("POST", "https://api.gsOTP.com/otp/verify", buf)
	request.Header.Add("apiKey", app.config.ApiKey)
	r, err := app.client.Do(request)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	res := &VerifyResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
//GetStatus show you details of your otp
func (app *App) GetStatus(req StatusRequest) (*StatusResponse, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	request, err := http.NewRequest("POST", "https://api.gsOTP.com/otp/status", buf)
	request.Header.Add("apiKey", app.config.ApiKey)
	r, err := app.client.Do(request)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	res := &StatusResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
//Send a new otp.
//
//type: SendAutoOTPCode:gsOTP generate otp code.you only set length of code
//
//type: SendOTP:send manual otp code
//
//type: SendIVRCode
//
//or use basic request BasicRequest
func (app *App) Send(req Request) (*SendResponse, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}
	var basicReq BasicRequest

	switch req.(type) {
	case *SendAutoOTPCode:
		var r = req.(*SendAutoOTPCode)
		basicReq.CountryCode = r.CountryCode
		basicReq.TemplateID = r.TemplateID
		basicReq.Length = r.Length
		basicReq.ExpireTime = r.ExpireTime
		basicReq.Param1 = r.Param1
		basicReq.Param2 = r.Param2
		basicReq.Param3 = r.Param3
		basicReq.Mobile = r.Mobile
		basicReq.Method = "sms"
		break
	case *SendOTP:
		var r = req.(*SendOTP)
		basicReq.CountryCode = r.CountryCode
		basicReq.TemplateID = r.TemplateID
		basicReq.Param1 = r.Param1
		basicReq.Param2 = r.Param2
		basicReq.Param3 = r.Param3
		basicReq.Mobile = r.Mobile
		basicReq.Method = "sms"
		basicReq.Code = r.Code
		break
	case *SendIVRCode:
		var r = req.(*SendIVRCode)
		basicReq.TemplateID = r.TemplateID
		basicReq.Mobile = r.Mobile
		basicReq.Method = "ivr"
		basicReq.Code = strconv.Itoa(r.Code)
		break
	case *BasicRequest:
		basicReq = *req.(*BasicRequest)
		break
	default:
		return nil, InvalidRequestErr
	}
	buf := &bytes.Buffer{}
	data, err := json.Marshal(basicReq)
	if err != nil {
		return nil, err
	}
	buf.Write(data)
	request, err := http.NewRequest("POST", "https://api.gsOTP.com/otp/send", buf)
	request.Header.Add("apiKey", app.config.ApiKey)
	r, err := app.client.Do(request)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	res := &SendResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
//New create new client .
func New(c Config) *App {
	if c.ApiKey == "" {
		panic("apiKey not set")
	}
	app := &App{config: c}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 5
	t.MaxConnsPerHost = 5
	t.MaxIdleConnsPerHost = 5

	app.client = &http.Client{
		Timeout:   5 * time.Second,
		Transport: t,
	}
	return app
}
func (app *App) SetConfig(c Config) (err error) {
	if c.ApiKey == "" {
		return errors.New("apiKey not set")
	}
	app.config = c
	return
}
func (app *App) GetCurrentApiKey() string {
	return app.config.ApiKey
}
