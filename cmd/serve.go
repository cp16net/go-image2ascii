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

func init() {
	RootCmd.AddCommand(serverCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		server()
	},
}

// The server itself
func server() {
	logger.Logger.Info("Starting up web application")

	// mux handler
	router := httprouter.New()
	router.POST("/convert", convert)

	logger.Logger.Info("Routes setup starting server")

	// Serve this program forever
	port := strconv.Itoa(AppConfig.Port)
	host := AppConfig.Host
	httpServer := &graceful.Server{Server: new(http.Server)}
	httpServer.Addr = host + ":" + port
	httpServer.Handler = router
	logger.Logger.Infof("listening at http://%s:%s", host, port)
	if err := httpServer.ListenAndServe(); err != nil {
		shutdown(err)
	}
}

// shutdown closes down the api server
func shutdown(err error) {
	logger.Logger.Fatalln(err)
}
