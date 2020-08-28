package m_user

import (
	"github.com/jinzhu/gorm"
)

type UserEntity struct {
	Name     string `json:"name" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Age      string `json:"age" gorm:"type:tinyint(4)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
	gorm.Model
}
