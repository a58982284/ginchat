package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



func Router() *gin.Engine{
	r:=gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index",service.GetIndex)
	r.GET("/user/getuserList",service.GetUserList)
	r.GET("/user/createUser",service.CreateUser)
	r.GET("/user/deleteUser",service.DeleteUser)
	r.POST("/user/updateUser",service.UpdateUser)
	return r
}