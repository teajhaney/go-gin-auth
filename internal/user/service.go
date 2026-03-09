package user

import (
	"context"
	"errors"
	"fmt"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


type service struct {
	repo *Repo

	jwtSecret string
}


func NewService(repo *Repo, jwtSecret string) *service {
	return &service{
		repo: repo,
		jwtSecret: jwtSecret,
	}
}


type RegisterInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type AuthResponse struct {
	Token string `json:"token"`
	User PublicUser `json:"user"`
}


func (s *service) Register(ctx context.Context, input RegisterInput) (AuthResponse, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	if email == "" || password == "" {
		return AuthResponse{}, errors.New("email and password are required")
	}

	if len(password) < 6 {
		return AuthResponse{}, errors.New("password must be at least 6 characters")
	}

// Check if user already exists
	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResponse{}, errors.New("user already exists, please try with a different email")
	}

	hashByte, err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error hashing password: %v", err)
	}


	 now := time.Now().UTC()

	user := User{
		Email: email,
		Password: string(hashByte),
		Role: "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	token, err := auth.CreateToken(s.jwtSecret, createdUser.ID.Hex(), createdUser.Role)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error creating token: %v", err)
	}
	return AuthResponse{
		Token: token, 
		User: ToPublicUser(createdUser),
		}, nil
}


func (s *service) Login(ctx context.Context, input LoginInput) (AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	if email == "" || password == "" {
		return AuthResponse{}, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return AuthResponse{}, errors.New("invalid email or password")
		}
		return AuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return AuthResponse{}, errors.New("invalid email or password")
	}

	token, err := auth.CreateToken(s.jwtSecret, user.ID.Hex(), user.Role)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("error creating token: %v", err)
	}

	return AuthResponse{
		Token: token,
		User: ToPublicUser(user),
	}, nil
}
