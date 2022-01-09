package reqtest

import "net/http"

func ChainHandler(handlers ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, h := range handlers {
			h.ServeHTTP(w, req)
		}
	})
}
