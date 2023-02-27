package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)
type UserBasic struct {
	gorm.Model
	Name string
	PassWord string
	Phone string	`valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email string	`valid:"email"`
	Identity string
	ClientIp string
	ClientPort string
	LoginTime uint64
	HeartbeatTime uint64
	LoginOutTime uint64	`gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout bool
	DeviceInfo string
}
//func (variable_name variable_data_type) function_name() [return_type]
func (table *UserBasic) TableName() string{
	return "user_basic"
}
// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getuserList [get]
func GetUserList()	[]*UserBasic{
	data := make([] *UserBasic,10)
	utils.DB.Find(&data)
	return data
}
// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(user UserBasic) *gorm.DB{
	fmt.Println("create_user:",&user)
	return utils.DB.Create(&user)
}
// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(user UserBasic) *gorm.DB{
	return utils.DB.Delete(&user)
}
// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(user UserBasic) *gorm.DB{
	return utils.DB.Model(&user).Updates(UserBasic{Name:user.Name,PassWord: user.PassWord})
}


func FindUserByName(name string) *gorm.DB{
	user := UserBasic{}
	return utils.DB.Where("name = ",name).First(&user)
}

func FindUserByPhone(phone string) *gorm.DB{
	user := UserBasic{}
	return utils.DB.Where("Phone = ",phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB{
	user := UserBasic{}
	return utils.DB.Where("Email = ",email).First(&user)
}