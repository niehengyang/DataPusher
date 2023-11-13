
package main

import (
	"dm/controller"
	"dm/service/config"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/net/context"
)
var adminConfigFile = flag.String("f", "etc/navi-business.yaml", "the config file")

func main() {

	var c config.Config
	conf.MustLoad(*adminConfigFile, &c)
	var ctx *context.Context

	c.AdminApi.MaxBytes = 1024 * 1024 * 50 //实体大小限制
	server := rest.MustNewServer(rest.RestConf(c.AdminApi))
	defer server.Stop()
	controller.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.AdminApi.Host, c.AdminApi.Port)
	server.Start()
}