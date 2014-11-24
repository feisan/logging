package logging

import (
	"testing"
)

func Test_LogDebug(t *testing.T) {
	Debug("Hello World")
	Debugf("%s %d%s", "Hello", 2014, "!")
	t.Log("Test_LogDebug OK")
}

func Test_LogInfo(t *testing.T) {
	Info("Hello World")
	Infof("%s %d%s", "Hello", 2014, "!")
	t.Log("Test_LogInfo OK")
}

func Test_LogWarning(t *testing.T) {
	Warning("Hello World")
	Warningf("%s %d%s", "Hello", 2014, "!")
	t.Log("Test_LogWarning OK")
}

func Test_LogError(t *testing.T) {
	Error("Hello World")
	Errorf("%s %d%s", "Hello", 2014, "!")
	t.Log("Test_LogError OK")
}

func Test_LogCritical(t *testing.T) {
	Critical("Hello World")
	Criticalf("%s %d%s", "Hello", 2014, "!")
	t.Log("Test_LogCritical OK")
}
