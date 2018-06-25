package controllers

import (
    "github.com/astaxie/beego"
    "log"
    "github.com/golang/protobuf/proto"
    "example/protobuf"
    "doc-cool/cool"
)

type ExampleController struct {
    beego.Controller
}

//@ReqData ExampleMessageReq
//@RespData ExampleMessageRsp
//@ReqProtoFile example.proto
//@Method *:DocCool
//@Url /v1/example/mytest
//@Description 支持bytes数据类型方式1
func (p *ExampleController) DocCool1() {
    reqData := &protobuf.ExampleMessageReq{}
    rspData := &protobuf.ExampleMessageRsp{
        Data: &protobuf.InfoMessage{},
    }
    log.Println("铝合金和房间看电视金黄色即可佛挡多撒多撒多杀佛第三方")
    log.Println("我要测试都能太链接............................")
    err := proto.Unmarshal(p.Ctx.Input.RequestBody, reqData)
    if err == nil {
        data := reqData.Data
        rspData.Data.Method = data.Method
        rspData.Data.Url = data.Url
    }
    rspData.DocCoolName = reqData.DocCoolName
    byte, err := proto.Marshal(rspData)
    cool.Access(p.Ctx)
    p.Ctx.ResponseWriter.Write(byte)
}

//@ReqData ExampleMessageReq
//@RespData ExampleMessageRsp
//@ReqProtoFile example.proto
//@Method *:DocCool
//@Url /v1/example/DocCool
//@Description 支持bytes数据类型方式1
func (p *ExampleController) DocCool() {
    reqData := &protobuf.ExampleMessageReq{}
    rspData := &protobuf.ExampleMessageRsp{
        Data: &protobuf.InfoMessage{},
    }
    log.Println("铝合金和房间看电视金黄色即可佛挡多撒多撒多杀佛第三方")
    log.Println("我要测试都能太链接............................")
    err := proto.Unmarshal(p.Ctx.Input.RequestBody, reqData)
    if err == nil {
        data := reqData.Data
        rspData.Data.Method = data.Method
        rspData.Data.Url = data.Url
    }
    rspData.DocCoolName = reqData.DocCoolName
    byte, err := proto.Marshal(rspData)
    cool.Access(p.Ctx)
    p.Ctx.ResponseWriter.Write(byte)
}

//@ReqData BytesMessageReq
//@RespData BytesMessageRsp
//@ReqProtoFile example.proto
//@Method *:BytesSupport
//@Url /v1/example/BytesSupport
//@Description 支持bytes数据类型
func (p *ExampleController) BytesSupport() {
    var b []byte
    reqData := &protobuf.BytesMessageReq{}
    rspData := &protobuf.BytesMessageRsp{
        Data: b,
    }

    err := proto.Unmarshal(p.Ctx.Input.RequestBody, reqData)
    log.Println("获取请求数据:", reqData)
    if err == nil {
        data := reqData.Data
        rspData.Data = data
    }
    rspData.BytesName = reqData.DocCoolName
    byte, err := proto.Marshal(rspData)
    cool.Access(p.Ctx)
    p.Ctx.ResponseWriter.Write(byte)
}

//@ReqData RepeatMessageReq
//@RespData RepeatMessageResp
//@ReqProtoFile example.proto
//@Method *:RepeatSupport
//@Url /v1/example/RepeatSupport
//@Description 支持bytes数据类型
func (p ExampleController) RepeatSupport() {
    reqData := &protobuf.RepeatMessageReq{}
    rspData := &protobuf.RepeatMessageResp{}

    err := proto.Unmarshal(p.Ctx.Input.RequestBody, reqData)
    if err == nil {
        data := reqData.Data
        rspData.Data = data
        rspData.Name = reqData.Name
        rspData.Subject = reqData.Subject
    }
    byte, err := proto.Marshal(rspData)
    cool.Access(p.Ctx)
    p.Ctx.ResponseWriter.Write(byte)
}
