package service

//
//import (
//	"fmt"
//	"ginchat/models"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func AuthHandler(c *gin.Context) {
//	// 用户发送用户名和密码过来
//	user := models.UserBasic{}
//	name := c.PostForm(user.Name)
//	pass := c.PostForm(user.PassWord)
//	fmt.Println("name", name)
//	fmt.Println("pass", pass)
//	//if err != nil {
//	//	c.JSON(http.StatusOK, gin.H{
//	//		"code": 2001,
//	//		"msg":  "无效的参数",
//	//	})
//	//	return
//	//}
//	dataBaseUser := models.FindUserByNameAndPwd(name, pass)
//	fmt.Println(">>>>>>>>>>>>>>>>>>>", dataBaseUser.Name, dataBaseUser.PassWord)
//	// 校验用户名和密码是否正确
//	if dataBaseUser.Name != "" {
//		// 生成Token
//		tokenString, _ := models.GenToken(user.Name, user.PassWord)
//		c.JSON(http.StatusOK, gin.H{
//			"code": 2000,
//			"msg":  "success",
//			"data": gin.H{"token": tokenString},
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 2002,
//		"msg":  "鉴权失败",
//	})
//	return
//}
