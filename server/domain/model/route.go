package model

// Route 表示路由模型
type Route struct {
	ID   int    `json:"id"`   // 唯一 ID
	Name string `json:"name"` // 路由名称
}
