package protobuf_cool
import "github.com/futurewalk/doc-cool/cool"
import "example/controllers"
import "example/protobuf"


func init() {
    cool.Register(
        &Resolver{},
        &controllers.ExampleController{},
        &protobuf.ExampleMessageReq{},
        &protobuf.ExampleMessageRsp{},
        &protobuf.BytesMessageReq{},
        &protobuf.BytesMessageRsp{},
        &protobuf.RepeatMessageReq{},
        &protobuf.RepeatMessageResp{},
    )
}

type Resolver struct{}

func (p *Resolver) Invoke(plugin *cool.Plugin) {

    if plugin.StructureName == "RepeatMessageReq" && plugin.FieldName == "Data" && plugin.Url == "/v1/example/RepeatSupport" {
        plugin.Swap(&protobuf.RepeatData{})
    }

    if plugin.StructureName == "BytesMessageReq" && plugin.FieldName == "Data" && plugin.Url == "/v1/example/BytesSupport"{
        plugin.Swap(&protobuf.BytesMessageInfo{})
    }
}
