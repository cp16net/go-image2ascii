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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cp16net/go-image2ascii/image"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	jsonOutput  bool
	out         io.Writer = os.Stdout
	imageWidth  int
	imageHeight int
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert [full filepath to image]",
	Short: "Converts an image to acsii art",
	Long: `Converts an image to acsii art. 
	
	Supports GIF, PNG, JPG formats for images.`,
	Run: convert,
}

func init() {
	RootCmd.AddCommand(convertCmd)
	convertCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "output in json format")
	convertCmd.Flags().IntVarP(&imageWidth, "width", "w", 80, "width of the output image")
	convertCmd.Flags().IntVarP(&imageHeight, "length", "l", 85, "height/length of the output image")
	viper.BindPFlag("width", convertCmd.Flags().Lookup("width"))
	viper.BindPFlag("length", convertCmd.Flags().Lookup("length"))
}

func convert(cmd *cobra.Command, args []string) {

	// check number of args (does cobra do this?)
	if len(args) != 1 {
		printer("**ERROR** Wrong number of arguments")
		return
	}
	filepath := args[0]

	// check for image file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		printer("**ERROR** File does not exist [", filepath, "]")
		return
	}

	// load file from path
	f, err := os.Open(filepath)
	if err != nil {
		printer("**ERROR** file failed to load ", err)
		return
	}

	// call the image convert function here.
	i := image.Image{}
	img, err := i.Execute(f, viper.GetInt("width"), viper.GetInt("length"))
	if err != nil {
		printer("**ERROR** converting image", err)
		return
	}

	if jsonOutput == true {
		j, _ := json.MarshalIndent(img, "", "  ")
		printer(string(j))
	} else {
		printer(img.Data)
	}
}

func printer(a ...interface{}) {
	fmt.Fprint(out, a...)
}
