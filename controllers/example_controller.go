package controllers

import (
	"github.com/astaxie/beego"
	"example/protobuf"
	"github.com/gogo/protobuf/proto"
	"doc-cool/cool"
    "log"
)

type ExampleController struct {
	beego.Controller
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
		Data:&protobuf.InfoMessage{},
	}

	err := proto.Unmarshal(p.Ctx.Input.RequestBody,reqData)
	if err == nil{
		data :=reqData.Data
		rspData.Data.Method = data.Method
		rspData.Data.Url = data.Url
	}
	rspData.DocCoolName = reqData.DocCoolName
	byte,err := proto.Marshal(rspData)
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
	log.Println("获取请求数据:",reqData)
	if err == nil {
		data := reqData.Data
		rspData.Data = data
	}
	rspData.BytesName = reqData.DocCoolName
	byte, err := proto.Marshal(rspData)
	cool.Access(p.Ctx)
	p.Ctx.ResponseWriter.Write(byte)
}
