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

func (order *Order) CreateOrder(orderInfo Order, detailsInfo []OrderDetail) (data *Order, err error) {

	if err := configDB.GormOpen.Table("Orders").Create(&orderInfo).Error; err != nil {
		return nil, errors.New(msg.SIGNUP_ERROR)
	}

	detailsInfo[0].OrderId = orderInfo.Id
	// for i, _ := range detailsInfo {
	// 	detailsInfo[i].OrderId = orderInfo.Id
	// }
	fmt.Println(detailsInfo[0].OrderId)
	// // var orderDetail OrderDetail

	// tmp := make([]OrderDetail, len(detailsInfo))
	// for i, detail := range detailsInfo {
	// 	tmp[i].Id = orderInfo.Id
	// 	tmp[i].Num = detail.Num
	// 	tmp[i].ProductId = detail.ProductId
	// }

	return &orderInfo, nil
}

func (orderDetail *OrderDetail) CreateOrderdetail(orderDetailInfo *OrderDetail) (data *OrderDetail, err error) {

	if err := configDB.GormOpen.Table("OrderDetails").Create(&orderDetailInfo).Error; err != nil {
		return nil, errors.New(msg.SIGNUP_ERROR)
	}
	fmt.Println(orderDetailInfo.OrderId)

	return orderDetailInfo, nil
}
