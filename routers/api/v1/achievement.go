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

func GetAchievement(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistAchievementByID(id) {
			data = models.GetAchievement(id)
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

func GetAchievements(c *gin.Context) {
	

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetAchievements(util.GetPage(c), models.GetAchievementTotal(maps), maps)
	data["total"] = models.GetAchievementTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddAchievement(c *gin.Context) {
	name := c.Query("name")
	address := c.Query("address")
	category := c.Query("category")
	id := 0
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(category, "category").Message("类别不能为空")
	valid.MaxSize(category, 100, "category").Message("类别最长为100字符")
	valid.Required(address, "address").Message("地址不能为空")
	valid.MaxSize(address, 100, "address").Message("地址最长为100字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistAchievementByName(id,name) {
			code = e.SUCCESS
			models.AddAchievement(name,address,category)
		} else {
			code = e.ERROR_EXIST_Achievement
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditAchievement(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	address := c.Query("address")
	category := c.Query("category")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.MaxSize(address, 100, "address").Message("地址最长为100字符")
	valid.MaxSize(category, 100, "category").Message("类别最长为100字符")
	code := e.INVALID_PARAMS
	
		
		if !valid.HasErrors(){
			code = e.SUCCESS
			if models.ExistAchievementByID(id) {
				if !models.ExistAchievementByName(id,name){
					data := make(map[string]interface{})
				if address!= "" {
					data["address"] = address
				}
				if name != "" {
					data["name"] = name
				}
				if category != "" {
					data["category"] = category
				}
				models.EditAchievement(id, data)
				}else{
					code = e.ERROR_EXIST_Achievement
				}
			} else {
				code = e.ERROR_NOT_EXIST_Achievement
			}
		}
	

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}


func DeleteAchievement(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistAchievementByID(id) {
			models.DeleteAchievement(id)
		} else {
			code = e.ERROR_NOT_EXIST_Achievement
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}