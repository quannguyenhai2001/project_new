package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService định nghĩa các phương thức liên quan đến xác thực
type AuthService interface {
	Signup(userData *model.UserSignup) (*model.TokenResponse, error)
	Login(credentials *model.UserLogin) (*model.TokenResponse, error)
	GenerateToken(user *model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserFromToken(tokenString string) (*model.User, error)
}

// AuthServiceImpl là implementation của AuthService
type AuthServiceImpl struct {
	userRepo    repository.UserRepository
	jwtSecret   []byte
	tokenExpiry time.Duration
}

// NewAuthService tạo instance mới của AuthServiceImpl
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &AuthServiceImpl{
		userRepo:    userRepo,
		jwtSecret:   []byte(jwtSecret),
		tokenExpiry: time.Hour * 24, // Token valid trong 24h
	}
}

// Signup xử lý việc đăng ký người dùng mới
func (s *AuthServiceImpl) Signup(userData *model.UserSignup) (*model.TokenResponse, error) {
	// Kiểm tra email đã tồn tại chưa
	existingUser, err := s.userRepo.FindByEmail(userData.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Mã hóa password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Tạo user mới
	user := &model.User{
		Email:     userData.Email,
		Password:  string(hashedPassword),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}

	// Lưu vào database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Tạo token
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Tạo response
	return &model.TokenResponse{
		AccessToken: token,
		User: model.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

// Login xử lý việc đăng nhập
func (s *AuthServiceImpl) Login(credentials *model.UserLogin) (*model.TokenResponse, error) {
	// Tìm user theo email
	user, err := s.userRepo.FindByEmail(credentials.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Kiểm tra password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Tạo token
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Tạo response
	return &model.TokenResponse{
		AccessToken: token,
		User: model.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

// GenerateToken tạo JWT token cho user đã xác thực
func (s *AuthServiceImpl) GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(s.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

// ValidateToken xác thực JWT token
func (s *AuthServiceImpl) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức mã hóa
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return s.jwtSecret, nil
	})

	return token, err
}

// GetUserFromToken lấy thông tin user từ token
func (s *AuthServiceImpl) GetUserFromToken(tokenString string) (*model.User, error) {
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid user id in token")
	}

	return s.userRepo.FindByID(int64(userID))
}
