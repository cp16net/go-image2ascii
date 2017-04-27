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
	"fmt"
	"net/http"
	"strconv"

	"github.com/cp16net/go-image2ascii/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
	"github.com/tylerb/graceful"
)

var (
	serverPort      int
	serverInterface string
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		// TODO separate this into server pacakage...

		// mux handler
		router := httprouter.New()
		router.POST("/convert", convert)
		logger.Logger.Info("Routes setup starting server")
		// Serve this program forever
		port := strconv.Itoa(serverPort)
		host := serverInterface
		httpServer := &graceful.Server{Server: new(http.Server)}
		httpServer.Addr = host + ":" + port
		httpServer.Handler = router
		logger.Logger.Infof("listening at http://%s:%s", host, port)
		if err := httpServer.ListenAndServe(); err != nil {
			shutdown(err)
		}
	},
}

func convert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// data, err := mysql.Write()
	// if err != nil {
	// 	logger.Logger.Error(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// output, err := json.MarshalIndent(data, "", "\t")
	// if err != nil {
	// 	logger.Logger.Error(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	logger.Logger.Debug("call to /convert")
	logger.Logger.Debug("repsonse writer: ", w)
	logger.Logger.Debug("request: ", r)
	logger.Logger.Debug("params: ", ps)
	w.Header().Set("content-type", "application/json")
	fmt.Fprintln(w, string(`{"message": "this is a test"}`))
}

// shutdown closes down the api server
func shutdown(err error) {
	logger.Logger.Fatalln(err)
}

func init() {
	RootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	httpCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "port on which the server will listen")
	httpCmd.Flags().StringVarP(&serverInterface, "bind", "", "127.0.0.1", "interface to which the server will bind")

	// 	Host string `env:"HOST" default:"0.0.0.0" long:"host" description:"HTTP listen server"`
	// 	Port int    `env:"PORT" default:"8080" long:"port" description:"HTTP listen port"`

}
