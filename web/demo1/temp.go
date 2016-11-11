package main

import (
	"github.com/astaxie/beego"
)

type MyController struct {
	beego.Controller
}

func (self *MyController) Get() {
	self.Ctx.WriteString("Hello,World! 你好，世界！")
}

func main() {
	beego.Router("/", &MyController{})
	beego.Run()
}
