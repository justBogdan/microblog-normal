package handlers

import (
	json2 "encoding/json"
	"io"
	"microblog-normal/internal/service"
	"net/http"
	"unicode/utf8"
)

type RegisterHandler struct {
	registrar service.Registrar
}

func NewRegisterHandler(registrar service.Registrar) *RegisterHandler {
	return &RegisterHandler{registrar: registrar}
}

func (handler *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var input struct {
		Nickname string `json:"nickname"`
	}
	err = json2.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if input.Nickname == "" {
		http.Error(w, "nick name can not be empty", http.StatusBadRequest)
	}
	if utf8.RuneCountInString(input.Nickname) > 20 {
		http.Error(w, "nickname can not be longer than 20 characters", http.StatusBadRequest)
	}
	user, err := handler.registrar.Register(input.Nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json, err := json2.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(json)
	if err != nil {
		return
	}
}
