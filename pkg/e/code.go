package e
//api错误码
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_News         = 10001
	ERROR_NOT_EXIST_News     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003
	ERROR_EXIST_ARTICLE = 10004
	ERROR_NOT_EXIST_Project = 10005
	ERROR_EXIST_Project = 10006
	ERROR_NOT_EXIST_Image = 10007
	ERROR_EXIST_Image = 10008
	ERROR_NOT_EXIST_Member = 10009
	ERROR_EXIST_Member = 10010
	ERROR_NOT_EXIST_Achievement = 10011
	ERROR_EXIST_Achievement = 10012

	ERROR_MANAGER_CHECK_TOKEN_FAIL    = 20001
	ERROR_MANAGER_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_MANAGER_TOKEN               = 20003
	ERROR_MANAGER                     = 20004
)
