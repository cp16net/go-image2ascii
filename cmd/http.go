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
	"github.com/cp16net/go-image2ascii/server"
	"github.com/spf13/cobra"
)

var (
	serverPort      int
	serverInterface string
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run as a web service",
	Long:  `This will allow you to run a web service that can convert images to ascii art.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("http called")
		conf := server.Config{}
		conf.Iface = serverInterface
		conf.Port = serverPort
		server.Run(&conf)
	},
}

func init() {
	RootCmd.AddCommand(httpCmd)
	httpCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "port on which the server will listen")
	httpCmd.Flags().StringVarP(&serverInterface, "bind", "", "127.0.0.1", "interface to which the server will bind")
}
