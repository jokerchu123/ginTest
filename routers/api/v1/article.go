package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
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

// 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}


	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), models.GetArticleTotal(maps), maps)
		data["total"] = models.GetArticleTotal(maps)

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

// 新增文章
func AddArticle(c *gin.Context) {
	title := c.Query("title")
	journal := c.Query("journal")
	author := c.Query("author")
	authors := c.Query("authors")
	date := c.Query("date")
	link := c.Query("link")
	papercode := c.Query("papercode")
	abstract := c.Query("abstract")
	theyear := c.Query("theyear")
	id := 0

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(journal, "journal").Message("期刊不能为空")
	valid.MaxSize(journal, 100, "journal").Message("期刊最长为100字符")
	valid.Required(author, "author").Message("第一作者不能为空")
	valid.MaxSize(author, 100, "author").Message("第一作者最长为100字符")
	valid.Required(authors, "authors").Message("其他作者不能为空")
	valid.MaxSize(authors, 100, "authors").Message("其他作者最长为100字符")
	valid.Required(date, "date").Message("日期不能为空")
	valid.MaxSize(date, 100, "date").Message("日期最长为100字符")
	valid.Required(link, "link").Message("详情页链接不能为空")
	valid.MaxSize(link, 100, "link").Message("详情页链接最长为100字符")
	valid.Required(papercode, "papercode").Message("代码链接不能为空")
	valid.MaxSize(papercode, 100, "papercode").Message("代码链接最长为100字符")
	valid.Required(abstract, "abstract").Message("摘要不能为空")
	valid.MaxSize(abstract, 65535, "abstract").Message("摘要最长为65535字符")
	valid.Required(theyear, "theyear").Message("论文年份不能为空")
	valid.MaxSize(theyear, 10, "theyear").Message("论文年份最长为10字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistArticleByTitle(id,title) {
			data := make(map[string]interface{})
			data["title"] = title
			data["journal"] = journal
			data["author"] = author
			data["authors"] = authors
			data["link"] = link
			data["papercode"] = papercode
			data["abstract"] = abstract
			data["date"] = date
			data["theyear"] = theyear

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	journal := c.Query("journal")
	author := c.Query("author")
	authors := c.Query("authors")
	date := c.Query("date")
	link := c.Query("link")
	papercode := c.Query("papercode")
	abstract := c.Query("abstract")
	theyear := c.Query("theyear")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(journal, "journal").Message("期刊不能为空")
	valid.MaxSize(journal, 100, "journal").Message("期刊最长为100字符")
	valid.Required(author, "author").Message("第一作者不能为空")
	valid.MaxSize(author, 100, "author").Message("第一作者最长为100字符")
	valid.Required(authors, "authors").Message("其他作者不能为空")
	valid.MaxSize(authors, 100, "authors").Message("其他作者最长为100字符")
	valid.Required(date, "date").Message("日期不能为空")
	valid.MaxSize(date, 100, "date").Message("日期最长为100字符")
	valid.Required(link, "link").Message("详情页链接不能为空")
	valid.MaxSize(link, 100, "link").Message("详情页链接最长为100字符")
	valid.Required(papercode, "papercode").Message("代码链接不能为空")
	valid.MaxSize(papercode, 100, "papercode").Message("代码链接最长为100字符")
	valid.Required(abstract, "abstract").Message("摘要不能为空")
	valid.MaxSize(abstract, 65535, "abstract").Message("摘要最长为65535字符")
	valid.Required(theyear, "theyear").Message("论文年份不能为空")
	valid.MaxSize(theyear, 10, "theyear").Message("论文年份最长为10字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
				data := make(map[string]interface{})
				if !models.ExistArticleByTitle(id,title){
					data["title"] = title
					data["journal"] = journal
					data["author"] = author
					data["authors"] = authors
					data["link"] = link
					data["papercode"] = papercode
					data["abstract"] = abstract
					data["date"] = date
					data["theyear"] = theyear
					models.EditArticle(id, data)
					code= e.SUCCESS
				}else{
					code = e.ERROR_EXIST_ARTICLE
				}
				
			
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
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
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
		"data": make(map[string]string),
	})
}
