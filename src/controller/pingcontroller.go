package controller

import (
	"ms-products/src/utils"
	"net/http"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (pingController *PingController) Ping(w http.ResponseWriter, req *http.Request) {
	pong := make(map[string]string)
	pong["ping"] = "pong"
	utils.CreateResponse(w, pong, http.StatusOK)
}
