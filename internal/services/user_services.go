package services

import (
	"errors"
	"time"

	"github.com/ayserragm/backend-project/internal/config"
	"github.com/ayserragm/backend-project/internal/db"
	"github.com/ayserragm/backend-project/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	cfg *config.Config
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{cfg: cfg}
}

func (s *UserService) Register(username, email, password string) (*models.User, error) {
	// email/username benzersiz mi?
	var count int64
	db.DB.Model(&models.User{}).Where("email = ? OR username = ?", email, username).Count(&count)
	if count > 0 {
		return nil, errors.New("username or email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		Role:         "user",
	}
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(emailOrUsername, password string) (string, *models.User, error) {
	var user models.User
	err := db.DB.Where("email = ? OR username = ?", emailOrUsername, emailOrUsername).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, errors.New("invalid credentials")
	}
	if err != nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Duration(s.cfg.AccessTokenTTLMin) * time.Minute).Unix(),
		"email": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", nil, err
	}
	return ss, &user, nil
}
