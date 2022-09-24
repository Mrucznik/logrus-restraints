package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

var std = logrus.StandardLogger()
var emptyLogger = &logrus.Logger{Out: ioutil.Discard}

func main() {
	for i := 0; i < 100; i++ {
		callTTL() // should be called 5 times
		time.Sleep(10 * time.Millisecond)
	}

	for i := 0; i < 100; i++ {
		callEveryN() //should be called 5 times
	}
}

func callTTL() {
	WithTTL(200 * time.Millisecond).Infoln("some ttl call")
}

func callEveryN() {
	EveryN(20).Infoln("some every n call")
}
