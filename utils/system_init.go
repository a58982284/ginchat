package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig(){
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err:=viper.ReadInConfig()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("config app:",viper.Get("app"))
	fmt.Println("config mysql:",viper.Get("mysql"))
}

func InitMySQL(){
	// newLogger := logger.New(
	// 	log.New(os.Stdout,"\r\n",log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,
	// 		LogLevel: logger.Info,
	// 		Colorful: true,
	// 	},
	// )
	DB,_ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),&gorm.Config{})
	// DB,_ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),&gorm.Config{Logger: newLogger})
	// user:=models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)

	// if err != nil{
	// 	panic("failed to connect database")
	// 	return
	//
}
