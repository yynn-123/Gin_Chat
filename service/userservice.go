package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @param phone query string false "电话号码"
// @param email query string false "电子邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
			"code":    -1,
			"data":    user,
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	dataName := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(-1, gin.H{
			"message": "用户名或密码不能为空",
			"code":    -1,
			"data":    user,
		})
		return
	}
	if dataName.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册",
			"code":    -1,
			"data":    user,
		})
		return
	}
	dataPhone := models.FindUserByPhone(user.Phone)
	if dataPhone.Phone != "" {
		c.JSON(-1, gin.H{
			"message": "手机号已注册",
		})
		return
	}
	dataEmail := models.FindUserByEmail(user.Email)
	if dataEmail.Email != "" {
		c.JSON(-1, gin.H{
			"message": "电子邮箱已注册",
		})
		return
	}
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "创建用户成功",
		"data":    user,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
		"code":    0,
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "修改用户成功",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary 用户登陆
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}
	if user.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已存在",
			"data":    data,
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	token, _ := models.GenToken(name, pwd)
	user.Identity = token
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登陆成功",
		"data":    data,
	})
}

// 防止跨域站点的伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(ws)
	MsgHandler(ws, c)
}
func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}

}
func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(userid)
	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "查询好友列表成功",
		"data":    users,
	})
	utils.RespOKList(c.Writer, users, len(users))
}
