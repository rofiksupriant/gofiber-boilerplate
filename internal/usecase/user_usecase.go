package usecase

import (
	"boilerplate/internal/entiity"
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"context"
	"crypto/aes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"os"
	"time"
)

type UserUseCase struct {
	DB         *gorm.DB
	Validate   *validator.Validate
	Repository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:         db,
		Validate:   validate,
		Repository: userRepository,
	}
}

func (uc *UserUseCase) Login(request *model.LoginRequest) (*model.LoginResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		log.Errorf("validate login err %v", err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   request.Username,
		Audience:  jwt.ClaimStrings{"ADMIN"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	secretKey := []byte(os.Getenv("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"))
	signedToken, errSignToken := token.SignedString(secretKey)
	if errSignToken != nil {
		return nil, errSignToken
	}

	cipher, errCipher := aes.NewCipher([]byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"))
	if errCipher != nil {
		return nil, errCipher
	}

	cipherText := make([]byte, aes.BlockSize+len(signedToken))
	cipher.Encrypt(cipherText, []byte(signedToken))

	response := new(model.LoginResponse)
	response.Username = request.Username
	response.Role = "ADMIN"
	response.Token = signedToken
	response.Avatar = ""

	return response, nil
}

func (uc *UserUseCase) CreateUser(ctx context.Context, request *model.CreateUserRequest) error {
	err := uc.Validate.Struct(&request)
	if err != nil {
		log.Errorf("create user request validation failed %v", err)
		return fiber.ErrBadRequest
	}

	db := uc.DB.WithContext(ctx).Begin()
	defer db.Rollback()

	user := &entiity.User{
		Username:  request.Username,
		Password:  request.Password,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = uc.Repository.Create(db, user); err != nil {
		log.Errorf("failed create new user %v", err)
		return err
	}

	return nil
}
