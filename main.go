package main

import (
    _ "example/protobuf/cool"
    "github.com/astaxie/beego"
)

func main() {

    beego.BConfig.WebConfig.DirectoryIndex = true
    beego.BConfig.WebConfig.StaticDir["/static"] = "static"
    beego.Run()
    //test()

}
