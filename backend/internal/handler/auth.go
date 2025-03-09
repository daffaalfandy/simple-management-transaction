package handler

import (
	"backend/internal/models"
	authService "backend/internal/service"
	"backend/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	db          *sql.DB
	authService authService.IAuth
}

func NewHandler(db *sql.DB, authService authService.IAuth) Handler {
	return Handler{
		db:          db,
		authService: authService,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := h.authService.Register(&user)
	if err != nil {
		utils.ParseJSON(w, r, utils.ErrorResponse(err), http.StatusInternalServerError)
		return
	}

	utils.ParseJSON(w, r, result, http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserLoginParams
	json.NewDecoder(r.Body).Decode(&credentials)

	token, err := h.authService.Login(&credentials)
	if err != nil {
		utils.ParseJSON(w, r, utils.ErrorResponse(err), http.StatusUnauthorized)
		return
	}

	utils.ParseJSON(w, r, utils.SuccessResponse(map[string]string{"token": token}), http.StatusOK)
}
