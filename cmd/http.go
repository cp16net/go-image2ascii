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
	"log"

	"github.com/cp16net/go-image2ascii/gen/models"
	"github.com/cp16net/go-image2ascii/gen/restapi"
	"github.com/cp16net/go-image2ascii/gen/restapi/operations"
	"github.com/cp16net/go-image2ascii/image"
	"github.com/cp16net/go-image2ascii/logger"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	httpPort        int
	httpAddressBind string
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run an http server",
	Long:  `This runs an http web server. (testing this out)`,
	Run:   serveHTTP,
}

func init() {
	RootCmd.AddCommand(httpCmd)
	httpCmd.Flags().IntVarP(&httpPort, "port", "p", 8080, "port on which the server will listen")
	httpCmd.Flags().StringVarP(&httpAddressBind, "bind", "", "127.0.0.1", "interface to which the server will bind")
	viper.BindPFlag("port", httpCmd.Flags().Lookup("port"))
	viper.BindPFlag("bind", httpCmd.Flags().Lookup("bind"))
}

func serveHTTP(cmd *cobra.Command, args []string) {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewConverterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = viper.GetInt("port")
	server.Host = viper.GetString("bind")

	// api.GetHandler = operations.GetHandlerFunc(index)
	api.GetHealthzCheckHandler = operations.GetHealthzCheckHandlerFunc(health)
	api.PostUploadHandler = operations.PostUploadHandlerFunc(uploaderPost)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

// NOTE: without extra work the html file can not be served because
// go-swagger does not have an HTMLProducer from the
// go-openapi/runtime library. Would need to create one that produced
// the mimetype of text/html and not sure how to do that right now.

// func index(params operations.GetParams) middleware.Responder {
// 	// greeting := fmt.Sprintf("Hello, world!")
// 	buf := bytes.NewBuffer(nil)
// 	f, err := os.Open("./templates/index.html")
// 	if err != nil {
// 		e := models.Htmlerror(err.Error())
// 		return operations.NewGetUploadDefault(500).WithPayload(e)
// 	}
// 	defer f.Close()
// 	io.Copy(buf, f)

// 	return operations.NewGetOK().WithPayload(string(buf.Bytes()))
// }

func health(params operations.GetHealthzCheckParams) middleware.Responder {
	running := true
	return operations.NewGetHealthzCheckOK().WithPayload(&models.Health{Running: &running})
}

func uploaderPost(params operations.PostUploadParams) middleware.Responder {
	img, err := image.Execute(params.Uploadfile.Data)
	if err != nil {
		logger.Logger.Error("Error from image convert: ", err.Error())
		es := "failed to upload file"
		e := models.Error{Message: &es}
		return operations.NewPostUploadDefault(500).WithPayload(&e)
	}
	return operations.NewPostUploadOK().WithPayload(&models.ASCII{String: &img.Data})

}
