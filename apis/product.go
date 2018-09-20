package apis

import (
	code "buyapi/config"
	config "buyapi/config"
	msg "buyapi/config"
	model "buyapi/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 展示全部商品
func ShowProducts(c *gin.Context) {
	var product model.Product

	// 執行-查詢全部商品
	result, err := product.GetProducts()

	if err != nil {
		showJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
		return
	}
	showJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)

}

// 增加商品
func CreateProduct(c *gin.Context) {
	var product model.Product

	// 取得參數
	product.Name = c.Request.FormValue("productName")
	product.Price = c.Request.FormValue("productPrice")
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	file, header, err := c.Request.FormFile("productImage")
	if err != nil {
		showJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
		return
	}

	//驗證參數是否有值
	if len(product.Name) > 0 && len(product.Price) > 0 {
		filename := header.Filename
		if file == nil && len(filename) <= 0 {
			//沒有圖片
			fmt.Println(msg.NOT_FOUND_IMAGE, err)
			showJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
			return
		}

		//圖檔重新命名
		product.Img = fileRename(filename)

		// 執行-增加商品
		err := product.Insert()
		if err != nil {
			//如果出錯，就刪除剛存的圖片
			os.Remove(config.IMAGE_PATH + product.Img)
			showJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}
		// 新增圖片
		AddImg(c, file, product.Img, config.IMAGE_PATH)

		showJsonDATA(c, code.SUCCESS, msg.CREATE_SUCCESS, "")

	} else {
		showJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

}

// 修改商品
func UpdateProduct(c *gin.Context) {
	var product model.Product
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	product.Name = c.Request.FormValue("productName")
	product.Price = c.Request.FormValue("productPrice")
	product.UpdatedAt = time.Now()

	file, header, err := c.Request.FormFile("productImage")
	if err != nil {
		showJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
		return
	}

	//驗證參數是否有值
	if len(product.Name) > 0 && len(product.Price) > 0 {
		filename := header.Filename
		if file == nil && len(filename) <= 0 {
			//找不到圖片
			showJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
			return
		}

		// 執行-查詢原本圖檔的名稱
		oldImgName, err := product.GetProductImg(id)
		if err != nil {
			//找不到圖片刪除，仍繼續
			fmt.Println(err)
		}

		//圖檔重新命名
		product.Img = fileRename(filename)

		// 執行-修改商品
		err = product.Update(id)
		if err != nil {
			//如果出錯，就刪除剛存的圖片
			os.Remove(config.IMAGE_PATH + product.Img)
			showJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}

		//刪除原本照片
		err = os.Remove(config.IMAGE_PATH + oldImgName)
		if err != nil {
			//找不到圖片刪除，仍繼續
			fmt.Println(msg.CONTINUE_NOT_FOUND_IMAGE)
		}

		// 新增圖片
		AddImg(c, file, product.Img, config.IMAGE_PATH)

		showJsonDATA(c, code.SUCCESS, msg.UPDATE_SUCCESS, "")

	} else {
		showJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

}

//刪除商品
func DestroyProduct(c *gin.Context) {
	var product model.Product
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// 執行-查詢原本圖檔的名稱
	oldImgName, err := product.GetProductImg(id)
	if err != nil {
		//找不到圖片刪除，仍繼續
		fmt.Println(err)
	}

	// 執行-刪除商品
	err = product.Destroy(id)
	if err != nil {
		//刪除失敗
		showJsonMSG(c, code.ERROR, msg.DELETE_ERROR)
		return
	}

	// 刪除原本照片
	err = os.Remove(config.IMAGE_PATH + oldImgName)
	if err != nil {
		// 找不到圖片刪除，仍繼續
		fmt.Println(msg.CONTINUE_NOT_FOUND_IMAGE)
	}

	showJsonDATA(c, code.SUCCESS, msg.DELETE_SUCCESS, "")

}

func fileRename(filename string) string {
	// 替換圖片檔名
	newFileName := GetMD5Hash(filename + time.Now().String())
	dotIndex := strings.LastIndex(filename, ".") //取得最後的.的索引值
	if dotIndex != -1 && dotIndex != 0 {
		newFileName += filename[dotIndex:] //取得副檔名
	}
	//輸出 檔名＋副檔名

	return newFileName
}

func showJsonMSG(c *gin.Context, code int64, msg string) {
	msg = strings.Replace(msg, "\b", "", -1)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}

func showJsonDATA(c *gin.Context, code int64, msg string, data interface{}) {
	msg = strings.Replace(msg, "\b", "", -1)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
}

func AddImg(c *gin.Context, file io.Reader, fileName string, filePath string) {
	//抓取新圖片到指定目錄
	out, err := os.Create(filePath + fileName)
	if err != nil {
		//沒有image資料夾
		showJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE_FOLDER)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		//寫入檔案失敗
		showJsonMSG(c, code.ERROR, msg.WRITE_FILE_ERROR)
		return
	}
}
