package assert

import "testing"

func IsTrue(t *testing.T, expr bool, message string) {
	if expr != true {
		t.Logf("Expected true but got false: %s", message)
		t.Fail()
	}
}

func StrEquals(t *testing.T, actual, expected, message string) {
	if actual != expected {
		t.Logf("Expected value '%s' but got '%s': %s", expected, actual, message)
		t.Fail()
	}
}
