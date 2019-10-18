package main

import (
	"app/newsapi/newsapi/controllers"
	_ "app/newsapi/newsapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/news/", &controllers.NewsController{}, "*:NewsRead")
	beego.Run()
}
