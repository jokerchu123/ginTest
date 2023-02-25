package v1
import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)
func GetNews(c *gin.Context) {
	

	maps := make(map[string]interface{})
	data := make(map[string]interface{})


	code := e.SUCCESS

	data["lists"] = models.GetNews(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetNewsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddNews(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	valid := validation.Validation{}
	id := 0
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistNewsByTitle(id,title) {
			code = e.SUCCESS
			models.AddNews(title,content)
		} else {
			code = e.ERROR_EXIST_News
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditNews(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	content := c.Query("content")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	
		
		if !valid.HasErrors(){
			code = e.SUCCESS
			if models.ExistNewsByID(id) {
				if !models.ExistNewsByTitle(id,title){
					data := make(map[string]interface{})
				if content!= "" {
					data["content"] = content
				}
				if title != "" {
					data["title"] = title
				}
				models.EditNews(id, data)
				}else{
					code = e.ERROR_EXIST_News
				}
			} else {
				code = e.ERROR_NOT_EXIST_News
			}
		}
	

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}


func DeleteNews(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistNewsByID(id) {
			models.DeleteNews(id)
		} else {
			code = e.ERROR_NOT_EXIST_News
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}