package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// 这是程序入口
// 这是IM及时通讯
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8081")

}
