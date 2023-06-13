package model

// Button 表示按钮模型
type Button struct {
	ID   int    `json:"id"`   // 唯一 ID
	Name string `json:"name"` // 按钮名称
}
