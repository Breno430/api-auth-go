package usecases

import (
	"context"
	"errors"

	"api-auth-go/internal/domain/entities"
	"api-auth-go/internal/domain/repositories"
	"api-auth-go/internal/infrastructure/services"
)

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateUserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Token     string `json:"token"`
}

type UserUseCase struct {
	userRepo   repositories.UserRepository
	jwtService *services.JWTService
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo:   userRepo,
		jwtService: services.NewJWTService(),
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	user, err := entities.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	exists, err := uc.userRepo.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	if err := entities.ValidateLoginData(input.Email, input.Password); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.CheckPassword(input.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := uc.jwtService.GenerateToken(user.ID.String(), user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Token:     token,
	}, nil
}
