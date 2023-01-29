// Code generated by hertz generator.

package api

import (
	"context"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login [POST]
// 使用 hertz-jwt 时需要替换掉这个 LoginHandler
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUserById .
// @router /douyin/user [GET]
func GetUserById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}