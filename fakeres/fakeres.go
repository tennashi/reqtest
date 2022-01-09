package fakeres

import (
	"encoding/json"
	"net/http"
)

type JSONHandler struct {
	Body       interface{}
	Header     map[string][]string
	StatusCode int
}

func NewJSONHandler(body interface{}) *JSONHandler {
	return &JSONHandler{
		Body:       body,
		StatusCode: http.StatusOK,
	}
}

func (h *JSONHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	for k, vs := range h.Header {
		for _, v := range vs {
			header.Set(k, v)
		}
	}

	w.WriteHeader(h.StatusCode)

	if err := json.NewEncoder(w).Encode(h.Body); err != nil {
		// There's nothing I can do.
		panic(err)
	}
}

type BytesHandler struct {
	Body       []byte
	Header     map[string][]string
	StatusCode int
}

func (h *BytesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	for k, vs := range h.Header {
		for _, v := range vs {
			header.Set(k, v)
		}
	}

	w.WriteHeader(h.StatusCode)

	_, err := w.Write(h.Body)
	if err != nil {
		// There's nothing I can do.
		panic(err)
	}
}
