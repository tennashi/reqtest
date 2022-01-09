package reqtest

import (
	"testing"
)

type OnFailure interface {
	Fail(mes string)
}

type TError struct {
	t *testing.T
}

func (n *TError) Fail(mes string) {
	n.t.Helper()

	n.t.Error(mes)
}

type TFatal struct {
	t *testing.T
}

func (n *TFatal) Fail(mes string) {
	n.t.Helper()

	n.t.Fatal(mes)
}
