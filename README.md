# reqtest
HTTP handlers for testing HTTP requests

[![Go Reference](https://pkg.go.dev/badge/github.com/tennashi/reqtest.svg)](https://pkg.go.dev/github.com/tennashi/reqtest)
[![CI](https://github.com/tennashi/reqtest/workflows/test/badge.svg)](https://github.com/tennashi/reqtest/actions)

## Usage
```go
package some_test

import (
	"testing"
	"github.com/tennashi/reqtest"
)

func TestSomeRequest(t *testing.T) {
	srv := httptest.NewServer(reqtest.CompareMethodHandler(t, "GET"))
	defer srv.Close()

	res, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Close()
}
```
