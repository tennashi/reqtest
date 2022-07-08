package reqtest

import (
	"net/http"
	"net/url"
	"testing"
)

func CompareMethodHandler(t *testing.T, method string) http.Handler {
	t.Helper()

	return (&HandlerGenerator{
		OnFailure: &TError{t: t},
	}).CompareMethod(method)
}

func ComparePathHandler(t *testing.T, path string) http.Handler {
	t.Helper()

	return (&HandlerGenerator{
		OnFailure: &TError{t: t},
	}).ComparePath(path)
}

func CompareQueryHandler(t *testing.T, want url.Values) http.Handler {
	t.Helper()

	return (&HandlerGenerator{
		OnFailure: &TError{t: t},
	}).CompareQuery(want)
}

func CompareHeaderValuesHandler(t *testing.T, key string, values []string) http.Handler {
	t.Helper()

	return (&HandlerGenerator{
		OnFailure: &TError{t: t},
	}).CompareHeaderValues(key, values)
}

func CompareJSONBodyHandler(t *testing.T, params interface{}) http.Handler {
	t.Helper()

	return (&HandlerGenerator{
		OnFailure: &TError{t: t},
	}).CompareJSONBody(params)
}
