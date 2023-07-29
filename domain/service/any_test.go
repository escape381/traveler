package service

import (
	"testing"
)

func TestAnyFunc(t *testing.T) {
	any := AnyService{}
	actual := any.AnyFunc()
	expect := "this is AnyFunc!!"
	if actual != expect {
		t.Errorf("I expected 'this is AnyFunc!!', but the result was '%s'", actual)
	}
}
