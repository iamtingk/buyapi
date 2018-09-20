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

// 增加商品
func SignUpMember(c *gin.Context) {
	var member model.Member

	// 取得參數
	member.Email = c.Request.FormValue("email")
	member.Phone = c.Request.FormValue("phone")
	member.Password = c.Request.FormValue("password")
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	fmt.Println(member.Email)
	fmt.Println(member.Phone)
	fmt.Println(member.Password)

	// 參數是否有值
	if len(member.Email) > 0 && len(member.Phone) > 0 && len(member.Password) > 0 {
		fmt.Println("OK")
		// 驗證規則
		if IsEmail(member.Email) && IsPhone(member.Phone) {

			// 執行-增加會員
			err := member.Insert()
			if err != nil {
				// 新增失敗
				ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
				return
			}
			ShowJsonDATA(c, code.SUCCESS, msg.CREATE_SUCCESS, "")
		} else {
			// 驗證失敗
			ShowJsonMSG(c, code.ERROR, msg.VERIFY_ERROR)
			return
		}

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

}
