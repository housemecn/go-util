package flag

import "testing"

func TestIsTesting(t *testing.T) {
	if !IsTesting() {
		t.Error("IsTesting failed")
	}
}
