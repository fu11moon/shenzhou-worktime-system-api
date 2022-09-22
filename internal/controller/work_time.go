package controller

import (
	"context"
	apiV1 "github.com/gogf/template-single/api/v1"
)

var WorkTime = WorkTimeController{}

type WorkTimeController struct{}

func (c *WorkTimeController) getThisWeek(ctx context.Context, req *apiV1.GetThisWeekReq) (res *apiV1.GetThisWeekRes, err error) {
	return
}

func (c *WorkTimeController) getLastWeek(ctx context.Context, req *apiV1.GetThisWeekReq) (res *apiV1.GetThisWeekRes, err error) {
	return
}

func (c *WorkTimeController) updateThisWeek(ctx context.Context, req *apiV1.UpdateThisWeekReq) (res *apiV1.UpdateThisWeekRes, err error) {
	return
}

func (c *WorkTimeController) updateLastWeek(ctx context.Context, req *apiV1.UpdateLastWeekReq) (res *apiV1.UpdateLastWeekRes, err error) {
	return
}
