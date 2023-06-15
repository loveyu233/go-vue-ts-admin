package dto

type MenuUpdate struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Icon  string `json:"image,omitempty"`
}
