package stor

import (
	"database/sql"
	"gocrud/internal/domain"
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
	_, err := s.DB.Exec("insert into users (name, lastName, age, email ) values ($1, $2, $3, $4)",
		u.Name, u.LastName, u.Age, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUser(id int64) (*sql.Rows, error) {
	row, err := s.DB.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (s *Storage) UpdateUser(u *domain.User, id int64) error {
	_, err := s.DB.Exec("UPDATE users SET name = $1, lastName = $2, age = $3 , email = $4 WHERE id = $5",
		u.Name, u.LastName, u.Age, u.Email, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUser(id int64) error {
	_, err := s.DB.Exec("delete from users where id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
