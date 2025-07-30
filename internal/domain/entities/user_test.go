package entities

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		password string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid user",
			userName: "John Doe",
			email:    "john@example.com",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "empty name",
			userName: "",
			email:    "john@example.com",
			password: "password123",
			wantErr:  true,
			errMsg:   "name is required",
		},
		{
			name:     "name with only spaces",
			userName: "   ",
			email:    "john@example.com",
			password: "password123",
			wantErr:  true,
			errMsg:   "name is required",
		},
		{
			name:     "name too short",
			userName: "J",
			email:    "john@example.com",
			password: "password123",
			wantErr:  true,
			errMsg:   "name must be between 2 and 100 characters",
		},
		{
			name:     "name too long",
			userName: "This is a very long name that exceeds the maximum allowed length of one hundred characters and should cause an error",
			email:    "john@example.com",
			password: "password123",
			wantErr:  true,
			errMsg:   "name must be between 2 and 100 characters",
		},
		{
			name:     "empty email",
			userName: "John Doe",
			email:    "",
			password: "password123",
			wantErr:  true,
			errMsg:   "email is required",
		},
		{
			name:     "email with only spaces",
			userName: "John Doe",
			email:    "   ",
			password: "password123",
			wantErr:  true,
			errMsg:   "email is required",
		},
		{
			name:     "invalid email format - missing @",
			userName: "John Doe",
			email:    "johnexample.com",
			password: "password123",
			wantErr:  true,
			errMsg:   "invalid email format",
		},
		{
			name:     "invalid email format - missing domain",
			userName: "John Doe",
			email:    "john@",
			password: "password123",
			wantErr:  true,
			errMsg:   "invalid email format",
		},
		{
			name:     "invalid email format - missing TLD",
			userName: "John Doe",
			email:    "john@example",
			password: "password123",
			wantErr:  true,
			errMsg:   "invalid email format",
		},
		{
			name:     "empty password",
			userName: "John Doe",
			email:    "john@example.com",
			password: "",
			wantErr:  true,
			errMsg:   "password is required",
		},
		{
			name:     "password with only spaces",
			userName: "John Doe",
			email:    "john@example.com",
			password: "   ",
			wantErr:  true,
			errMsg:   "password is required",
		},
		{
			name:     "password too short",
			userName: "John Doe",
			email:    "john@example.com",
			password: "123",
			wantErr:  true,
			errMsg:   "password must be at least 6 characters",
		},
		{
			name:     "password exactly 5 characters",
			userName: "John Doe",
			email:    "john@example.com",
			password: "12345",
			wantErr:  true,
			errMsg:   "password must be at least 6 characters",
		},
		{
			name:     "password exactly 6 characters",
			userName: "John Doe",
			email:    "john@example.com",
			password: "123456",
			wantErr:  false,
		},
		{
			name:     "valid email with special characters",
			userName: "John Doe",
			email:    "john.doe+test@example.co.uk",
			password: "password123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.userName, tt.email, tt.password)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewUser() expected error but got none")
					return
				}
				if tt.errMsg != "" && err.Error() != tt.errMsg {
					t.Errorf("NewUser() error = %v, want %v", err.Error(), tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("NewUser() unexpected error: %v", err)
				return
			}

			if user == nil {
				t.Errorf("NewUser() expected user but got nil")
				return
			}

			if user.Name != tt.userName {
				t.Errorf("NewUser() name = %v, want %v", user.Name, tt.userName)
			}

			if user.Email != tt.email {
				t.Errorf("NewUser() email = %v, want %v", user.Email, tt.email)
			}

			if user.Password == tt.password {
				t.Errorf("NewUser() password should be hashed, not plain text")
			}

			// Verify that the password can be checked
			if !user.CheckPassword(tt.password) {
				t.Errorf("NewUser() created password hash cannot be verified")
			}
		})
	}
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "password123")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "correct password",
			password: "password123",
			want:     true,
		},
		{
			name:     "incorrect password",
			password: "wrongpassword",
			want:     false,
		},
		{
			name:     "empty password",
			password: "",
			want:     false,
		},
		{
			name:     "similar password",
			password: "password124",
			want:     false,
		},
		{
			name:     "case sensitive password",
			password: "PASSWORD123",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := user.CheckPassword(tt.password); got != tt.want {
				t.Errorf("User.CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
