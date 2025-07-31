package repositories

import (
	"context"

	"api-auth-go/internal/domain/entities"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, passwordReset *entities.PasswordReset) error
	FindByToken(ctx context.Context, token string) (*entities.PasswordReset, error)
	FindByUserID(ctx context.Context, userID string) (*entities.PasswordReset, error)
	Update(ctx context.Context, passwordReset *entities.PasswordReset) error
	DeleteExpired(ctx context.Context) error
}
