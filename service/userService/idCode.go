package userService

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var digitDriver = &base64Captcha.DriverString{
	Height:          32,
	Width:           100,
	NoiseCount:      0,
	ShowLineOptions: 2,
	Length:          5,
	Source:          "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

var store = base64Captcha.DefaultMemStore

// CaptchaGenerate 生成验证码
func CaptchaGenerate() (string, string, string, error) {
	b := base64Captcha.NewCaptcha(digitDriver, store)
	id, b64s, _, err := b.Generate()
	hcode := store.Get(id, false)
	if err != nil {
		fmt.Println("Error generating captcha:", err)
		return "", "", "", err
	}
	return id, b64s, hcode, nil
}

// GetCodeAnswer 验证验证码
func GetCodeAnswer(id, code string) bool {
	return store.Verify(id, code, true)
}
