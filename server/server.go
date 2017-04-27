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
