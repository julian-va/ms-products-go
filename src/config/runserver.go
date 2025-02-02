package config

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port string, dependencies *ResolveDependencies) {
	routers := http.NewServeMux()

	routers = dependencies.ProductsRouter.RegisterRouters(routers)
	fmt.Printf("Run server http://localhost%s", port)

	if err := http.ListenAndServe(port, routers); err != nil {
		log.Fatal("Faild to start http server")
	}
}
