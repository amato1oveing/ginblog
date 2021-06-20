package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCate(c *gin.Context) {
	//todo 添加分类
	var cates model.Category
	_ = c.ShouldBindJSON(&cates)
	code = model.CheckCate(cates.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&cates)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cates,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询单个分类
//func GetOneUser(c *gin.Context) {
//
//}

//查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	cates := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cates,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCate(c *gin.Context) {
	var cates model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&cates)
	code = model.CheckCate(cates.Name)
	if code == errmsg.SUCCESS {
		code = model.EditCate(id, &cates)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}
