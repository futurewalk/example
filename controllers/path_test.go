package controllers

import (
    "testing"
    "log"
    "doc-cool/cool"
)

func TestPath(t *testing.T) {
    cool.Export()

   /* d := demo
    fv := reflect.ValueOf(d)
    fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
    fv.Call(nil)*/

}
func demo() {
    log.Println("获取当前的犯法.......")
}

/*if plugin.StructName == "BytesMessageReq" && plugin.FieldName == "Data" && plugin.Url == "/v1/example/BytesSupport" {
    plugin.Swap(&protobuf.InfoMessage{})
}*/
