package config

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port string, dependencies *ResolveDependencies) {
	routers := http.NewServeMux()
	routers = dependencies.ProductsRouter.RegisterRouters(routers)
	routers = dependencies.PingRouter.RegisterRouters(routers)
	requestLoggerMiddleware(routers)

	serve := http.Server{
		Addr:    port,
		Handler: requestLoggerMiddleware(routers),
	}
	fmt.Printf("Run server http://localhost%s", port)

	if err := serve.ListenAndServe(); err != nil {
		log.Fatal("Faild to start http server")
	}
}

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Header %s, method %s, path: %s, body: %s\n", r.Header, r.Method, r.URL.Path, r.Body)
		next.ServeHTTP(w, r)
	}
}
