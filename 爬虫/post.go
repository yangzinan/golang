package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

const (
	IMAGE_URL = "https://biz.kaslyju.net/kasly-web/kaptchaImage"
	LOGIN_URL = "https://biz.kaslyju.net/kasly-web/user/login"
	HOME_URL  = "http://biz.kaslyju.net/kasly-web/home/homepage"
)

func main() {
	cookies := GetCodeImage()
	var kcode string
	fmt.Println("请输入验证码：")
	fmt.Scanln(&kcode)
	Login(cookies, "wangdongyu", "Jly120612.", kcode)

}

func GetCodeImage() (cookies string) {
	response, err := http.Get(IMAGE_URL)
	image := response.Body
	if err != nil {
		panic(err)
	}

	file, r := os.Create("test.jpg")
	if r != nil {
		panic(r)
	}
	io.Copy(file, image)
	file.Close()
	cookies = response.Cookies()[0].Name + "=" + response.Cookies()[0].Value
	return cookies
}

func Login(cookies, username, password, keycode string) {
	client := &http.Client{}
	data := "kaptchaCode=" + keycode + "&username=" + username + "&password=" + password
	req, _ := http.NewRequest("POST", LOGIN_URL, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookies)
	client.Do(req)

	req, _ = http.NewRequest("GET", HOME_URL, strings.NewReader(""))
	req.Header.Set("Cookie", cookies)
	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	context := string(body)
	root, err := html.Parse(strings.NewReader(context))
	if err != nil {
		panic(err)
	}
	node := htmlquery.Find(root, "//*[@id=\"navigation\"]/div/div/ul/li[1]/a/span")
	if len(node) == 0 {
		fmt.Println("登录失败")
	} else {
		fmt.Printf("当前用户为：%s", htmlquery.InnerText(node[0]))
	}

}
