package recovroutine

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

// PCallback defines the callback func after pannic
// occurs.
type PCallback func(err error)

// WorkerFunc defines what the goroutines worker func
// to execute.
type WorkerFunc func()

// RecovRoutine defines the recover.
type RecovRoutine struct {
	error
	// RetryCnts defines the times to retry the workerfunc
	//
	// Optional.
	RetryCnt int
	// Worker defines the times to retry the workerfunc
	//
	// Required.
	Worker WorkerFunc
	// PCallback defines the times to retry the workerfunc
	//
	// Optional. Default: nil
	PCallback PCallback
}

func (r *RecovRoutine) Recover() {
	logger := log.New(os.Stderr, "\n\n\x1b[31m", log.LstdFlags)
	if e := recover(); e != nil {
		r.error = errors.New(fmt.Sprintf("%s, %s", r, string(debug.Stack())))
		if logger != nil {
			logger.Printf("[Recov Recovery] %s panic recovered:\n%s\n%s",
				timeFormat(time.Now()), r.error, debug.Stack())
		}

		if r.PCallback != nil {
			r.PCallback(r)
		}

		if r.RetryCnt > 0 {
			logger.Printf("[Recov retry]retry %v times:\n", r.RetryCnt)
			r.RetryCnt--
			go r.Run()
		}
	}
}

func (r *RecovRoutine) Error() string {
	if r.error != nil {
		return r.error.Error()
	}
	return ""
}

func (r *RecovRoutine) Run() {
	defer r.Recover()
	r.Worker()
}

func timeFormat(t time.Time) string {
	var timeString = t.Format("2006/01/02 - 15:04:05")
	return timeString
}
