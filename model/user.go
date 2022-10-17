package model

import (
	"gin-mysql-vue/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int
	Name  string `gorm:"type:varchar(20);not null" json:"name" validate:"required,min=4,max=12" label:"username"`
	Email string `gorm:"type:varchar(50);not null" json:"email" validate:"required" label:"email"`
	Role  int    `gorm:"type:int;DEFAULT:1" json:"role" label:"role"`
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	db := GetDB()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	db := GetDB()
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	db := GetDB()
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// GetUsers 查询用户列表
func GetUsers(name string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	db := GetDB()

	if name != "" {
		db.Select("id,name,role,created_at").Where(
			"name LIKE ?", name+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"name LIKE ?", name+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,name,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}
