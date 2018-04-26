package main

import (
    _ "github.com/bigbluebutton/bbb-api-meetings/docs"
    _ "github.com/bigbluebutton/bbb-api-meetings/routers"
    "github.com/bigbluebutton/bbb-api-meetings/controllers"

    "github.com/astaxie/beego"
)

func main() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Run()
}
