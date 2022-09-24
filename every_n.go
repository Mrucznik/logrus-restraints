package logrus_restraints

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

var everyNCallers sync.Map

func EveryN(n uint) *logrus.Logger {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("cannot find caller")
	}

	key := callKey(file, line)
	timesRaw, alreadyCalled := everyNCallers.Load(key)

	if alreadyCalled {
		times, ok := timesRaw.(uint)
		if !ok {
			panic("cannot cast timesRaw")
		}

		if times < n {
			everyNCallers.Store(key, times+1)
			return emptyLogger
		}
	}

	everyNCallers.Store(key, uint(1))
	return std
}
