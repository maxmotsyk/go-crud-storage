package stor

import (
	"database/sql"
	"fmt"
	"gocrud/internal/domain"

	log "github.com/sirupsen/logrus"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(d *sql.DB) *Storage {
	return &Storage{
		DB: d,
	}
}

func (s *Storage) CreateUser(u *domain.User) error {
	_, err := s.DB.Exec(`
        INSERT INTO users (id,name, lastName, age, email)
        VALUES ($1, $2, $3, $4, $5)
    `, u.Id, u.Name, u.LastName, u.Age, u.Email)

	if err != nil {
		log.Printf("CreateUser DB error: %v", err) // <<< добавь лог
		return err
	}

	log.WithFields(log.Fields{
		"layer":   "storage",
		"action":  "createUser",
		"user_id": u.Id,
	}).Info("User created successfully")

	return nil
}

func (s *Storage) GetUser(id int64) (*sql.Rows, error) {
	row, err := s.DB.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "getUser",
	}).Info(fmt.Sprintf("Successfully retrieved user with id %d", id))

	return row, nil
}

func (s *Storage) UpdateUser(u *domain.User, id int64) error {
	_, err := s.DB.Exec("UPDATE users SET  name = $1, lastName = $2, age = $3 , email = $4 WHERE id = $5",
		u.Name, u.LastName, u.Age, u.Email, id)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "updateUser",
	}).Info(fmt.Sprintf("Successfully update user with id %d", id))

	return nil
}

func (s *Storage) DeleteUser(id int64) error {
	_, err := s.DB.Exec("DELETE FROM  users WHERE id = $1", id)

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"layer":  "storage",
		"action": "deleteUser",
	}).Info(fmt.Sprintf("Successfully delete user with id %d", id))

	return nil
}
