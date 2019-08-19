package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"strings"
	"time"
)


/**
*  设备管理相关 API
 */
//注册设备
func RegisterDevice(ProductKey,DeviceName string) string  {
	parm := make(map[string] string)
	parm["Action"] =  "RegisterDevice"
	parm["DeviceName"] =  DeviceName
	parm["ProductKey"] =  ProductKey

	objUrl := signUrl(parm)
	return objUrl
}
/**
 * 生成签名后的url
 * @requests 具体请求参数/ 公共参数统一在内部写死 / 公共参数可通过配置传入
 * @return 签名后的请求url
 */
func signUrl(requests map[string] string) string  {

	iotUrl := "http://iot.cn-shanghai.aliyuncs.com/?"
	accessKeyID := "xxx"
	accessSecret := "xxx"

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
//Signature = Base64( HMAC-SHA1( AccessSecret, UTF-8-Encoding-Of(StringToSign) ) )
func sign(accessSecret , parmStr string)string {

	stringToSign := "GET&%2F&" + percent_encode(parmStr)
	key :=  []byte(accessSecret+"&")
	h := hmac.New(sha1.New, key )
	h.Write([]byte(stringToSign))
	signature:= string(h.Sum(nil))
	src := []byte(signature)
	signature = base64.StdEncoding.EncodeToString(src)
	return signature
}
//特殊字符进行操作
func percent_encode(encodeStr string) string  {
	str := encodeStr

	str = strings.ReplaceAll(str,"%3A", "%253A")    // 时间处的 :   处理

	str = strings.ReplaceAll(str,"&", "%26")
	str = strings.ReplaceAll(str,"=", "%3D")

	str = strings.ReplaceAll(str,"+", "%20")
	str = strings.ReplaceAll(str,"*", "%2A")
	str = strings.ReplaceAll(str,"%7E", "~")

	return str
}