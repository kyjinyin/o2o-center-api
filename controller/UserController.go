package controller
import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "fmt"
    //"vq0599/models"
    //"vq0599/common"
    //"strings"
)

func GetUser(hasPermission bool) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"resut":"ok",
		})
	}
}