package reqtest

import (
	"net/http"
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
