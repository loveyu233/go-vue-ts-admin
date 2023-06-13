package vo

type UserVo struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Avatar   string     `json:"avatar"`
	Identity string     `json:"identity"`
	Routers  []RouteVo  `json:"routers"`
	Buttons  []ButtonVo `json:"buttons"`
	Roles    []RoleVo   `json:"roles"`
}
