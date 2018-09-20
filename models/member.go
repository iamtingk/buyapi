package models

import (
	configDB "buyapi/database"
	"fmt"
	"time"
)

type Member struct {
	Id            int64     `json:"id"`              // 會員Id
	Email         string    `json:"email"`           // 信箱
	Phone         string    `json:"phone"`           // 手機
	Password      string    `json:"password"`        // password
	Token         string    `json:"token"`           // token憑證
	IsEmailVerify int64     `json:"is_email_verify"` // 信箱是否驗證
	IsPhoneVerify int64     `json:"is_phone_verify"` // 手機是否驗證
	CreatedAt     time.Time `json:"createdAt"`       // 開始時間
	UpdatedAt     time.Time `json:"updatedAt"`       // 更新時間
}

type ShowToken struct {
	Token string `json:"token"` // token憑證
}

// 註冊會員
func (member *Member) Insert() (data *Member, err error) {
	if err := configDB.GormOpen.Create(&member).Error; err != nil {
		return member, err
	}
	fmt.Println(member.Email)
	return member, nil
}

// 登入會員
func (member *Member) Query(email string, password string) (data *Member, err error) {
	if err := configDB.GormOpen.Table("members").Where("email=? AND password=?", email, password).Scan(&member).Error; err != nil {
		return member, err
	}
	return member, nil
}
