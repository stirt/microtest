package assert

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type comparatorFunc func() bool

func assert[T, S any](t *testing.T, want T, have S, f comparatorFunc) {
	t.Helper()
	if !f() {
		t.Errorf(`want: "%v", have: "%v"`+"\n"+"%s", want, have, cmp.Diff(want, have))
	}
}

func Assert[T comparable](t *testing.T, want T, have T) {
	assert(t, want, have, func() bool {
		return want == have
	})
}

func AssertObject(t *testing.T, want any, have any) {
	assert(t, want, have, func() bool {
		return cmp.Equal(want, have)
	})
}

func AssertNil(t *testing.T, have any) {
	t.Helper()
	if have != nil {
		t.Errorf(`want: "%v", have: "%v"`+"\n"+"%s", nil, have, cmp.Diff(nil, have))
	}
}

func AssertNotNil(t *testing.T, have any) {
	assert(t, "value shouldn't be nil", have, func() bool {
		return have != nil
	})
}

func AssertErrorIs(t *testing.T, want error, have error) {
	assert(t, want, have, func() bool {
		return errors.Is(have, want)
	})
}

func AssertErrorAs(t *testing.T, want any, have error) {
	assert(t, want, have, func() bool {
		return errors.As(have, want)
	})
}
