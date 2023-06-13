package response

import "server/domain/valid"

// APIResponse 表示 API 的响应结构体
type APIResponse struct {
	Code    valid.ResponseCode `json:"code,omitempty" `   // 表示请求是否成功
	Message string             `json:"message,omitempty"` // 返回的消息
	Data    interface{}        `json:"data,omitempty"`    // 返回的数据
}

// NewSuccessResponse 创建一个成功的 API 响应
func NewSuccessResponse(data interface{}) *APIResponse {
	return &APIResponse{
		Code: 200,
		Data: data,
	}
}

// NewErrorResponse 创建一个错误的 API 响应
func NewErrorResponse(code valid.ResponseCode) *APIResponse {
	return &APIResponse{
		Code:    code,
		Message: valid.ResponseMessage[code],
	}
}
