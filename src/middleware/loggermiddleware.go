package middleware

import (
	"log"
	"ms-products/src/utils"
	"net/http"
	"strings"
)

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestId := r.Header.Get(utils.X_REQUEST_ID)
		w.Header().Add(utils.X_REQUEST_ID, requestId)
		path := r.URL.Path

		next.ServeHTTP(w, r)
		if !strings.EqualFold(path, utils.PING_PATH) {
			log.Printf("X-Request-Id: %s, Method: %s, Path: %s, Status-Code: %s\n", requestId, r.Method, r.URL.Path, w.Header().Get(utils.HTTP_STATUS))
		}
	}
}
