package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"learn/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Serivce
}

func NewHandler(service *service.Serivce) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world")

	})
	router.HandleFunc("/send", h.SendMail)

	return router
}
