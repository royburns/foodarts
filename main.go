package main

import (
	"github.com/astaxie/beego"
	_ "github.com/royburns/foodarts/models"
	_ "github.com/royburns/foodarts/routers"
)

func main() {
	beego.Run()
}
