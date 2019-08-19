package main

import (
	"fmt"
	"gitchat/src/github.com/AliIoTApi/7/util"
)

func main()  {
	//创建产品
	createProductUrl := util.IoT_CreateProduct("product_name","node_type","data_format")
	fmt.Println( "createProductUrl:"+createProductUrl)

	//注册设备
	registerDeviceUrl := util.IoT_RegisterDevice("product_key","device_name")
	fmt.Println( "registerDeviceUrl:"+registerDeviceUrl )
}