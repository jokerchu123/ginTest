package routers
//注册路由
import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	api"github.com/EDDYCJY/go-gin-example/routers/api"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//获取token
	r.GET("/manager", api.GetManager)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//新建新闻
		apiv1.POST("/news", v1.AddNews)
		//更新指定新闻
		apiv1.PUT("/news/:id", v1.EditNews)
		//删除指定新闻
		apiv1.DELETE("/news/:id", v1.DeleteNews)
		
		
		//新建文章
		apiv1.POST("/article", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)

		
		//新建项目
		apiv1.POST("/project", v1.AddProject)
		//更新指定项目
		apiv1.PUT("/project/:id", v1.EditProject)
		//删除指定项目
		apiv1.DELETE("/project/:id", v1.DeleteProject)
		
		//新建成员
		apiv1.POST("/member", v1.AddMember)
		//更新指定成员
		apiv1.PUT("/member/:id", v1.EditMember)
		//删除指定成员
		apiv1.DELETE("/member/:id", v1.DeleteMember)

		//新建图片
		apiv1.POST("/image", v1.AddImage)
		//更新指定图片
		apiv1.PUT("/image/:id", v1.EditImage)
		//删除指定图片
		apiv1.DELETE("/image/:id", v1.DeleteImage)

		//新建成果
		apiv1.POST("/achievement", v1.AddAchievement)
		//更新指定成果
		apiv1.PUT("/achievement/:id", v1.EditAchievement)
		//删除指定成果
		apiv1.DELETE("/achievement/:id", v1.DeleteAchievement)
	}
	apiget := r.Group("/api/get")
	{
		//获取新闻列表
		apiget.GET("/news", v1.GetNews)

		//获取项目列表
		apiget.GET("/projects", v1.GetProjects)
		//获取文章列表
		apiget.GET("/articles", v1.GetArticles)
		//获取成员列表
		apiget.GET("/members", v1.GetMembers)
		//获取图片列表
		apiget.GET("/images", v1.GetImages)
		//获取成果列表
		apiget.GET("/achievements", v1.GetAchievements)

		//获取指定文章
		apiget.GET("/article/:id", v1.GetArticle)
		//获取指定项目
		apiget.GET("/project/:id", v1.GetProject)
		//获取指定成员
		apiget.GET("/member/:id", v1.GetMember)
		//获取指定图片
		apiget.GET("/image/:id", v1.GetImage)
		//获取指定成果
		apiget.GET("/achievement/:id", v1.GetAchievement)
	}
	return r
}
