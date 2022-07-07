package httpMethods

import (
	"fmt"
	"net/http"
	"test_go_webserver/http/httpStatus"
)

const (
	GET     string = http.MethodGet
	POST    string = http.MethodPost
	PUT     string = http.MethodPut
	DELETE  string = http.MethodDelete
	PATCH   string = http.MethodPatch
	OPTIONS string = http.MethodOptions
	CONNECT string = http.MethodConnect
	HEAD    string = http.MethodHead
	TRACE   string = http.MethodTrace
)

func HandleGet(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != GET {
			message := "400: Bad Request"

			w.WriteHeader(httpStatus.BadRequest)
			_, _ = w.Write([]byte(message))

			println(fmt.Errorf("%s, %s", r.Method, message).Error())

			return
		}

		f(w, r)
	})
}

func HandlePost(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != POST {
			message := "400: Bad Request"

			w.WriteHeader(httpStatus.BadRequest)
			_, _ = w.Write([]byte(message))

			println(fmt.Errorf("%s, %s", r.Method, message).Error())

			return
		}

		f(w, r)
	})
}

func HandlePut(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != PUT {
			message := "400: Bad Request"

			w.WriteHeader(httpStatus.BadRequest)
			_, _ = w.Write([]byte(message))

			println(fmt.Errorf("%s, %s", r.Method, message).Error())

			return
		}

		f(w, r)
	})
}

func HandleDelete(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != DELETE {
			message := "400: Bad Request"

			w.WriteHeader(httpStatus.BadRequest)
			_, _ = w.Write([]byte(message))

			println(fmt.Errorf("%s, %s", r.Method, message).Error())

			return
		}

		f(w, r)
	})
}
