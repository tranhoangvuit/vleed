package handler

import "github.com/tranhoangvuit/vleed/pkg/server"

// Handler wrap common infomation to send to view
type Handler struct {
	*server.Server
}

// NewHander return new handler
func NewHander(app *server.Server) Handler {
	return Handler{
		Server: app,
	}
}
