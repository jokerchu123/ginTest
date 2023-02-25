package v1
import (
	"net/http"
	"log"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

func GetImage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistImageByID(id) {
			data = models.GetImage(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}


func GetImages(c *gin.Context) {
	

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetImages(util.GetPage(c), models.GetImageTotal(maps), maps)
	data["total"] = models.GetImageTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddImage(c *gin.Context) {
	name := c.Query("name")
	date := c.Query("date")
	address := c.Query("address")
	id := 0
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(date, "date").Message("日期不能为空")
	valid.MaxSize(date, 100, "date").Message("日期最长为100字符")
	valid.Required(address, "address").Message("地址不能为空")
	valid.MaxSize(address, 65535, "address").Message("地址最长为65535字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistImageByName(id,name) {
			code = e.SUCCESS
			models.AddImage(name,date,address)
		} else {
			code = e.ERROR_EXIST_Image
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditImage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	date := c.Query("date")
	address := c.Query("address")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(name, "name").Message("图片名称不能为空")
	valid.MaxSize(name, 100, "name").Message("图片名称最长为100字符")
	valid.Required(date, "date").Message("图片日期不能为空")
	valid.MaxSize(date, 100, "date").Message("图片日期最长为100字符")
	valid.Required(address, "address").Message("图片地址不能为空")
	valid.MaxSize(address, 65535, "address").Message("图片地址年份最长为65535字符")
	code := e.INVALID_PARAMS
	
		
		if !valid.HasErrors(){
			code = e.SUCCESS
			if models.ExistImageByID(id) {
				if !models.ExistImageByName(id,name){
					data := make(map[string]interface{})
				if date!= "" {
					data["date"] = date
				}
				if name != "" {
					data["name"] = name
				}
				data["address"] = address
				models.EditImage(id, data)
				}else{
					code = e.ERROR_EXIST_Image
				}
			} else {
				code = e.ERROR_NOT_EXIST_Image
			}
		}
	

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}


func DeleteImage(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistImageByID(id) {
			models.DeleteImage(id)
		} else {
			code = e.ERROR_NOT_EXIST_Image
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}