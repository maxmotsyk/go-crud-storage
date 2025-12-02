package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"gocrud/internal/domain"
	"gocrud/internal/stor"
)

// UserHandler отвечает за обработку HTTP-запросов, связанных с пользователями.
type UserHandler struct {
	s *stor.Storage // зависимость: слой работы с БД
}

// NewUserHandler конструктор, прокидываем сюда Storage.
func NewUserHandler(newStor *stor.Storage) *UserHandler {
	return &UserHandler{
		s: newStor,
	}
}

// CreateUser обрабатывает запрос POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Создаём пустую структуру, в которую будем декодировать JSON
	var user domain.User

	// Декодируем JSON из тела запроса сразу в структуру
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// Если JSON кривой — возвращаем 400
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	// Вызываем метод уровня хранения (Storage), который пишет в БД
	if err := h.s.CreateUser(&user); err != nil {
		// Если БД вернула ошибку — это уже 500
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	// Успех: возвращаем 201 + созданного пользователя
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

// GetUser обрабатывает запрос GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Достаём id из URL: /users/{id}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Если id не число — 400
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// Запрашиваем данные пользователя из БД
	rows, err := h.s.GetUser(id)
	if err != nil {
		// Ошибка уровня БД — 500
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user domain.User

	// Проверяем, есть ли строка с таким id
	if !rows.Next() {
		// Пользователь не найден — 404
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// Сканируем значения из строки в структуру user
	// Порядок полей должен совпадать с SELECT * FROM users:
	// id, name, lastName, age, email
	if err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Age, &user.Email); err != nil {
		http.Error(w, "scan error", http.StatusInternalServerError)
		return
	}

	// Возвращаем пользователя в JSON
	_ = json.NewEncoder(w).Encode(user)
}

// UpdateUser обрабатывает запрос PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Берём id из URL: /users/{id}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// Декодируем новые данные пользователя из body
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	// Вызываем метод обновления в Storage, передаём id отдельно
	if err := h.s.UpdateUser(&user, id); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	// Проставляем id в структуру, чтобы вернуть актуальные данные
	user.Id = id

	// Возвращаем обновлённого пользователя
	_ = json.NewEncoder(w).Encode(user)
}
