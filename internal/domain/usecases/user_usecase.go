package usecases

import (
	"context"
	"errors"
	"fmt"

	"api-auth-go/internal/domain/entities"
	"api-auth-go/internal/domain/repositories"
	"api-auth-go/internal/infrastructure/services"

	"golang.org/x/crypto/bcrypt"
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

type ResetPasswordInput struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type ResetPasswordOutput struct {
	Message string `json:"message"`
}

type RequestPasswordResetInput struct {
	Email string `json:"email" validate:"required,email"`
}

type RequestPasswordResetOutput struct {
	Message string `json:"message"`
}

type UserUseCase struct {
	userRepo          repositories.UserRepository
	passwordResetRepo repositories.PasswordResetRepository
	jwtService        *services.JWTService
	emailService      *services.EmailService
}

func NewUserUseCase(userRepo repositories.UserRepository, passwordResetRepo repositories.PasswordResetRepository) *UserUseCase {
	return &UserUseCase{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
		jwtService:        services.NewJWTService(),
		emailService:      services.NewEmailService(),
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

func (uc *UserUseCase) RequestPasswordReset(ctx context.Context, input RequestPasswordResetInput) (*RequestPasswordResetOutput, error) {
	if err := entities.ValidatePasswordResetData(input.Email); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &RequestPasswordResetOutput{
			Message: "Se o email existir, você receberá um código de verificação por email.",
		}, nil
	}

	existingReset, err := uc.passwordResetRepo.FindByUserID(ctx, user.ID.String())
	if err != nil {
		return nil, err
	}

	if existingReset != nil && existingReset.IsValid() {
		return &RequestPasswordResetOutput{
			Message: "Um código de verificação já foi enviado. Aguarde 15 minutos para solicitar um novo.",
		}, nil
	}

	passwordReset, err := entities.NewPasswordReset(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	if err := uc.passwordResetRepo.Create(ctx, passwordReset); err != nil {
		return nil, err
	}

	go func() {
		if err := uc.emailService.SendPasswordResetEmail(user.Email, user.Name, passwordReset.Token); err != nil {
			fmt.Printf("Erro ao enviar email: %v\n", err)
		}
	}()

	return &RequestPasswordResetOutput{
		Message: "Código de verificação enviado por email.",
	}, nil
}

func (uc *UserUseCase) ResetPassword(ctx context.Context, input ResetPasswordInput) (*ResetPasswordOutput, error) {
	if err := entities.ValidateResetPasswordInput(input.Token, input.Password); err != nil {
		return nil, err
	}

	passwordReset, err := uc.passwordResetRepo.FindByToken(ctx, input.Token)
	if err != nil {
		return nil, err
	}
	if passwordReset == nil {
		return nil, errors.New("invalid or expired token")
	}

	user, err := uc.userRepo.FindByID(ctx, passwordReset.UserID.String())
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := passwordReset.ValidateToken(input.Token); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	passwordReset.MarkAsUsed()
	if err := uc.passwordResetRepo.Update(ctx, passwordReset); err != nil {
		return nil, err
	}

	return &ResetPasswordOutput{
		Message: "Senha alterada com sucesso.",
	}, nil
}
