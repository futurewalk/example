package main

import (
    "github.com/astaxie/beego"
    //_"example/protobuf/cool"
)

func main() {

    beego.BConfig.WebConfig.DirectoryIndex = true
    beego.BConfig.WebConfig.StaticDir["/cooler"] = "static"
    beego.Run()
}
