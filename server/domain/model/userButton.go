package model

// UserButton 表示用户按钮关联模型
type UserButton struct {
	ID       int `json:"id"`        // 唯一 ID
	UserID   int `json:"user_id"`   // 关联User表的id
	ButtonID int `json:"button_id"` // 关联Button表的id
}
