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

package image

import (
	"bytes"
	"image"
	"image/color"
	"os"
	"reflect"

	"github.com/nfnt/resize"

	// load these for image decoding
	// add more here for additional image types
	_ "image/jpeg"
	_ "image/png"
)

const (
	width = 100

	// ASCIISTR is the darkness level for characters
	ASCIISTR = "MND8OZ$7I?+=~:,.."
)

// Image data describing the image
type Image struct {
	Data string `json:"string"`
}

// Execute image conversion to ascii represenation
func Execute(filepath string) (*Image, error) {
	// load file from path
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	// logger.Logger.Debug("image decoded with format", format)
	f.Close()

	// set output image size
	sz := img.Bounds()
	h := (sz.Max.Y * width * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(width), uint(h), img, resize.Lanczos3)

	// create ascii represenation of light to dark contrasts
	// break image up into even blocks
	// iterate over each block in the image
	// get darkness by adding up the color values
	// write the ascii that represents the range of darkness in location

	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < width; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}

	return &Image{Data: string(buf.Bytes())}, nil
	// return &Image{Data: []byte("****")}, nil
}
