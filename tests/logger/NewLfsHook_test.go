package logger_test

import (
	"testing"

	"coursesheduling/lib/log"
)

func TestFL(t *testing.T) {
	log.InitLog(0)
	log.Logger.Debugln("debug")
	log.Logger.Errorln("error")
}
