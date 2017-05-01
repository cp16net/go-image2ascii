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

package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertJsonOutput(t *testing.T) {
	jsonOutput = true
	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{"../gopher.png"})
	assert.Contains(t, buf.String(), "{\n  \"string\": \"")
}

func TestConvertOutput(t *testing.T) {
	jsonOutput = false
	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{"../gopher.png"})
	assert.NotContains(t, buf.String(), "\"string\": \"")
}

func TestConvertWrongArgs(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{"test", "test"})
	t.Log(buf.String())
	assert.Contains(t, buf.String(), "Wrong number of arguments")
}

func TestConvertFileDNE(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{"./file-does-not-exist"})
	assert.Contains(t, buf.String(), "File does not exist")
}

func TestConvertNonImageFile(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{"../README.org"})
	assert.Contains(t, buf.String(), "converting image")
}

func TestConvertBadFileProps(t *testing.T) {
	f, _ := ioutil.TempFile("", "test")
	f.WriteString("test")
	f.Chmod(0000)
	defer f.Close()
	defer os.Remove(f.Name())

	buf := &bytes.Buffer{}
	out = buf
	convert(convertCmd, []string{f.Name()})
	assert.Contains(t, buf.String(), "file failed to load")
}
