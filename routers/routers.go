package routers

import (
	baseController "o2o-center-api/controller/baseController"
	userController "o2o-center-api/controller/userController"
	//"bubble/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// if setting.Conf.Release {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	r := gin.Default()
	// where is static reources
	//r.Static("/static", "static")
	// where is template -> for web not for api
	//r.LoadHTMLGlob("templates/*")

	r.GET("/", baseController.IndexHandler)

	r.POST("/", userController.IndexHandler)

	// v1
	// v1Group := r.Group("v1")
	// {
	// 	// 待办事项
	// 	// 添加
	// 	v1Group.POST("/todo", controller.CreateTodo)
	// 	// 查看所有的待办事项
	// 	v1Group.GET("/todo", controller.GetTodoList)
	// 	// 修改某一个待办事项
	// 	v1Group.PUT("/todo/:id", controller.UpdateATodo)
	// 	// 删除某一个待办事项
	// 	v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	// }
	return r
}