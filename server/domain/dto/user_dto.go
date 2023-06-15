package dto

type UserDto struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type UserUpdateDto struct {
	Uid      int    `json:"uid,omitempty"`
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty"`
	Icon     string `json:"image,omitempty" binding:"required"`
	Token    string `json:"token,omitempty" binding:"required"`
}
