// Copyright © 2017 Craig Vyvial <cp16net@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"io"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

const (
	// DEBUG level messages will be outputed to log
	DEBUG = "debug"

	// INFO level messages will be outputed to log
	INFO = "info"

	// WARN level messages and above output to log
	WARN = "warn"

	// ERROR level messages will be outputed to log
	ERROR = "error"

	// FATAL level messages will be outputed to log
	FATAL = "fatal"
)

// NewLogger creates a new logger
func NewLogger(level string, out io.Writer) *logrus.Logger {

	minLogLevel := logrus.DebugLevel
	switch strings.ToLower(level) {
	case INFO:
		minLogLevel = logrus.InfoLevel
	case WARN:
		minLogLevel = logrus.WarnLevel
	case ERROR:
		minLogLevel = logrus.ErrorLevel
	case FATAL:
		minLogLevel = logrus.FatalLevel
	case DEBUG:
		minLogLevel = logrus.DebugLevel
	}

	logger := logrus.New()
	logger.Out = out
	logger.Level = minLogLevel

	return logger
}

// Logger is a default logger for application use
var Logger = NewLogger("debug", os.Stdout)
