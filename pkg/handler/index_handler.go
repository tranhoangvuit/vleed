package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexHandler handle index page
func (h *Handler) IndexHandler() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		h.Server.Rndr.HTML(w, http.StatusOK, "index", nil)
	}
}

// NotFoundHandler handle not found page
// func (h *Handler) NotFoundHandler() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("ABC")
// 	}
// }
