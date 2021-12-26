package mysql

import (
	"database/sql"
	"errors"
	"github.com/record-collection/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) GetUserById(id int) (*models.User, error) {
	s := &models.User{}

	err := u.DB.QueryRow("SELECT id, first_name, last_name, email, password, "+
		"access_level FROM users WHERE id = ?", id).Scan(&s.ID, &s.FirstName, &s.LastName,
		&s.Email, &s.Password, &s.AccessLevel)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, err
}

// Authenticate authenticates user
func (u *UserModel) Authenticate(email, testPassword string) (int, string, error) {

	var id int
	var hashedPassword string

	err := u.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(id,
		hashedPassword)

	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}
