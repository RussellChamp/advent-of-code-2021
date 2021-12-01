package timer

import "testing"

func TestTimer(t *testing.T) {

	Start()
	Tick()

	t.Log("Yup")
}
