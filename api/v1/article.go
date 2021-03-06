package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

///添加文章
func AddArt(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)
	code = model.CreateArt(&art)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	artList, code := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   artList,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	artList, code := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   artList,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&art)
	code = model.EditArt(id, &art)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//批量删除文章
func DeleteArts(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	code = model.DeleteArts(ids) //批量删除
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
