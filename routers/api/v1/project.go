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
func GetProject(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistProjectByID(id) {
			data = models.GetProject(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_Project
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


func GetProjects(c *gin.Context) {
	

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetProjects(util.GetPage(c), models.GetProjectTotal(maps), maps)
	data["total"] = models.GetProjectTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddProject(c *gin.Context) {
	name := c.Query("name")
	link := c.Query("link")
	theyear := c.Query("theyear")
	id := 0
	valid := validation.Validation{}
	valid.Required(name, "name").Message("项目名称不能为空")
	valid.MaxSize(name, 100, "name").Message("项目名称最长为100字符")
	valid.Required(link, "link").Message("项目链接不能为空")
	valid.MaxSize(link, 100, "link").Message("项目链接最长为100字符")
	valid.Required(theyear, "theyear").Message("项目年份不能为空")
	valid.MaxSize(theyear, 10, "theyear").Message("项目年份最长为10字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistProjectByName(id,name) {
			code = e.SUCCESS
			models.AddProject(name,link,theyear)
		} else {
			code = e.ERROR_EXIST_Project
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditProject(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	link := c.Query("link")
	theyear := c.Query("theyear")
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(name, 100, "name").Message("项目名称最长为100字符")
	valid.MaxSize(link, 100, "link").Message("项目链接最长为100字符")
	valid.Required(theyear, "theyear").Message("项目年份不能为空")
	valid.MaxSize(theyear, 10, "theyear").Message("项目年份最长为10字符")
	code := e.INVALID_PARAMS
	
		
		if !valid.HasErrors(){
			code = e.SUCCESS
			if models.ExistProjectByID(id) {
				if !models.ExistProjectByName(id,name){
					data := make(map[string]interface{})
				if link!= "" {
					data["link"] = link
				}
				if name != "" {
					data["name"] = name
				}
				data["theyear"] = theyear
				models.EditProject(id, data)
				}else{
					code = e.ERROR_EXIST_Project
				}
			} else {
				code = e.ERROR_NOT_EXIST_Project
			}
		}
	

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}


func DeleteProject(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistProjectByID(id) {
			models.DeleteProject(id)
		} else {
			code = e.ERROR_NOT_EXIST_Project
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}