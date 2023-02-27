package main

import (
	"ginchat/models"
	"ginchat/router"

	"ginchat/utils"
	"log"
	"os"
	"time"

	//"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB
func main(){
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	DB,_ = gorm.Open(mysql.Open("root:123456@tcp(10.124.148.61:3306)/ginchat"),&gorm.Config{Logger: newLogger})
	DB.AutoMigrate(models.UserBasic{})	//没有表的话会自动创建
	// user:= &models.UserBasic{}
	// user.Name = "test_user"
	utils.InitConfig()
	utils.InitMySQL()
	r :=router.Router()
	r.Run(":8081")
}