package model

// UserRoute 表示用户路由关联模型
type UserRoute struct {
	ID      int `json:"id"`       // 唯一 ID
	UserID  int `json:"user_id"`  // 关联User表的id
	RouteID int `json:"route_id"` // 关联Route表的id
}
