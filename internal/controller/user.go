package controller

import (
	"context"
	apiV1 "github.com/gogf/template-single/api/v1"
)

var User = userController{}

type userController struct{}

func (c *userController) Login(ctx context.Context, req *apiV1.AuthReq) (res *apiV1.AuthRes, err error) {
	return
}
