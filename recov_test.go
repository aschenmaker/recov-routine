package recovroutine

import (
	"strings"
	"testing"
	"time"
)

func TestRunWithCallbackFunc(t *testing.T) {
	var reRo RecovRoutine

	reRo.Worker = func() {
		var slice []int
		slice[2] = 1 // Panic happens
	}

	reRo.PCallback = func(err error) {
		t.Log("retry times", reRo.RetryCnt)
		if err == nil {
			t.Fatal("unexpected error")
		}

		s := err.Error()
		t.Log(s)
		if s == "" || !strings.Contains(s, "goroutine") {
			t.Fatal("unexpected error")
		}
	}

	reRo.RetryCnt = 3

	reRo.Run()
	time.Sleep(1 * time.Second)
}
