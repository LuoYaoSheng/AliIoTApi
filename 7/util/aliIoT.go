package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"strconv"
	"strings"
	"time"
)


/**
*  设备管理相关 API
 */
//创建产品
func IoT_CreateProduct(ProductName,NodeType,DataFormat string) string {
	parm := make(map[string] string)
	parm["Action"] =  "CreateProduct"
	parm["ProductName"] =  ProductName
	parm["NodeType"] =  NodeType
	parm["DataFormat"] = DataFormat

	objUrl := signUrl(parm)
	return objUrl
}
//注册设备
func IoT_RegisterDevice(ProductKey,DeviceName string) string  {
	parm := make(map[string] string)
	parm["Action"] =  "RegisterDevice"
	parm["DeviceName"] =  DeviceName
	parm["ProductKey"] =  ProductKey

	objUrl := signUrl(parm)
	return objUrl
}
//该接口用于查看一个产品下多个设备的运行状态，单次最多可查询50个设备
func IoT_BatchGetDeviceState(ProductKey string, DeviceName []string) string  {
	parm := make(map[string] string)
	parm["Action"] =  "BatchGetDeviceState"
	parm["ProductKey"] =  ProductKey

	for k,v := range DeviceName {
		key := "DeviceName."+strconv.Itoa((k + 1))
		parm[key] = v
	}
	objUrl := signUrl(parm)
	return objUrl
}

//调用该接口在指定产品下批量自定义设备名称。IoT平台将检查名称的合法性。
func IoT_BatchCheckDeviceNames(ProductKey string,DeviceName []string)string  {
	parm := make(map[string] string)
	parm["Action"] =  "BatchCheckDeviceNames"
	parm["ProductKey"] =  ProductKey

	for k,v := range DeviceName {
		key := "DeviceName."+strconv.Itoa((k + 1))
		parm[key] = v
	}
	objUrl := signUrl(parm)
	return objUrl
}

//调用该接口根据申请批次ID（ApplyId）批量注册设备.
func IoT_BatchRegisterDeviceWithApplyId(ProductKey,ApplyId string)string  {
	parm := make(map[string] string)
	parm["Action"] =  "BatchRegisterDeviceWithApplyId"
	parm["ApplyId"] =  ApplyId
	parm["ProductKey"] =  ProductKey

	objUrl := signUrl(parm)
	return objUrl
}

//调用该接口查询批量注册的设备信息
func IoT_QueryPageByApplyId(ApplyId,PageSize,CurrentPage string)string  {
	parm := make(map[string] string)
	parm["Action"] =  "QueryPageByApplyId"
	parm["ApplyId"] =  ApplyId
	parm["PageSize"] =  PageSize
	parm["CurrentPage"] =  CurrentPage

	objUrl := signUrl(parm)
	return objUrl
}


func signUrl(requests map[string] string) string  {

	iotUrl := "http://iot.cn-shanghai.aliyuncs.com/?"
	accessKeyID := "LTAIaLC3JoT3Djqv"
	accessSecret := "hsRrxobaFmbhgyuafEI3f39EBbAYx3"

	v := url.Values{}
	//公共参数
	v.Add("Format", "JSON")
	v.Add("Version", "2018-01-20")
	v.Add("AccessKeyId", accessKeyID)
	v.Add("SignatureMethod", "HMAC-SHA1")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "1.0")
	v.Add("SignatureNonce", uuid.NewV4().String())
	v.Add("RegionId", "cn-shanghai")

	//请求参数
	for  key,value := range requests{
		v.Add(key,value)
	}
	//fmt.Println(v)

	urlBody := v.Encode()
	urlBody = percent_encode(urlBody)

	signature :=  sign(accessSecret,urlBody)
	urlBody = v.Encode()
	objUrl := iotUrl + urlBody + "&Signature=" + signature
	return objUrl
}

func sign(accessSecret , parmStr string)string {

	stringToSign := "GET&%2F&" + percent_encode(parmStr)

	//fmt.Println("待签名："+stringToSign)

	key :=  []byte(accessSecret+"&")
	h := hmac.New(sha1.New, key )
	h.Write([]byte(stringToSign))
	signature:= string(h.Sum(nil))
	src := []byte(signature)
	signature = base64.StdEncoding.EncodeToString(src)
	return signature
}

func percent_encode(encodeStr string) string  {
	str := encodeStr

	str = strings.ReplaceAll(str,"%3A", "%253A")    // 时间处的 :   处理

	str = strings.ReplaceAll(str,"%5B", "%255B")
	str = strings.ReplaceAll(str,"%5D", "%255D")



	str = strings.ReplaceAll(str,"&", "%26")
	str = strings.ReplaceAll(str,"=", "%3D")

	str = strings.ReplaceAll(str,"+", "%20")
	str = strings.ReplaceAll(str,"*", "%2A")
	str = strings.ReplaceAll(str,"%7E", "~")

	return str
}