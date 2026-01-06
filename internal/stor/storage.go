package stor

import (
	"database/sql"
	"fmt"
	"gocrud/internal/domain"
	"time"

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

func (s *Storage) CreateUser(u *domain.SignUpInput) error {
	registered_time := time.Now()
	_, err := s.DB.Exec(`
        INSERT INTO users (name, lastName, age, email, password, registered_time)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, u.Name, u.LastName, u.Age, u.Email, u.Password, registered_time)

	if err != nil {
		log.Printf("CreateUser DB error: %v", err) // <<< добавь лог
		return err
	}

	log.WithFields(log.Fields{
		"layer":   "storage",
		"action":  "createUser",
		"user_id": u.Email,
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
	//ToDo UpdateUser in struct.go before fix this handler
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
