package main

import (
	"fmt"
	"mygofarm/Infrastructures"
)

//GO Web开发较通用的脚手架模板
func main() {
	//初始化基础配置
	if err := Infrastructures.Init(); err != nil {
		fmt.Printf("init infrastructure failed, err:%v\n", err)
	}
	//初始化路由服务和项目服务启动

}
