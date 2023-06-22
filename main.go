package main

import (
	"GoHomework/api"
	"GoHomework/boot"
)

func main() {
	boot.InitMysql()
	api.InitRouter()
}
