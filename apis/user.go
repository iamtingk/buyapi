package apis

import (
	config "buyapi/config"
	model "buyapi/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//列表数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	if len(user.Username) > 0 && len(user.Password) > 0 {

		id, err := user.Insert()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "添加失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "添加成功",
			"data":    id,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失敗",
		})
	}
}

//添加数据
func Storee(c *gin.Context) {
	// namee := c.Request.FormValue("namee")
	// fmt.Println(namee)
	// file, header, err := c.Request.FormFile("upload")
	// filename := header.Filename
	// fmt.Println(header.Filename)
	// out, err := os.Create("./tmp/" + filename + ".png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer out.Close()
	// _, err = io.Copy(out, file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	file, header, err := c.Request.FormFile("upload")
	fmt.Println(header.Filename)
	filename := header.Filename
	fmt.Println(filename)
	fmt.Println(file)
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(config.IMAGE_PATH + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// 替換圖片檔名
	newFileName := GetMD5Hash(filename)
	dotIndex := strings.LastIndex(filename, ".") //取得最後的.的索引值
	if dotIndex != -1 && dotIndex != 0 {
		newFileName += filename[dotIndex:] //取得副檔名
	}
	err = os.Rename(config.IMAGE_PATH+filename+".png", config.IMAGE_PATH+newFileName+".png")
	//Rename(oldName,newName)
	if err != nil {
		fmt.Println("reName Error", err)
	}

	if len(user.Username) > 0 && len(user.Password) > 0 {

		id, err := user.Insert()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "添加失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "添加成功",
			"data":    id,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失敗",
		})
	}
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}
