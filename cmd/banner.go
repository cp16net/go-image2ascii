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
	"strings"

	"github.com/cp16net/go-image2ascii/image"
	"github.com/spf13/cobra"
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

		b := image.Banner{Input: input}
		img, err := b.Execute()
		if err != nil {
			printer("**ERROR** failed to convert banner", err)
			return
		}
		printer(img.Data)
	},
}

func init() {
	RootCmd.AddCommand(bannerCmd)
}
