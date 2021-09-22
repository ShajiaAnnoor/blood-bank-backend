package server

import (
	"github.com/sirupsen/logrus"
)

//Logrus returns logrus related configurations
func Logrus() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}
