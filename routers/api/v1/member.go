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


func GetMember(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistMemberByID(id) {
			data = models.GetMember(id)
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


func GetMembers(c *gin.Context) {
	

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	code := e.SUCCESS

	data["lists"] = models.GetMembers(util.GetPage(c), models.GetMemberTotal(maps), maps)
	data["total"] = models.GetMemberTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddMember(c *gin.Context) {
	name := c.Query("name")
	identity := c.Query("identity")
	phone := c.Query("phone")
	mail := c.Query("mail")
	achievement := c.Query("achievement")
	research := c.Query("research")
	introduction := c.Query("introduction")
	id := 0
	valid := validation.Validation{}
	valid.Required(name, "name").Message("姓名不能为空")
	valid.MaxSize(name, 100, "name").Message("姓名最长为100字符")
	valid.Required(identity, "identity").Message("身份不能为空")
	valid.MaxSize(identity, 100, "identity").Message("身份最长为100字符")
	valid.Required(phone, "phone").Message("电话年份不能为空")
	valid.MaxSize(phone, 100, "phone").Message("电话年份最长为10字符")
	valid.Required(research, "research").Message("研究方向不能为空")
	valid.MaxSize(research, 100, "research").Message("研究方向最长为100字符")
	valid.Required(mail, "mail").Message("邮箱不能为空")
	valid.MaxSize(mail, 100, "mail").Message("邮箱最长为100字符")
	valid.Required(achievement, "achievement").Message("成果不能为空")
	valid.MaxSize(achievement, 65535, "achievement").Message("成果最长为65535字符")
	valid.Required(introduction, "introduction").Message("个人介绍不能为空")
	valid.MaxSize(introduction, 65535, "introduction").Message("个人介绍最长为65535字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistMemberByName(id,name) {
			code = e.SUCCESS
			models.AddMember(name,phone,achievement,mail,research,identity,introduction)
		} else {
			code = e.ERROR_EXIST_Member
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditMember(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	identity := c.Query("identity")
	phone := c.Query("phone")
	mail := c.Query("mail")
	achievement := c.Query("achievement")
	research := c.Query("research")
	introduction := c.Query("introduction")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("姓名不能为空")
	valid.MaxSize(name, 100, "name").Message("姓名最长为100字符")
	valid.Required(research, "research").Message("研究方向不能为空")
	valid.MaxSize(research, 100, "research").Message("研究方向最长为100字符")
	valid.Required(identity, "identity").Message("身份不能为空")
	valid.MaxSize(identity, 100, "identity").Message("身份最长为100字符")
	valid.Required(phone, "phone").Message("电话年份不能为空")
	valid.MaxSize(phone, 100, "phone").Message("电话年份最长为10字符")
	valid.Required(mail, "mail").Message("邮箱不能为空")
	valid.MaxSize(mail, 100, "mail").Message("邮箱最长为100字符")
	valid.Required(achievement, "achievement").Message("成果不能为空")
	valid.MaxSize(achievement, 65535, "achievement").Message("成果最长为65535字符")
	valid.Required(introduction, "introduction").Message("个人介绍不能为空")
	valid.MaxSize(introduction, 65535, "introduction").Message("个人介绍最长为65535字符")
	code := e.INVALID_PARAMS
	
		
		if !valid.HasErrors(){
			code = e.SUCCESS
			if models.ExistMemberByID(id) {
				if !models.ExistMemberByName(id,name){
					data := make(map[string]interface{})
				
				data["name"] = name
				data["phone"] = phone
				data["mail"] = mail
				data["identity"] = identity
				data["achievement"] = achievement
				data["introduction"] = introduction
				data["research"] = research
				
				models.EditMember(id, data)
				}else{
					code = e.ERROR_EXIST_Member
				}
			} else {
				code = e.ERROR_NOT_EXIST_Member
			}
		}
	

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}


func DeleteMember(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistMemberByID(id) {
			models.DeleteMember(id)
		} else {
			code = e.ERROR_NOT_EXIST_Member
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}