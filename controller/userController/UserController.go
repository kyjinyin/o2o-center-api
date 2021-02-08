package userController
import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
        "message":"userController",
    })
}

// func GetUser(hasPermission bool) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK,gin.H{
// 			"resut":"ok",
// 		})
// 	}
// }