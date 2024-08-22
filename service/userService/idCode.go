package userService

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"strings"
)

var digitDriver = &base64Captcha.DriverString{
	Height:          90,
	Width:           240,
	NoiseCount:      0,
	ShowLineOptions: 2,
	Length:          4,
	Source:          "abcdefghijklmnopqrstuvwxyz",
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
	return store.Verify(id, strings.ToLower(code), true)
}
