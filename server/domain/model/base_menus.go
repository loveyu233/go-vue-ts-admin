package model

type SysBaseMenus struct {
	ID        int             `json:"id"`         // 唯一id
	RID       int             `json:"r_id"`       // 路由id
	MenuLevel int             `json:"menu_level"` // 路由显示层级
	ParentID  int             `json:"parent_id"`  // 父路由Id
	Path      string          `json:"path"`       // 访问路径
	Name      string          `json:"name"`       // 路由name
	Hidden    int             `json:"hidden"`     // 该路由是否隐藏
	Component string          `json:"component"`  // 对应前端的vue文件名
	Sort      int             `json:"sort"`       // 同一层级下的显示顺序
	Title     string          `json:"title"`      // 路由标题
	Icon      string          `json:"icon"`       // 路由图标
	Children  []*SysBaseMenus `json:"children"`   // 该路由下的子路由
}
