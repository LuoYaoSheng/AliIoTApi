package main

import (
	"fmt"
	"gitchat/src/github.com/AliIoTApi/6/util"
)

func main()  {
	objUrl := util.RegisterDevice("a1EkBx6nvxB","1533023037")
	fmt.Println( "objUrl:"+objUrl )
}