// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/cp16net/go-image2ascii/image"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
)

const (
	outputFile = "local-file-output.png"
)

var (
	input        string
	fontLocation string
)

// bannerCmd represents the banner command
var bannerCmd = &cobra.Command{
	Use:   "banner [string to print]",
	Short: "Make a banner with some text for dot matrix printer",
	Long:  `Generates a banner with text given for a dot matrix printer.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}
		input := strings.Join(args[:], " ")

		//check os
		if runtime.GOOS == "windows" {
			fmt.Println("sorry this wont work on windows [dont know where the fonts live]")
			return
		} else if runtime.GOOS == "linux" {
			fontLocation = "/usr/share/fonts/truetype/freefont/FreeMonoBold.ttf"
		} else if runtime.GOOS == "darwin" {
			fontLocation = "/Library/Fonts/Courier New Bold.ttf"
		}

		const S = 1024
		dc := gg.NewContext(S, S)
		w, h := dc.MeasureString(input)
		h = 80
		w = w * 7.5

		dc = gg.NewContext(int(w), int(h))
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		if err := dc.LoadFontFace(fontLocation, 80); err != nil {
			panic(err)
		}
		dc.DrawStringAnchored(input, w/2, h/2, 0.5, 0.35)

		dc.SavePNG(outputFile)
		src, err := imaging.Open(outputFile)
		if err != nil {
			log.Fatalf("Open failed: %v", err)
		}
		dst := imaging.Rotate270(src)
		imaging.Save(dst, outputFile)

		// load file from path
		f, err := os.Open(outputFile)
		if err != nil {
			printer("**ERROR** file failed to load ", err)
			return
		}
		defer f.Close()
		defer os.Remove(outputFile)

		i := image.Image{}
		img, err := i.Execute(f, -1, -1)
		if err != nil {
			printer("**ERROR** converting image", err)
			return
		}

		printer(img.Data)
	},
}

func init() {
	RootCmd.AddCommand(bannerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bannerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bannerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
