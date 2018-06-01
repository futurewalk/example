package protobuf_cool

import "github.com/futurewalk/doc-cool/cool"
import "example/controllers"
import "example/protobuf"


func init() {
    container := make(map[string]interface{})
    container["ExampleController"]= &controllers.ExampleController{}
    container["ExampleMessageReq"]= &protobuf.ExampleMessageReq{}
    container["ExampleMessageRsp"]= &protobuf.ExampleMessageRsp{}
    container["BytesMessageReq"]= &protobuf.BytesMessageReq{}
    container["BytesMessageRsp"]= &protobuf.BytesMessageRsp{}
    container["Extension"] = &Resolver{} 
    cool.Start(container)
}

type Resolver struct{}

func (p *Resolver) Invoke(plugin *cool.Plugin) {
    if plugin.StructName == "BytesMessageReq" && plugin.FieldName == "Data" && plugin.Url == "/v1/example/BytesSupport"{
        plugin.Swap(&protobuf.BytesMessageInfo{})
    }
}
