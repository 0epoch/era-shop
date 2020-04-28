package main

import (
	"era-shop.xyz/era-shop/model"
	"era-shop.xyz/era-shop/pkg/config"
	"era-shop.xyz/era-shop/router"
)

func main() {
	//fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	config.Load()
	model.Load()
	r := router.Load()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

