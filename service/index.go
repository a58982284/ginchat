package service

import "github.com/gin-gonic/gin"
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"welcome!",
	})
}