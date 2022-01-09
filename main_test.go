package reqtest_test

import (
	"github.com/tennashi/reqtest"
)

type OnFailureForTest struct {
	mesCh chan<- string
}

func (f *OnFailureForTest) Fail(failMes string) {
	f.mesCh <- failMes
}

func HandlerGeneratorForTest(mesCh chan<- string) *reqtest.HandlerGenerator {
	return &reqtest.HandlerGenerator{
		OnFailure: &OnFailureForTest{
			mesCh: mesCh,
		},
	}
}
