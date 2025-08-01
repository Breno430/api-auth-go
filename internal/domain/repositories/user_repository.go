package repositories

import (
	"context"

	"api-auth-go/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByID(ctx context.Context, id string) (*entities.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Update(ctx context.Context, user *entities.User) error
	FindAll(ctx context.Context) ([]*entities.User, error)
	FindAllWithFilters(ctx context.Context, filters *entities.UserFilters) ([]*entities.User, error)
	Delete(ctx context.Context, id string) error
}
