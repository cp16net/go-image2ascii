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
	"context"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/cp16net/go-image2ascii/image"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

// Iris server type
type Iris struct {
	Conf *Config
	Log  *logrus.Logger
}

// Server runs the iris server
func (i Iris) Server() {
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		view.HTML("./templates", ".html"),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}),
		iris.EventPolicy{
			// Interrupt Event means when control+C pressed on terminal.
			Interrupted: func(*iris.Framework) {
				// shut down gracefully, but wait 5 seconds the maximum before closed
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				app.Shutdown(ctx)
			},
		},
	)

	// Serve the index page with html form for upload
	app.Get("/", i.index)

	// Serve the health check to verify service is up
	app.Get("/healthz-check", i.healthz)

	// Serve the form.html to the user
	app.Get("/upload", i.index)

	// Handle the post request from the upload_form.html to the server
	app.Post("/upload", iris.LimitRequestBodySize(10<<20), i.upload)

	// start the server at 127.0.0.1:8080
	port := strconv.Itoa(i.Conf.Port)
	app.Listen(i.Conf.Iface + ":" + port)
}

// index serves the webpage
func (i Iris) index(ctx *iris.Context) {
	i.Log.Debug("call to ", ctx.Path())
	ctx.Render("index.html", nil)
}

// healthz is a simple function to check that server is up and
// accepting requests
func (i Iris) healthz(ctx *iris.Context) {
	i.Log.Debug("call to ", ctx.Path())
	ctx.JSON(iris.StatusOK, health{Running: true})
}

// upload handler
func (i Iris) upload(ctx *iris.Context) {
	i.Log.Debug("call to ", ctx.Path())
	// Get the file from the request
	file, _, err := ctx.FormFile("uploadfile")
	if err != nil {
		i.Log.Error("Error uploading image: ", err.Error())
		ctx.HTML(iris.StatusInternalServerError,
			"Error while uploading: <b>"+err.Error()+"</b>")
		return
	}
	defer file.Close()

	// run conversion
	img, err := image.Execute(file)
	if err != nil {
		i.Log.Error("Error from image convert: ", err.Error())
		ctx.HTML(iris.StatusInternalServerError,
			"Error while uploading: <b>"+err.Error()+"</b>")
		return
	}
	ctx.JSON(iris.StatusOK, img)
}
