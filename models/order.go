package models

import (
	msg "buyapi/config"
	configDB "buyapi/database"
	"errors"
	"fmt"
	"time"
)

type RequestOrder struct {
	Token       string               `json:"token"`         // 會員憑證
	OrderDetail []RequestOrderDetail `json:"order_details"` // 訂單明細
}

type RequestOrderDetail struct {
	ProductId int64 `json:"product_id"` // 商品Id
	Num       int64 `json:"num"`        // 數量
}

type Order struct {
	Id        int64     `json:"id"`        // 訂單Id
	MemberId  int64     `json:"member_id"` // 會員Id
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間
}

type OrderDetail struct {
	Id        int64 `json:"id"`         // 訂單明細Id
	OrderId   int64 `json:"order_id"`   // 訂單Id
	ProductId int64 `json:"product_id"` // 商品Id
	Num       int64 `json:"num"`        // 數量
}

func (order *Order) InsertOrder(orderInfo Order, detailsInfo []OrderDetail) (data *Order, err error) {
	if err := configDB.GormOpen.Table("Orders").Create(&orderInfo).Error; err != nil {
		return nil, errors.New(msg.SIGNUP_ERROR)
	}

	for i, _ := range detailsInfo {
		detailsInfo[i].OrderId = orderInfo.Id
	}

	//新增訂單明細
	if err := InsertOrderdetail(detailsInfo){
		return nil, errors.New(msg.SQL_WRITE_ERROR)
	}
	
	return &orderInfo, nil
}

func InsertOrderdetail(detailsInfo []OrderDetail) (err error) {
	sql := "INSERT INTO `OrderDetails` (`order_id`,`product_id`,`num`) VALUES "
	count := len(detailsInfo[0:]) - 1
	for i, detail := range detailsInfo {
		if i == (count) {
			sql += fmt.Sprintf("('%d','%d','%d');", detail.OrderId, detail.ProductId, detail.Num)
		} else {
			sql += fmt.Sprintf("('%d','%d','%d'),", detail.OrderId, detail.ProductId, detail.Num)
		}
	}
	fmt.Println(sql)

	tx := configDB.GormOpen.Begin()

	if err := configDB.GormOpen.Exec(sql).Error; err != nil {
		tx.Rollback()
		return errors.New(msg.SQL_WRITE_ERROR)
	}

	tx.Commit()
	return nil
}
