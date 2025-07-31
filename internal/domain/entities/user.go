package entities

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewUser(name, email, password string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}
	if len(name) < 2 || len(name) > 100 {
		return nil, errors.New("name must be between 2 and 100 characters")
	}

	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password is required")
	}
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func ValidateLoginData(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	if len(email) > 255 {
		return errors.New("email is too long (maximum 255 characters)")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password is required")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if len(password) > 128 {
		return errors.New("password is too long (maximum 128 characters)")
	}

	return nil
}

func ValidateResetPasswordInput(token, password string) error {
	if strings.TrimSpace(token) == "" {
		return errors.New("token is required")
	}

	if len(token) != 6 {
		return errors.New("token must be 6 digits")
	}

	for _, char := range token {
		if char < '0' || char > '9' {
			return errors.New("token must contain only digits")
		}
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password is required")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if len(password) > 128 {
		return errors.New("password is too long (maximum 128 characters)")
	}

	return nil
}
