package log

import (
	"bytes"
	"fmt"
	"testing"
)

type params struct {
	appName string
	options []Option
}

type expect struct {
	ok  bool
	err error
}

func TestLog(t *testing.T) {
	stdout := new(bytes.Buffer)
	l, err := NewLogger("TestApp", WithOutput(stdout))
	if err != nil {
		panic(err)
	}

	l.Log(I, "Test")
	fmt.Print(stdout.String())
	stdout.Reset()

	l.WithField("AddKey1", "AddVal1").Log(I, "Test")
	fmt.Print(stdout.String())
	stdout.Reset()

	l.WithField("AddKey2", "AddVal2").Log(I, "Test2")
	fmt.Print(stdout.String())
	stdout.Reset()

}

func TestNewLogger(t *testing.T) {

	factors := []struct {
		p params
		e expect
	}{
		{
			params{appName: "App01", options: nil},
			expect{ok: true, err: nil},
		},
	}

	for _, f := range factors {
		actual, err := NewLogger(f.p.appName, f.p.options...)
		if f.e.ok {
			if actual == nil {
				t.Error("actual is nil")
			}
			if err != nil {
				t.Error("err is not nil")
			}
		} else {
			if actual != nil {
				t.Error("actual is not nil")
			}
			if err == nil {
				t.Error("err is nil")
			}
		}
	}
}
