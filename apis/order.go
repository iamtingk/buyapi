package apis

import (
	code "buyapi/config"
	msg "buyapi/config"
	model "buyapi/models"
	. "buyapi/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 增加訂單
func TestOrder(c *gin.Context) {

	type Pig struct {
		Head string `json:"head"`
		Num  int64  `json:"num"`
	}

	type KDRespBody struct {
		Text []Pig `json:"pig"`
		Num  int64 `json:"num"`
	}

	var abc KDRespBody
	err := c.BindJSON(&abc)
	if err != nil {
		fmt.Println("errKKKK")
	}
	fmt.Println(abc.Text)
	fmt.Println(len(abc.Text))
	fmt.Println(abc.Text[1])
	fmt.Println(abc.Text[2].Head)

}

// 增加訂單
func CreateOrder(c *gin.Context) {

	var request model.RequestOrder

	// 取得回傳的JSON
	err := c.BindJSON(&request)
	if err != nil {
		//沒有資料
		// ShowJsonMSG(c, code.ERROR, msg.REQUEST_DATA_ERROR)
		// return
		fmt.Println(msg.REQUEST_DATA_ERROR)
	}

	var token string
	var order model.Order
	var orderDetails []model.OrderDetail

	token = request.Token
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	memberId, err := model.CheckToken(token)
	if err != nil {
		fmt.Println(msg.NOT_SIGNIN)
		ShowJsonMSG(c, code.ERROR, msg.NOT_SIGNIN)
		return
	}
	order.MemberId = memberId

	tmp := make([]model.OrderDetail, len(request.OrderDetail))

	for i, detail := range request.OrderDetail {
		tmp[i].Num = detail.Num
		tmp[i].ProductId = detail.ProductId
		fmt.Println(detail.ProductId)
	}
	orderDetails = tmp

	// 執行-新增訂單
	result, err := order.CreateOrder(order, orderDetails)
	fmt.Println(result) // id
	if err != nil {
		// 註冊失敗
		ShowJsonMSG(c, code.ERROR, msg.SIGNUP_ERROR)
		return
	}

	ShowJsonDATA(c, code.SUCCESS, msg.CREATE_SUCCESS, "")

}
