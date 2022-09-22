package v1

import "github.com/gogf/gf/v2/frame/g"

type AuthReq struct {
	g.Meta   `path:"/auth" method:"post" summary:"登录授权"`
	UserName string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}
type AuthRes struct{}

type GetThisWeekReq struct {
	g.Meta `path:"/this-week" method:"get" summary:"获取本周工时"`
	Name   string `json:"name"`
}
type GetThisWeekRes struct {
}

type GetLastWeekReq struct {
	g.Meta `path:"/last-week" method:"get" summary:"获取上周工时"`
	Name   string `json:"name"`
}
type GetLastWeekRes struct{}

type UpdateThisWeekReq struct {
	g.Meta `path:"/this-week" method:"post" summary:"更新本周工时"`
}
type UpdateThisWeekRes struct{}

type UpdateLastWeekReq struct {
	g.Meta `path:"/last-week" method:"post" summary:"更新上周工时"`
}
type UpdateLastWeekRes struct{}
