package config

import (
	"fmt"
	"log"
	"ms-products/src/middleware"
	"net/http"
)

func RunServer(port string, dependencies *ResolveDependencies) {
	routers := http.NewServeMux()
	routers = dependencies.ProductsRouter.RegisterRouters(routers)
	routers = dependencies.PingRouter.RegisterRouters(routers)

	serve := http.Server{
		Addr:    port,
		Handler: middleware.RequestLoggerMiddleware(routers),
	}
	fmt.Printf("Run server http://localhost%s\n", port)

	if err := serve.ListenAndServe(); err != nil {
		log.Fatal("Faild to start http server")
	}
}
