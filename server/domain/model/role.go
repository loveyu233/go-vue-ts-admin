package model

// Role 表示角色模型
type Role struct {
	ID   int    `json:"id"`   // 唯一 ID
	Name string `json:"name"` // 角色名称
}
