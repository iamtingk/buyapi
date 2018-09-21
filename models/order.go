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

// 查詢全部
func (order *Order) QueryOrders(memberId int64) (data interface{}, err error) {
	var orders []Order
	result := configDB.GormOpen.Table("Orders").Where("member_id=?", memberId).Find(&orders)
	if result.Error != nil {
		err = result.Error
		return nil, err
	} else if len(orders[0:]) == 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}
	return orders, nil
}

// 查詢訂單明細
func QueryOrderDetail(orderId int64) (data interface{}, err error) {
	var orderDetail []OrderDetail
	result := configDB.GormOpen.Table("OrderDetails").Where("order_id=?", orderId).Find(&orderDetail)
	if result.Error != nil {
		err = result.Error
		return nil, err
	} else if len(orderDetail[0:]) == 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}

	return orderDetail, nil
}

// 新增訂單
func (order *Order) InsertOrder(orderInfo Order, detailsInfo []OrderDetail) (data *Order, err error) {
	if err := configDB.GormOpen.Table("Orders").Create(&orderInfo).Error; err != nil {
		return nil, errors.New(msg.SQL_WRITE_ERROR)
	}

	// 設置訂單明細OrderId
	for i, _ := range detailsInfo {
		detailsInfo[i].OrderId = orderInfo.Id
	}

	//新增訂單明細
	if err := InsertOrderdetail(detailsInfo); err != nil {
		fmt.Println("332211")
		return nil, errors.New(msg.SQL_WRITE_ERROR)
	}

	return &orderInfo, nil
}

//新增訂單明細
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

// 刪除訂單
func (order *Order) Destroy(order_id int64) (err error) {

	var tmpOrder Order
	var tmpOrderDetail OrderDetail
	if err = configDB.GormOpen.Table("Orders").Select([]string{"id"}).First(&tmpOrder, order_id).Error; err != nil {
		fmt.Println("111")
		return err
	}

	if err = configDB.GormOpen.Table("OrderDetails").Select([]string{"order_id"}).First(&tmpOrderDetail, order_id).Error; err != nil {
		fmt.Println("222")
		return err
	}

	fmt.Println(order_id)
	if err = configDB.GormOpen.Table("Orders").Where("id=?", order_id).Delete(&tmpOrder).Error; err != nil {
		return err
	}

	if err = configDB.GormOpen.Table("OrderDetails").Where("order_id=?", order_id).Delete(&tmpOrderDetail).Error; err != nil {
		return err
	}

	return nil
}
