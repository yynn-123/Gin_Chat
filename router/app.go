package router

import (
	"ginchat/docs"
	"ginchat/models"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/assert", "assert")
	//r.LoadHTMLGlob("views/**/*")
	docs.SwaggerInfo.BasePath = ""
	userGroup := r.Group("/user")
	{
		userGroup.POST("/getUserList", models.JWTAuthMiddleware(), service.GetUserList)
		userGroup.POST("/createUser", service.CreateUser)
		userGroup.POST("/deleteUser", models.JWTAuthMiddleware(), service.DeleteUser)
		userGroup.POST("/updateUser", models.JWTAuthMiddleware(), service.UpdateUser)
		userGroup.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)
	// 获取token
	//r.GET("/auth", service.AuthHandler)
	// 发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	r.POST("/searchFriends", service.SearchFriends)
	return r
}
