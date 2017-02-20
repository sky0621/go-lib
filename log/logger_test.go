package log

import "testing"

type params struct {
	appName string
	options []Option
}

type expect struct {
	ok  bool
	err error
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
