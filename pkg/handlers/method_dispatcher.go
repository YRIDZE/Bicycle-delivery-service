package handlers

import "net/http"

type Methods struct {
	get    http.Handler
	post   http.Handler
	put    http.Handler
	delete http.Handler
}

func (h *AppHandlers) MethodDispatcher(methods Methods) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set(
				"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,  Special-Request-Header, Authorization",
			)

			switch req.Method {
			case "GET":
				reqHandler(methods.get, w, req)
			case "POST":
				reqHandler(methods.post, w, req)
			case "PUT":
				reqHandler(methods.put, w, req)
			case "DELETE":
				reqHandler(methods.delete, w, req)
			case "OPTIONS":
				w.WriteHeader(http.StatusOK)
				return
			}
		},
	)
}

func reqHandler(h http.Handler, w http.ResponseWriter, req *http.Request) {
	if h != nil {
		h.ServeHTTP(w, req)
	} else {
		http.Error(w, "request not processed", http.StatusMethodNotAllowed)
	}
}
