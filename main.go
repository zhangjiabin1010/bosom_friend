package main

import "bosom_friend/router"

func main(){
	r := router.Router()
	r.Run() //默认在本地8080端口启动服务
}