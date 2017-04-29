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
	"io"

	"github.com/cp16net/go-image2ascii/logger"
	"github.com/nfnt/resize"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG, GIF, and PNG formatted images.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	width = 100

	// ASCIISTR is the 16 darkness levels of characters
	ASCIISTR = "@ND8OZ$7I?+=~:,.."
)

// Image data describing the image
type Image struct {
	Data string `json:"string"`
}

// Execute image conversion to ascii represenation
func Execute(f io.Reader) (*Image, error) {

	img, _, err := image.Decode(f)
	if err != nil {
		logger.Logger.Error("could not decode the image")
		return nil, err
	}
	return convert(img)
}

// convert image here
//
// Steps
//
// 1. parse the size of the image.
// 2. resize the image to smaller size by set width.
// 3. iterate of the pixels of the image and get the greyscale value
// 4. convert the greyscale value to ASCII mapping of 16 colors
// 5. write the value to the new ascii buffer and continue 4-6 until end of image.
// 6. return Image object of Data as an ASCII string
func convert(img image.Image) (*Image, error) {
	// set output image size
	sz := img.Bounds()
	h := (sz.Max.Y * width * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(width), uint(h), img, resize.Lanczos3)

	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < width; j++ {
			p := img.At(j, i)
			g := color.GrayModel.Convert(p)
			y, _, _, _ := g.RGBA()
			pos := int(y * 16 / 1 >> 16)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}

	return &Image{Data: string(buf.Bytes())}, nil
}
