package glog_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

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

func TestLogger_NewWithLevels(t *testing.T) {
	logger.LogPrefix = "Test"
	logger = glog.New(glog.LevelInfo)
	assert.Equal(t, 5, len(logger.DirLevels))

	logger = glog.New(glog.LevelFatal)
	assert.Equal(t, 2, len(logger.DirLevels))

	logger = glog.New(glog.LevelDebug)
	assert.Equal(t, 6, len(logger.DirLevels))
}

func TestLogger_NewBot(t *testing.T) {}

func TestLogger_Trace(t *testing.T) {
	readOutput()
	logger.Trace("trace")
	logger.Info(fmt.Sprintf("tracing log %v \n", len(logger.DirLevels)))
	t.Log(buf.String())
}

func TestLogger_Debug(t *testing.T) {
	readOutput()
	logger.Debug("debug")
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
