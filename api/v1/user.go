package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	//todo 添加用户
	var users model.User
	_ = c.ShouldBindJSON(&users)
	code = model.CheckUser(users.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&users)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   users,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询单个用户
func GetUser(c *gin.Context) {
	//todo 查询单个用户
	id := c.Param("id")
	user, err := model.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR_USER_NOT_EXIST,
			"msg":    errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_EXIST),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.SUCCESS,
		"data":   user,
		"msg":    errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   users,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var users model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&users)
	code = model.CheckUser(users.Username)
	if code == errmsg.SUCCESS {
		code = model.EditUsers(id, &users)
	}
	if code == errmsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUsers(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}
