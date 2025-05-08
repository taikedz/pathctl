package putest

import (
	"testing"
)

func CheckEqual[V string|bool](t *testing.T, exp_value V, got_value V) {
	if exp_value != got_value {
		t.Errorf("Got %#v // Exp %#v", got_value, exp_value)
	}
}

func CheckEqualArr[V string|bool](t *testing.T, exp_value []V, got_value []V) {
	if len(exp_value) != len(got_value) {
		t.Errorf("Got %#v // Exp %#v", got_value, exp_value)
		return
	}

	for i:=0; i<len(exp_value) && i<len(got_value); i++ {
		if exp_value[i] != got_value[i] {
			t.Errorf("Got %#v // Exp %#v", got_value[i], exp_value[i])
			return
		}
	}
}
