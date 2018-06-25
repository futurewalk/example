var Resolver = {}

//for bytes data resolver
//exmple proto message
// message Demo {
//    required bytes data = 1;
// }
// in this method ,you can replace this data to another protoBuf message for field type bytes
Resolver.byteInvoke = function (plugin) {

    if (plugin.StructureName == "CoolRequest" && plugin.FieldName == "data") {
        plugin.Swap("DataName");
    }
    if (plugin.StructureName == "CoolRequest" && plugin.FieldName == "Optype") {
        //plugin.Remove()
    }
    if (plugin.StructureName == "Request" && plugin.FieldName == "data" && plugin.Url == "/v1/device/login") {
        plugin.Swap("LoginReqData")
    }
    if (plugin.StructureName == "Request" && plugin.FieldName == "data" && plugin.Url == "/v1/device/device-bind-status") {
        plugin.Swap("QueryDeviceBindStatusReqData")
    }
    if (plugin.StructureName == "Request" && plugin.FieldName == "data" && plugin.Url == "/v1/device/get-deviceInfo") {
        plugin.Swap("GetDeviceInfoReqData","userBind.proto");
    }
    if (plugin.StructureName == "BytesMessageReq" && plugin.FieldName == "Data" && plugin.Url == '/v1/example/BytesSupport') {
        plugin.Swap("BytesMessageInfo");
    }

}
// in this callBack method ,you can set a field with a custom data type on another protobuf file。
Resolver.dataInvoke = function (plugin){

}
Resolver.RspInvoke = function (plugin) {
    if (plugin.StructureName == "BytesMessageRsp" && plugin.FieldName == "Data" && plugin.Url == '/v1/example/BytesSupport') {
        plugin.Swap("BytesMessageInfo");
    }
}
window.Resolver = Resolver;