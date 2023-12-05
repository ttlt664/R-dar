package routes

import (
	"github.com/gin-gonic/gin"
	"radar_src/config"
	"radar_src/middleware"
)

func InitRouter() {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	//定义接口版本路由
	authv1 := r.Group("api/v1")
	//给authv1增加中间件认证
	authv1.Use(middleware.JwtToken())
	{
		//User模块的路由
		authv1.PUT("user/:id", v1.EditUser)
		authv1.DELETE("user/:id", v1.DelUser)

		//Category模块的路由
		authv1.POST("category/add", v1.AddCategory)
		authv1.PUT("category/:id", v1.EditCategory)
		authv1.DELETE("category/:id", v1.DelCategory)

		//Article模块的路由
		authv1.POST("article/add", v1.AddArticle)
		authv1.PUT("article/:id", v1.EditArticle)
		authv1.DELETE("article/:id", v1.DelArticle)
	}

	router := r.Group("api/v1")
	{
		//User模块的路由
		router.GET("user/exist", v1.UserExist)
		router.GET("user/get_user_list", v1.GetUserList)
		router.GET("user/:id", v1.GetUser)
		router.POST("user/add", v1.AddUser)
		//Category模块的路由
		router.GET("category/exist", v1.CategoryExist)
		router.GET("category/get_category_list", v1.GetCategoryList)
		router.GET("category/:id", v1.GetCategoryInfo)
		//Article模块的路由
		router.GET("article/get_article_list", v1.GetArticleList)
		router.GET("article/get_article_list_by_category", v1.GetArticleListByCatrgory)
		router.GET("article/get_article/:id", v1.GetArticle)
		router.POST("login", v1.Login)
		router.POST("upload", v1.Upload)
	}

	r.Run(config.HttpPort)
}
