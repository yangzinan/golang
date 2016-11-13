package main

import (
	"github.com/astaxie/beego"
)

type MyController struct {
	beego.Controller
}

func (self *MyController) Get() {
	self.Ctx.WriteString("Hello,Golang! 你好，Go语言！")
}

func main() {
	beego.Router("/", &MyController{})
	beego.Run()
}
