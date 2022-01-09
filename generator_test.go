package reqtest_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHandlerGenerator_CompareMethod(t *testing.T) {
	t.Parallel()

	httpMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPut,
		http.MethodPost,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	cases := append(httpMethods, []string{
		"hoge",
		"",
	}...)

	for _, tt := range cases {
		t.Run(tt, func(t *testing.T) {
			t.Parallel()

			failMesCh := make(chan string, 1)
			defer close(failMesCh)

			got := HandlerGeneratorForTest(failMesCh).CompareMethod(tt)
			srv := httptest.NewServer(got)
			defer srv.Close()

			for _, reqMethod := range httpMethods {
				req, err := http.NewRequest(reqMethod, srv.URL, nil)
				if err != nil {
					t.Fatal(err)
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer res.Body.Close()

				if tt == reqMethod {
					if diff := cmp.Diff(res.StatusCode, http.StatusOK); diff != "" {
						t.Fatal(diff)
					}

					select {
					case mes := <-failMesCh:
						t.Fatalf("expect the cause of the failure will not be output, but got: %s", mes)
					default:
					}

					return
				}

				select {
				case <-failMesCh:
				default:
					t.Fatal("expect the cause of the failure to be output, but it was not")
				}
			}
		})
	}
}

func TestHandlerGenerator_ComparePath(t *testing.T) {
	t.Parallel()

	cases := []string{
		"/",
		"/hoge",
		"/hoge/fuga",
		"/hoge//fuga",
	}

	for _, tt := range cases {
		t.Run(tt, func(t *testing.T) {
			t.Parallel()

			failMesCh := make(chan string, 1)
			defer close(failMesCh)

			got := HandlerGeneratorForTest(failMesCh).ComparePath(tt)
			srv := httptest.NewServer(got)
			defer srv.Close()

			u, err := url.Parse(srv.URL)
			if err != nil {
				t.Fatal(err)
			}

			u.Path = tt

			t.Run("matching path", func(t *testing.T) {
				res, err := http.Get(u.String())
				if err != nil {
					t.Fatal(err)
				}
				defer res.Body.Close()

				if diff := cmp.Diff(res.StatusCode, 200); diff != "" {
					t.Fatal(diff)
				}

				select {
				case mes := <-failMesCh:
					t.Fatalf("expect the cause of the failure will not be output, but got: %s", mes)
				default:
				}
			})

			t.Run("mismatched path", func(t *testing.T) {
				res, err := http.Get(u.String() + "/mismatch")
				if err != nil {
					t.Fatal(err)
				}
				defer res.Body.Close()

				select {
				case <-failMesCh:
				default:
					t.Fatal("expect the cause of the failure to be output, but it was not")
				}
			})
		})
	}
}

func TestHandlerGenerator_CompareHeaderValues(t *testing.T) {
	t.Parallel()

	cases := [][]string{
		{"some", "value"},
		{"value", "some"},
		{"some"},
		{"value"},
		{"mismatch"},
		{""},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			failMesCh := make(chan string, 1)
			defer close(failMesCh)

			got := HandlerGeneratorForTest(failMesCh).CompareHeaderValues("X-Some-Header", tt)
			srv := httptest.NewServer(got)
			defer srv.Close()

			t.Run("matching key", func(t *testing.T) {
				req, err := http.NewRequest("GET", srv.URL, nil)
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Add("X-Some-Header", "some")
				req.Header.Add("X-Some-Header", "value")

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer res.Body.Close()

				if reflect.DeepEqual(req.Header.Values("X-Some-Header"), tt) {
					select {
					case mes := <-failMesCh:
						t.Fatalf("expect the cause of the failure will not be output, but got: %s", mes)
					default:
					}

					return
				}

				select {
				case <-failMesCh:
				default:
					t.Fatal("expect the cause of the failure to be output, but it was not")
				}
			})

			t.Run("mismatched key", func(t *testing.T) {
				req, err := http.NewRequest("GET", srv.URL, nil)
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Add("X-Mismatched-Header", "some")
				req.Header.Add("X-Mismatched-Header", "value")

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer res.Body.Close()

				select {
				case <-failMesCh:
				default:
					t.Fatal("expect the cause of the failure to be output, but it was not")
				}
			})
		})
	}
}
