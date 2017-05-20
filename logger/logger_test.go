package logger

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	testObj := NewLogger()

	assert.NotNil(t, testObj, "Unexpected error.")
	assert.Equal(t, reflect.TypeOf(testObj).String(), "*logrus.Logger", "Invalid object returned.")
}
