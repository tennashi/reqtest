package reqtest

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-cmp/cmp"
)

type HandlerGenerator struct {
	OnFailure OnFailure
}

func (g *HandlerGenerator) CompareMethod(method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if diff := cmp.Diff(req.Method, method); diff != "" {
			g.OnFailure.Fail(diff)
			return
		}
	})
}

func (g *HandlerGenerator) ComparePath(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if diff := cmp.Diff(req.URL.Path, path); diff != "" {
			g.OnFailure.Fail(diff)
			return
		}
	})
}

func (g *HandlerGenerator) CompareHeaderValues(key string, values []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if diff := cmp.Diff(req.Header.Values(key), values); diff != "" {
			g.OnFailure.Fail(diff)
			return
		}
	})
}

func (g *HandlerGenerator) CompareJSONBody(jsonBody interface{}) http.Handler {
	got := jsonBody

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(got); err != nil {
			g.OnFailure.Fail(err.Error())
			return
		}

		if diff := cmp.Diff(got, jsonBody); diff != "" {
			g.OnFailure.Fail(diff)
			return
		}
	})
}
