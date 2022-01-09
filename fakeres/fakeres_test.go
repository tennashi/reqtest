package fakeres_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tennashi/reqtest/fakeres"
)

func TestJSONHandler(t *testing.T) {
	cases := []*fakeres.JSONHandler{
		{
			Body:       1,
			StatusCode: 200,
		},
		{
			Body:       "hoge",
			StatusCode: 200,
		},
		{
			Body:       []string{"hoge", "fuga"},
			StatusCode: 200,
		},
		{
			Body:       map[string]string{"hoge": "fuga"},
			StatusCode: 200,
		},
		{
			Body:       struct{ Hoge string }{Hoge: "fuga"},
			StatusCode: 200,
		},
		{
			Body:       struct{ Hoge string }{Hoge: "fuga"},
			StatusCode: 200,
		},
		{
			Body:       1,
			StatusCode: 201,
		},
		{
			Body:       1,
			StatusCode: 400,
		},
		{
			Body:       1,
			StatusCode: 200,
			Header:     map[string][]string{"X-Hoge": {"hoge"}},
		},
	}
	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			srv := httptest.NewServer(tt)
			defer srv.Close()

			res, err := http.Get(srv.URL)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			if diff := cmp.Diff(res.StatusCode, tt.StatusCode); diff != "" {
				t.Fatal(diff)
			}

			for k, vs := range tt.Header {
				if diff := cmp.Diff(res.Header.Values(k), vs); diff != "" {
					t.Fatal(diff)
				}
			}

			d, err := json.Marshal(tt.Body)
			if err != nil {
				t.Fatal(err)
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(bytes.TrimSpace(resBody)), string(bytes.TrimSpace(d))); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
