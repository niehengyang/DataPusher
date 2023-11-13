package controller

import (
	"dm/service/resp"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/net/context"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *context.Context) {

	//错误响应处理
	resp.SetErrorHandler()

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/testSubmit",
				Handler: TestSubmit(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}