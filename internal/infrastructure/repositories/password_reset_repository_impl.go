package repositories

import (
	"context"
	"time"

	"api-auth-go/internal/domain/entities"
	"api-auth-go/internal/domain/repositories"

	"gorm.io/gorm"
)

type PasswordResetRepositoryImpl struct {
	db *gorm.DB
}

func NewPasswordResetRepositoryImpl(db *gorm.DB) repositories.PasswordResetRepository {
	return &PasswordResetRepositoryImpl{db: db}
}

func (r *PasswordResetRepositoryImpl) Create(ctx context.Context, passwordReset *entities.PasswordReset) error {
	return r.db.WithContext(ctx).Create(passwordReset).Error
}

func (r *PasswordResetRepositoryImpl) FindByToken(ctx context.Context, token string) (*entities.PasswordReset, error) {
	var passwordReset entities.PasswordReset
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&passwordReset).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &passwordReset, nil
}

func (r *PasswordResetRepositoryImpl) FindByUserID(ctx context.Context, userID string) (*entities.PasswordReset, error) {
	var passwordReset entities.PasswordReset
	err := r.db.WithContext(ctx).Where("user_id = ? AND used = false AND expires_at > ?", userID, time.Now()).First(&passwordReset).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &passwordReset, nil
}

func (r *PasswordResetRepositoryImpl) Update(ctx context.Context, passwordReset *entities.PasswordReset) error {
	return r.db.WithContext(ctx).Save(passwordReset).Error
}

func (r *PasswordResetRepositoryImpl) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&entities.PasswordReset{}).Error
}
