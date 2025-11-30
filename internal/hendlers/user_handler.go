package handler

import (
	"encoding/json"
	"gocrud/internal/domain"
	"gocrud/internal/stor"
	"io"
	"net/http"
)

type UserHandler struct {
	s *stor.Storage
}

func NewUserHandler(newStor *stor.Storage) *UserHandler {
	return &UserHandler{
		s: newStor,
	}
}

func (u *UserHandler) CreatUser(w http.ResponseWriter, r *http.Request) {

	user := domain.User{}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(resp, user)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	u.s.CreateUser(&user)
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(resp, user)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	u.s.GetUser(user.Id)
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(resp, user)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

}
