package service

import (
	"fmt"
	"ginchat/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context){
	data := make([]*models.UserBasic,10)
	data = models.GetUserList()

	c.JSON(200,gin.H{
		"message":data,
	})
}

func CreateUser(c *gin.Context){
	user :=models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword :=c.Query("repassword")

	data := models.FindUserByName(user.Name)

	if data !=nil{
		
	}

	if password != repassword{
		c.JSON(
			-1,gin.H{
				"message":"two password is diffenert!",
			},
		)
		return
	}
	user.PassWord = password
	fmt.Println("user:",user)
	models.CreateUser(user)
	c.JSON(200,gin.H{
		"message":"create user successful!",
	})
}

func DeleteUser(c *gin.Context){
	user := models.UserBasic{}
	id,_ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200,gin.H{
		"message":"delete user successful!",
	})
}

func UpdateUser(c *gin.Context){
	user := models.UserBasic{}
	id,_ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name=c.PostForm("name")
	user.PassWord=c.PostForm("password")

	_,err:=govalidator.ValidateStruct(user)
	if err !=nil {
		fmt.Println(err)
		c.JSON(200,gin.H{
			"message":"修改参数不匹配!",
		})
	}

	models.UpdateUser(user)
	c.JSON(200,gin.H{
		"message":"update user successful!",
	})
}