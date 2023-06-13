package model

// UserRole 表示用户角色关联模型
type UserRole struct {
	ID     int `json:"id"`      // 唯一 ID
	UserID int `json:"user_id"` // 关联User表的id
	RoleID int `json:"role_id"` // 关联Role表的id
}
