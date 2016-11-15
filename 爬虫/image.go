package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	GetCodeImage()
}

func GetCodeImage() []*http.Cookie {
	url := "https://biz.kaslyju.net/kasly-web/kaptchaImage"
	response, err := http.Get(url)
	image := response.Body
	if err != nil {
		panic(err)
	}

	file, r := os.Create("test.jpg")
	if r != nil {
		panic(r)
	}
	io.Copy(file, image)
	return response.Cookies()
}
