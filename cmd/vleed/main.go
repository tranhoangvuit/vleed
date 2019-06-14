package main

import (
	"flag"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/tranhoangvuit/vleed/pkg/handler"
	"github.com/tranhoangvuit/vleed/pkg/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "path to config file for this app")
}

func main() {
	flag.Parse()
	// Load configuration
	err := server.LoadConfiguration((configPath + "/configs/"), "vfeed_config.json")
	if err != nil {
		panic(err)
	}
	staticFilePath := path.Join(viper.GetString("path"), "static")
	templateFolderPath := path.Join(viper.GetString("path"), "templates")

	router := httprouter.New()
	cookieSecretKey := viper.GetString("cookieSecret")
	server := server.SetupServer(router, []byte(cookieSecretKey), templateFolderPath)

	handler := handler.NewHander(server)

	router.GET("/", handler.IndexHandler())

	router.ServeFiles("/static/*filepath", http.Dir(staticFilePath))
	// router.NotFound = http.HandlerFunc(handler.NotFoundHandler())

	http.ListenAndServe(":8090", router)
}
