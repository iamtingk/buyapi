package models

import (
	msg "buyapi/config"
	configDB "buyapi/database"
	"errors"
	"time"
)

type Product struct {
	ID        int64     `json:"id"`        // 商品Id
	Name      string    `json:"name"`      // 商品名稱
	Img       string    `json:"img"`       // 圖片
	Price     string    `json:"price"`     // 價錢
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間

}

var Products []Product

// 新增商品
func (product Product) Insert() (err error) {
	result := configDB.GormOpen.Create(&product)
	if result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}

// 查詢全部
func (product *Product) GetProducts() (data interface{}, err error) {
	var products []Product
	result := configDB.GormOpen.Find(&products)
	if result.Error != nil {
		err = result.Error
		return "", err
	}
	return products, nil
}

// 修改商品
func (product *Product) Update(id int64) (updateProduct Product, err error) {

	if err = configDB.GormOpen.Select([]string{"id"}).First(&updateProduct, id).Error; err != nil {
		return
	}
	if err = configDB.GormOpen.Model(&updateProduct).Updates(&product).Error; err != nil {
		return
	}
	return
}

// 查詢圖片名稱 - 指定id
func (product *Product) GetProductImg(id int64) (imgName string, err error) {

	var productRow Product
	configDB.GormOpen.Table("products").Where("id=?", id).Scan(&productRow)

	result := productRow.Img
	if len(result) > 0 {
		return result, nil
	}

	return result, errors.New(msg.CONTINUE_NOT_FOUND_IMAGE)

}
