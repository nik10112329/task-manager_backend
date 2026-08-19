package user

import (
	"learning-project/internal/handlers"
	"learning-project/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, h.GetList)
	router.GET(userUrl, h.GetUser)
	router.POST(usersUrl, h.CreateUser)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartialyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is list users"))
	w.WriteHeader(200)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is get users"))
	w.WriteHeader(200)
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create users"))
	w.WriteHeader(201)
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update users"))
	w.WriteHeader(204)
}
func (h *handler) PartialyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is particialy update users"))
	w.WriteHeader(204)
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is delete users"))
	w.WriteHeader(204)
}
