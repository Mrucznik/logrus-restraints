package main

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

	if _, alreadyCalled := ttlCallers.Load(ttlKey(file, line)); alreadyCalled {
		return emptyLogger
	}

	ttlCallers.Store(ttlKey(file, line), struct{}{})
	go time.AfterFunc(duration, func() {
		ttlCallers.Delete(fmt.Sprintf("%s%d", file, line))
	})

	return std.Logger
}

func ttlKey(file string, line int) string {
	return fmt.Sprintf("%s%d", file, line)
}
