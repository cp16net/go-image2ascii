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
	"os"
	"testing"

	"github.com/cp16net/go-image2ascii/gen/restapi/operations"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	p := operations.GetHealthzCheckParams{}
	o := health(p)
	assert.NotNil(t, o)
}

func TestUploadPostBadInput(t *testing.T) {
	p := operations.PostUploadParams{}
	o := uploaderPost(p)
	assert.NotNil(t, o)
}

func TestUploadPost(t *testing.T) {
	f, _ := os.Open("../gopher.png")
	defer f.Close()
	p := operations.PostUploadParams{}
	p.Uploadfile = runtime.File{Data: f}
	o := uploaderPost(p)
	assert.NotNil(t, o)
}
