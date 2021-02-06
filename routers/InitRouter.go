  package routers
  import (
	"github.com/gin-gonic/gin"
	"o2o-center-api/controller"
	"fmt"
	//"vq0599/middleware"
  )
  func InitRouter() *gin.Engine {
	router := gin.Default()
	_ , err := fmt.Println(controller.GetUser(true))
	if err != nil{
		fmt.Printf("InitRouter failed, err:%v\n", err)
	}
	// api := router.Group("/api/v1")
	// apiAdmin := api.Group("/admin")
  
	// api.GET("/articles", controller.GetArticles(false))
	// api.GET("/articles/:id", controller.GetArticle(false))
  
	// api.GET("/pv/:id", controller.UpdatePvs)
	// api.POST("/like/:id", controller.UpdateLikes)
  
	// ----------- 管理后台接口 -----------
	// apiAdmin.POST("/login", controller.Login)
	// apiAdmin.GET("/login", controller.GetLoginStatus)
  
	// apiAdmin.Use(middleware.Jwt())
	// {
	//   apiAdmin.GET("/articles", controller.GetArticles(true))
	//   apiAdmin.GET("/articles/:id", controller.GetArticle(true))
	//   apiAdmin.POST("/articles", controller.AddArticle)
	//   apiAdmin.DELETE("/articles/:id", controller.DeleteArticle)
	//   apiAdmin.PATCH("/articles/:id", controller.UpdateArticle)
	//   apiAdmin.POST("/upload/images", controller.UploadImage)
	// }
  
	return router
  }