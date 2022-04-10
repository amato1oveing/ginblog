package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类的路由接口
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章的路由接口
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		auth.DELETE("article/ids", v1.DeleteArts)

		auth.POST("quit", v1.Logout)
	}

	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("user_only", v1.GetUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCate)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)
	}
	_ = r.Run(utils.HttpPort)
}
