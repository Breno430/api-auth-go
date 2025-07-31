package entities

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type UserFilters struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Role      string `json:"role" form:"role"`
	Page      int    `json:"page" form:"page"`
	Limit     int    `json:"limit" form:"limit"`
	SortBy    string `json:"sort_by" form:"sort_by"`
	SortOrder string `json:"sort_order" form:"sort_order"`
}

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;default:'user'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func ValidateUUID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id is required")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid UUID format")
	}

	return nil
}

func ValidateRole(role string) error {
	if strings.TrimSpace(role) == "" {
		return errors.New("role is required")
	}

	if role != RoleAdmin && role != RoleUser {
		return errors.New("role must be 'admin' or 'user'")
	}

	return nil
}

func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	if len(name) < 2 {
		return errors.New("name must be at least 2 characters")
	}

	if len(name) > 100 {
		return errors.New("name must be at most 100 characters")
	}

	nameRegex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)
	if !nameRegex.MatchString(name) {
		return errors.New("name must contain only letters and spaces")
	}

	return nil
}

func ValidateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}

	if len(email) > 255 {
		return errors.New("email is too long (maximum 255 characters)")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func ValidatePassword(password string) error {
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

func ValidateUserFilters(filters *UserFilters) error {
	if filters.Page < 1 {
		filters.Page = 1
	}

	if filters.Limit < 1 {
		filters.Limit = 10
	}

	if filters.Limit > 100 {
		filters.Limit = 100
	}

	if filters.SortBy == "" {
		filters.SortBy = "created_at"
	}

	validSortFields := []string{"name", "email", "role", "created_at", "updated_at"}
	isValidSortField := false
	for _, field := range validSortFields {
		if filters.SortBy == field {
			isValidSortField = true
			break
		}
	}

	if !isValidSortField {
		return errors.New("invalid sort_by field")
	}

	if filters.SortOrder == "" {
		filters.SortOrder = "desc"
	}

	if filters.SortOrder != "asc" && filters.SortOrder != "desc" {
		return errors.New("sort_order must be 'asc' or 'desc'")
	}

	if filters.Role != "" {
		if err := ValidateRole(filters.Role); err != nil {
			return err
		}
	}

	return nil
}

func NewUser(name, email, password string) (*User, error) {
	if err := ValidateName(name); err != nil {
		return nil, err
	}

	if err := ValidateEmail(email); err != nil {
		return nil, err
	}

	if err := ValidatePassword(password); err != nil {
		return nil, err
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
		Role:     RoleUser,
	}, nil
}

func NewAdminUser(name, email, password string) (*User, error) {
	user, err := NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	user.Role = RoleAdmin
	return user, nil
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func ValidateLoginData(email, password string) error {
	if err := ValidateEmail(email); err != nil {
		return err
	}

	if err := ValidatePassword(password); err != nil {
		return err
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

	if err := ValidatePassword(password); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateUserInput(name, email, role string) error {
	if err := ValidateName(name); err != nil {
		return err
	}

	if err := ValidateEmail(email); err != nil {
		return err
	}

	if err := ValidateRole(role); err != nil {
		return err
	}

	return nil
}
