package user

import (
	"encoding/json"
	"learning-project/internal/handlers"
	"learning-project/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	usersUrl      = "/users"
	userUrl       = "/users/:uuid"
	createUserUrl = "/user/create"
)

type handler struct {
	logger  *logging.Logger
	service Service
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, h.GetList)
	router.GET(userUrl, h.GetUser)
	router.POST(createUserUrl, h.CreateUser)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartialyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
}

func NewHandler(logger *logging.Logger, service Service) handlers.Handler {
	return &handler{
		logger:  logger,
		service: service,
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
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	created, err := h.service.Register(r.Context(), req)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
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
