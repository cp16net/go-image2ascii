package logger

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoNewLogger(t *testing.T) {
	testObj := NewLogger("info", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}

func TestWarnNewLogger(t *testing.T) {
	testObj := NewLogger("warn", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}

func TestDebugNewLogger(t *testing.T) {
	testObj := NewLogger("debug", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}

func TestErrorNewLogger(t *testing.T) {
	testObj := NewLogger("error", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}

func TestFatalNewLogger(t *testing.T) {
	testObj := NewLogger("fatal", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}

func TestDefaultNewLogger(t *testing.T) {
	testObj := NewLogger("anything", os.Stdout)

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}
