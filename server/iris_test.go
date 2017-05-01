package server

import (
	"os"
	"testing"

	"github.com/cp16net/go-image2ascii/logger"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/httptest"
)

func TestIrisServer(t *testing.T) {
	i := Iris{Log: logger.Logger}
	app := i.Server()
	app.Adapt(view.HTML("../templates", ".html"))
	e := httptest.New(app, t)
	h := e.GET("/healthz-check").Expect().Status(iris.StatusOK).JSON().Object()
	h.Equal(map[string]bool{
		"running": true,
	})
	e.GET("/").Expect().Status(iris.StatusOK)
	e.GET("/upload").Expect().Status(iris.StatusOK)
	f, _ := os.Open("../gopher.png")
	e.POST("/upload").WithMultipart().WithFile("uploadfile", "../gopher.png", f).Expect().Status(iris.StatusOK).JSON().Object().ContainsKey("string")
	e.POST("/upload").WithMultipart().WithFile("uploadfi", "../gopher.png", f).Expect().Status(iris.StatusInternalServerError)
	f2, _ := os.Open("../README.org")
	e.POST("/upload").WithMultipart().WithFile("uploadfile", "../README.org", f2).Expect().Status(iris.StatusInternalServerError)
}
