package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"time"
)

func main()  {

	accessKeyID := "xxxx"

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

	fmt.Println( v )
}
