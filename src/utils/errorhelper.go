package utils

import "net/http"

func ErrorHandleHttp(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)

}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
