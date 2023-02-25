package e
//api错误码判断
var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_News:                "该新闻已存在",
	ERROR_NOT_EXIST_News:            "该新闻不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该论文不存在",
	ERROR_EXIST_ARTICLE:        	"该论文已存在",
	ERROR_NOT_EXIST_Project:        "该项目不存在",
	ERROR_EXIST_Project:        	"该项目已存在",
	ERROR_NOT_EXIST_Image:        "该图片不存在",
	ERROR_EXIST_Image:        	"该图片已存在",
	ERROR_NOT_EXIST_Member:        "该成员不存在",
	ERROR_EXIST_Member:        	"该成员已存在",
	ERROR_NOT_EXIST_Achievement:        "该成果不存在",
	ERROR_EXIST_Achievement:        	"该成果已存在",
	ERROR_MANAGER_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_MANAGER_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_MANAGER_TOKEN:               "Token生成失败",
	ERROR_MANAGER:                     "Token错误",
}
//??????
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
