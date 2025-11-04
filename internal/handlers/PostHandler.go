package handlers

import (
	"encoding/json"
	"io"
	"microblog-normal/internal/service"
	"net/http"
)

type PostHandler struct {
	Publisher service.Publisher
}

func NewPostHandler(publisher service.Publisher) *PostHandler {
	return &PostHandler{Publisher: publisher}
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var input struct {
		Author string `json:"author"`
		Text   string `json:"text"`
	}
	if err := json.Unmarshal(body, &input); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if input.Author == "" {
		http.Error(w, "Author nickname can't be empty", http.StatusBadRequest)
	}
	if input.Text == "" {
		http.Error(w, "Text can't be empty", http.StatusBadRequest)
	}
	post, err := handler.Publisher.Publish(input.Author, input.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} // использовать json.NewEncoder Предпосчительнее тк как 1) Выгоднее по памяти нет аллокации в []byte
}
