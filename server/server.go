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

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cp16net/go-image2ascii/image"
	"github.com/cp16net/go-image2ascii/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
)

// Config struct
type Config struct {
	Iface string
	Port  int
}

// Run starts the server
func Run(conf *Config) {
	// mux handler
	router := httprouter.New()
	router.POST("/convert", convert)
	logger.Logger.Info("Routes setup starting server")
	// Serve this program forever
	port := strconv.Itoa(conf.Port)
	host := conf.Iface
	httpServer := &graceful.Server{Server: new(http.Server)}
	httpServer.Addr = host + ":" + port
	httpServer.Handler = router
	logger.Logger.Infof("listening at http://%s:%s", host, port)
	if err := httpServer.ListenAndServe(); err != nil {
		shutdown(err)
	}

}

func convert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger.Logger.Debug("call to /convert")
	logger.Logger.Debug("repsonse writer: ", w)
	logger.Logger.Debug("request: ", r)
	logger.Logger.Debug("params: ", ps)

	filepath := ""
	img, err := image.Execute(filepath)
	if err != nil {
		logger.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json")
	fmt.Fprintln(w, string(`{"message": "this is a test"}`))
	j, _ := json.MarshalIndent(img, "", "  ")
	fmt.Fprintln(w, string(j))
}

// shutdown closes down the api server
func shutdown(err error) {
	logger.Logger.Fatalln(err)
}
