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
	"github.com/cp16net/go-image2ascii/logger"
	"github.com/cp16net/go-image2ascii/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	irisPort        int
	irisAddressBind string
)

// irisCmd represents the iris command
var irisCmd = &cobra.Command{
	Use:   "iris",
	Short: "Run an iris server",
	Long:  `This runs an iris web server. (testing this out)`,
	Run:   serve,
}

func init() {
	RootCmd.AddCommand(irisCmd)
	irisCmd.Flags().IntVarP(&irisPort, "port", "p", 8080, "port on which the server will listen")
	irisCmd.Flags().StringVarP(&irisAddressBind, "bind", "", "127.0.0.1", "interface to which the server will bind")
	viper.BindPFlag("port", irisCmd.Flags().Lookup("port"))
	viper.BindPFlag("bind", irisCmd.Flags().Lookup("bind"))
}

func serve(cmd *cobra.Command, args []string) {
	i := server.Iris{}
	i.Conf = &server.Config{}
	i.Conf.Iface = viper.GetString("bind")
	i.Conf.Port = viper.GetInt("port")
	i.Log = logger.Logger
	i.Server()
}
