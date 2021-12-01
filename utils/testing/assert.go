package assert

import "testing"

func StrEquals(t *testing.T, actual, expected, message string) {
	if actual != expected {
		t.Logf("Expected value '%s' but got '%s': %s", expected, actual, message)
		t.Fail()
	}
}
