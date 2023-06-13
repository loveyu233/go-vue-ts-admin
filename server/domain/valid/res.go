package valid

// ResponseCode 表示响应的状态码
type ResponseCode int

// 定义常量表示不同的响应状态码
const (
	SuccessCode               ResponseCode = 200
	UserNotFound                           = 301
	ParameterValidationFailed              = 302
	TokenValidationFailed                  = 303
	TokenDoesNotExist                      = 304
	InvalidRequestCode        ResponseCode = 400
	UnauthorizedCode          ResponseCode = 401
	ForbiddenCode             ResponseCode = 403
	NotFoundCode              ResponseCode = 404
	ServerErrorCode           ResponseCode = 500
	UnprocessableEntity       ResponseCode = 422
	TooManyRequestsCode       ResponseCode = 429
	ServiceUnavailable        ResponseCode = 503
	IconUpdateErr             ResponseCode = 504
	UserInfoUpdateErr         ResponseCode = 505
	MenuInfoUpdateErr                      = 506
	CarouselsGetErr                        = 507
)

// ResponseMessage 表示响应的状态说明
var ResponseMessage = map[ResponseCode]string{
	SuccessCode:               "成功",
	UserNotFound:              "用户名或密码错误",
	ParameterValidationFailed: "参数校验错误",
	TokenValidationFailed:     "token校验失败",
	TokenDoesNotExist:         "token不存在",
	InvalidRequestCode:        "无效请求",
	UnauthorizedCode:          "未授权",
	ForbiddenCode:             "禁止访问",
	NotFoundCode:              "未找到",
	ServerErrorCode:           "内部服务器错误",
	UnprocessableEntity:       "无法处理的实体",
	TooManyRequestsCode:       "请求过多",
	ServiceUnavailable:        "服务不可用",
	IconUpdateErr:             "用户头像上传失败",
	UserInfoUpdateErr:         "用户信息更新失败",
	MenuInfoUpdateErr:         "路由修改失败",
	CarouselsGetErr:           "轮播图获取失败",
}
