package main

import (
	"flag"

	"MuXiFresh-Be-2.0/common/nacos"

	"github.com/zeromicro/go-zero/gateway"
)

func main() {
	flag.Parse()

	var c gateway.GatewayConf
	nacos.MustLoad(nacos.Service("gateway", &c))

	gw := gateway.MustNewServer(c)
	defer gw.Stop()

	gw.Start()
}
