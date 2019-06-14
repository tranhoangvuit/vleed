package server

import (
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

// Server is the main application
type Server struct {
	Rndr   *render.Render
	Router *httprouter.Router
	Store  *sessions.CookieStore
	Bm     *bluemonday.Policy
}

// LoadConfiguration will load config file
func LoadConfiguration(pwd, fileName string) error {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(filepath.Dir(pwd))
	return viper.ReadInConfig()
}

// SetupServer will set up the basic information
func SetupServer(r *httprouter.Router, cookieSecretKey []byte, templateFolderPath string) *Server {
	rndr := render.New(render.Options{
		Directory:  templateFolderPath,
		Extensions: []string{".html"},
		Layout:     "base",
	})

	bm := bluemonday.UGCPolicy()

	return &Server{
		Rndr:   rndr,
		Router: r,
		Store:  sessions.NewCookieStore(cookieSecretKey),
		Bm:     bm,
	}
}
