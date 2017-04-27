package cmd

// import (
// 	// "encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/jessevdk/go-flags"
// 	"github.com/julienschmidt/httprouter"
// 	"github.com/tylerb/graceful"

// 	"github.com/cp16net/go-image2ascii/logger"
// )

// // Config for webappliation
// type Config struct {
// 	Host string `env:"HOST" default:"0.0.0.0" long:"host" description:"HTTP listen server"`
// 	Port int    `env:"PORT" default:"8080" long:"port" description:"HTTP listen port"`
// }

// var (
// 	// AppConfig configuration for web application
// 	AppConfig Config
// 	parser    = flags.NewParser(&AppConfig, flags.Default)
// )

// // Parse all of the bindata templates
// func init() {
// 	// parse the flags
// 	_, err := parser.Parse()
// 	if e, ok := err.(*flags.Error); ok {
// 		if e.Type == flags.ErrHelp {
// 			os.Exit(0) //exit without error in case of help
// 		} else {
// 			os.Exit(1) //exit with error for other cases
// 		}
// 	}

// }

// func convert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// data, err := mysql.Write()
// 	// if err != nil {
// 	// 	logger.Logger.Error(err)
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// output, err := json.MarshalIndent(data, "", "\t")
// 	// if err != nil {
// 	// 	logger.Logger.Error(err)
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	w.Header().Set("content-type", "application/json")
// 	fmt.Fprintln(w, string(`{"message": "this is a test"}`))
// }

// // The server itself
// func main() {
// 	logger.Logger.Info("Starting up web application")

// 	// mux handler
// 	router := httprouter.New()
// 	router.POST("/convert", convert)

// 	logger.Logger.Info("Routes setup starting server")

// 	// Serve this program forever
// 	port := strconv.Itoa(AppConfig.Port)
// 	host := AppConfig.Host
// 	httpServer := &graceful.Server{Server: new(http.Server)}
// 	httpServer.Addr = host + ":" + port
// 	httpServer.Handler = router
// 	logger.Logger.Infof("listening at http://%s:%s", host, port)
// 	if err := httpServer.ListenAndServe(); err != nil {
// 		shutdown(err)
// 	}
// }

// // shutdown closes down the api server
// func shutdown(err error) {
// 	logger.Logger.Fatalln(err)
// }
