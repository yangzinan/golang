package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

func main() {
	response, _ := http.Get("https://biz.kaslyju.net") //获取http返回
	body, _ := ioutil.ReadAll(response.Body)           //获取返回体
	context := string(body)                            //返回主题转化为字符串
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/"}
	cookies := response.Cookies()
	cookies = append(cookies, &cookie)
	fmt.Println(cookies)
	//fmt.Println(context)
	fmt.Println("-------------------------------------------------------")
	root, err := html.Parse(strings.NewReader(context)) //初始化xpath文本
	if err != nil {
		panic(err)
	}
	node := htmlquery.Find(root, "//*[@id=\"loginForm\"]") //执行xpanth查找
	fmt.Println(node[0].Attr[2].Val)                       //获取当前标签的第二属性（class）的值
	fmt.Println(len(node))                                 //获取查找到的个数
	fmt.Println(htmlquery.InnerText(node[0]))              //获取标签的文本（当前没有向下获取最近的一个）

}
