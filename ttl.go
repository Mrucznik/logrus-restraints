package logrus_restraints

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
	"time"
)

var ttlCallers sync.Map

func WithTTL(duration time.Duration) *logrus.Logger {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("cannot find caller")
	}

	if _, alreadyCalled := ttlCallers.Load(callKey(file, line)); alreadyCalled {
		return emptyLogger
	}

	ttlCallers.Store(callKey(file, line), struct{}{})
	go time.AfterFunc(duration, func() {
		ttlCallers.Delete(fmt.Sprintf("%s%d", file, line))
	})

	return std
}
