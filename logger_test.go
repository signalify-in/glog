package glog_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/signalify-in/glog"
	"github.com/stretchr/testify/assert"
)

var (
	logger *glog.Logger
	buf    bytes.Buffer
)

func TestLogger_New(t *testing.T) {
	logger = glog.New(glog.LevelInfo)
	assert.Equal(t, reflect.TypeOf(logger), reflect.TypeOf(&glog.Logger{}))
}

func TestLogger_NewBot(t *testing.T) {}

func TestLogger_Trace(t *testing.T) {
	readOutput()
	logger = glog.New(glog.LevelTrace, "Test")
	assert.Equal(t, 7, len(logger.DirLevels))

	logger.Trace(fmt.Sprintf("trace %v", time.Now()))
	t.Log(buf.String())
}

func TestLogger_Debug(t *testing.T) {
	readOutput()
	logger = glog.New(glog.LevelDebug)
	assert.Equal(t, 6, len(logger.DirLevels))
	logger.Debug("debug")
	logger = glog.New(glog.LevelDebug, "Test Debug")
	logger.Debug("debug test with prefix")
	//add more assertions here
	t.Log(buf.String())
}

func TestLogger_Info(t *testing.T) {
	readOutput()
	logger.Lot = true
	logger.Info("info")
	t.Log(buf.String())
}

func TestLogger_Warning(t *testing.T) {
	readOutput()
	logger.Warning("warn")
	t.Log(buf.String())
}

func TestLogger_Error(t *testing.T) {
	readOutput()
	logger.Error("error")
	t.Log(buf.String())
}

func TestLogger_Fatal(t *testing.T) {}

func TestLogger_Panic(t *testing.T) {}

func readOutput() {
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
}
