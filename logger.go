package logrus_restraints

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var std = logrus.StandardLogger()
var emptyLogger = &logrus.Logger{Out: ioutil.Discard}

func callKey(file string, line int) string {
	return fmt.Sprintf("%s%d", file, line)
}
