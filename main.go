package main

import (
	"context"
	"fmt"
	
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
)

func main() {
	db.Init()
	fmt.Println("DB init success")

	h := server.Default(server.WithHostPorts("0.0.0.0:18000"))
	
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.Spin()
}