package main

import (
	"time"

	rv "github.com/aschenmaker/recov-routine"
)

func main() {

	r := rv.RecovRoutine{
		RetryCnt: 1,
		PCallback: func(err error) {
			if err != nil {
				println(err.Error())
			}
		},
	}
	r.Worker = func() {
		var slice []int
		slice[2] = 1
	}
	r.Run()
	// Waiting for retry.
	time.Sleep(10 * time.Second)
}
