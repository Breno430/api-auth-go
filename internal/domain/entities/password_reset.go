package entities

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Token     string    `json:"token" gorm:"not null;uniqueIndex"`
	Email     string    `json:"email" gorm:"not null"`
	Used      bool      `json:"used" gorm:"default:false"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewPasswordReset(userID uuid.UUID, email string) (*PasswordReset, error) {
	if err := ValidatePasswordResetData(email); err != nil {
		return nil, err
	}

	token := generatePIN()

	return &PasswordReset{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		Email:     email,
		Used:      false,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}, nil
}

func (pr *PasswordReset) IsExpired() bool {
	return time.Now().After(pr.ExpiresAt)
}

func (pr *PasswordReset) IsValid() bool {
	return !pr.Used && !pr.IsExpired()
}

func (pr *PasswordReset) MarkAsUsed() {
	pr.Used = true
}

func (pr *PasswordReset) ValidateToken(token string) error {
	if pr.Token != token {
		return errors.New("invalid token")
	}

	if !pr.IsValid() {
		return errors.New("token is expired or already used")
	}

	return nil
}

func ValidateNewPassword(password string) error {
	if password == "" {
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

func generatePIN() string {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "123456"
	}
	return fmt.Sprintf("%06d", n.Int64())
}

func ValidatePasswordResetData(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}
