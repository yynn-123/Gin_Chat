package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// 这是程序入口
// 这是IM及时通讯
// 这是hot-fix分支提交
// 这是master分支提交
// 这是hot-fix分支第二次提交
// push test
// 远程库 pull test
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8081")

}
