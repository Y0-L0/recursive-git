package testutils

import (
	"reflect"
	"testing"
)

// assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf(msg, v...)
	}
}

// ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("unexpected error: %s", err.Error())
	}
}

// equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	expType := reflect.TypeOf(exp).String()
	actType := reflect.TypeOf(act).String()

	if expType != actType {
		tb.Fatalf("\nexp: %#v\ngot: %#v", expType, actType)
		return
	}

	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\nexp: %#v\ngot: %#v", exp, act)
	}
}
