package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "/views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	//c.JSON(200, gin.H{
	//	"message": "welcome !!",
	//})
}
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	//c.JSON(200, gin.H{
	//	"message": "welcome !!",
	//})
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	userID, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userID)
	user.Identity = token
	ind.Execute(c.Writer, user)
	//c.JSON(200, gin.H{
	//	"message": "welcome !!",
	//})
}

func Chat(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
