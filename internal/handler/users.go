package handler

import (
	todo "Fragaed"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var validate = validator.New()

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input todo.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		// Обработка ошибки при декодировании JSON
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.Users.CreateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр id из URL
	idStr := chi.URLParam(r, "id")

	// Преобразуем idStr в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Получаем пользователя по id
	u, err := h.service.Users.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ с данными пользователя в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var input todo.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	if err := validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, vErr := range validationErrors {
			errorMessage := fmt.Sprintf("%s is %s", vErr.Field(), vErr.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}
	input.Id = id
	if err := h.service.Users.UpdateUser(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	answer, err := h.service.Users.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answer)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	var cond todo.Conditions
	if err := json.NewDecoder(r.Body).Decode(&cond); err != nil {
	}
	answer, err := h.service.ListAllUsers(cond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answer)
}
