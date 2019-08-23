## 内容说明
使用GO语言来对接物联网平台API，通过具体调用阿里物联API来学习常规的加密算法，GO语言的加密算法使用。

此处仅限于学习交流，真实项目开发建议使用官方方法实现。



## 阿里云在线示例

[API Explorer][open-api] 提供在线调用阿里云产品，并动态生成 SDK 代码和快速检索接口等能力，能显著降低使用云 API 的难度。

阿里云SDK发布地址: https://develop.aliyun.com/tools/sdk#/go 


### 例子：注册设备
```go
package main

import (
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

func main() {
	client, err := iot.NewClientWithAccessKey("cn-shanghai", "<accessKeyId>", "<accessSecret>")

	request := iot.CreateGetGatewayBySubDeviceRequest()
	request.Scheme = "https"

	response, err := client.GetGatewayBySubDevice(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

```

[open-api]: https://api.aliyun.com/
