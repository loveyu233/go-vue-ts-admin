package model

// User 表示用户模型
type User struct {
	ID       int    `json:"id"`       // 唯一 ID
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Avatar   string `json:"avatar"`   // 头像图片路径
	Identity string `json:"identity"` // 身份说明
}
