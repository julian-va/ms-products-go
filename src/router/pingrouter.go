package router

import (
	"fmt"
	"ms-products/src/controller"
	"net/http"
)

type PingRouter struct {
	pingController controller.PingController
}

func NewPingRouter(pingController controller.PingController) *PingRouter {
	return &PingRouter{pingController: pingController}
}

func (router *PingRouter) RegisterRouters(routers *http.ServeMux) *http.ServeMux {
	routers.HandleFunc(fmt.Sprintf("%s /ping", http.MethodGet), router.pingController.Ping)
	return routers
}
