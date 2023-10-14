// Code generated by hertz generator.

package calculate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/XZ0730/tireCV/biz/dal/cache"
	calculate "github.com/XZ0730/tireCV/biz/model/calculate"
	"github.com/XZ0730/tireCV/biz/pack"
	"github.com/XZ0730/tireCV/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

// Calculate .
// @router /cs/calculate [POST]
func Calculate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculate.CalculateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	klog.Info("req :", req.Expression)
	if req.Expression == "" {
		pack.SendFailResponse(c, errors.New("expression not found"))
		return
	}
	resp := new(calculate.CalculateResponse)
	result, err := service.NewDo_CalculateService(ctx).BaseCalculate(ctx, req.Expression)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	err = cache.SetHistory(ctx, c.ClientIP(), fmt.Sprint(time.Now().Unix()), req.Expression+"="+result)
	if err != nil {
		pack.SetCalculateResponse(30001, err.Error(), result, resp)
		c.JSON(consts.StatusOK, resp)
		return
	}
	pack.SetCalculateResponse(consts.StatusOK, consts.StatusMessage(consts.StatusOK), result, resp)
	c.JSON(consts.StatusOK, resp)
}

// RateCall .
// @router /cs/rate [POST]
func RateCall(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculate.RateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(calculate.RateResponse)

	interest, err := service.NewDo_CalculateService(ctx).RateCalculate(ctx, &req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	pack.SetRateResponse(consts.StatusOK, consts.StatusMessage(consts.StatusOK), interest, resp)
	c.JSON(consts.StatusOK, resp)
}

// SetRate .
// @router /cs/rate/set [POST]
func SetRate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculate.SetRateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(calculate.BaseResponse)
	err = service.NewDo_CalculateService(ctx).SetRate(ctx, &req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	pack.SetBaseResponse(consts.StatusOK, consts.StatusMessage(consts.StatusOK), resp)
	c.JSON(consts.StatusOK, resp)
}

// GetHistory .
// @router /cs/calculate/history [GET]
func GetHistory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculate.HistoryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp := new(calculate.HistoryResponse)
	key := c.ClientIP()
	// klog.Info("key:", key)
	history, err := service.NewDo_CalculateService(ctx).History(ctx, key)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	if len(history) > 10 {
		history = history[len(history)-9:]
	}
	pack.SetHistoryResponse(consts.StatusOK, consts.StatusMessage(consts.StatusOK), history, resp)
	c.JSON(consts.StatusOK, resp)
}

// GetRate .
// @router /cs/rate/get [GET]
func GetRate(ctx context.Context, c *app.RequestContext) {
	var err error

	resp := new(calculate.GetRateResponse)
	res, err := service.NewDo_CalculateService(ctx).GetRate(ctx)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	pack.SetGetRateResponse(consts.StatusOK, consts.StatusMessage(consts.StatusOK), res, resp)
	c.JSON(consts.StatusOK, resp)
}