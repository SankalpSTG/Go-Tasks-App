package misc

import "goapp/dao"

type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
type IntData struct {
	Data int64 `json:"data,omitempty"`
}
type TasksData struct {
	Data []dao.Task `json:"data,omitempty"`
}
