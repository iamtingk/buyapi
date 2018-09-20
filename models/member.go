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

// 新增商品
func (member *Member) Insert() (err error) {
	result := configDB.GormOpen.Create(&member)

	fmt.Println(member.Id)
	if result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}
