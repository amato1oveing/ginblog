package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password, data.Role)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.SetCookie("token", token, 36000, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, false)
	c.Redirect(http.StatusFound, "/login")
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"msg":    errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}
