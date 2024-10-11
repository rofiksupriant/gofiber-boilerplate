package usecase

import (
	"boilerplate/internal/entiity"
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"boilerplate/internal/security"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	user, err := uc.Repository.FindByUsername(uc.DB, request.Username)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrNotFound.Code, "User tidak ditemukan")
	}

	wrongPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if wrongPass != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Password Salah")
	}

	// generate new token
	token, err := security.NewToken(user.Username, user.Role)
	if err != nil {
		log.Errorf("error generate new token %v", err)
		return nil, err
	}

	// encrypt token
	encryptedToken, err := security.Encrypt(*token)
	if err != nil {
		log.Errorf("error encrypt token %v", err)
		return nil, err
	}

	// build response
	response := new(model.LoginResponse)
	response.Username = request.Username
	response.Role = "ADMIN"
	response.Token = *encryptedToken
	response.Avatar = ""

	return response, nil
}

func (uc *UserUseCase) CreateUser(request *model.CreateUserRequest) error {
	err := uc.Validate.Struct(&request)
	if err != nil {
		log.Errorf("create user request validation failed %v", err)
		return err
	}

	user := &entiity.User{
		Username:  request.Username,
		Password:  request.Password,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = uc.Repository.Create(uc.DB, user); err != nil {
		log.Errorf("failed create new user %v", err)
		return err
	}

	return nil
}

func (uc *UserUseCase) Verify(username string) (*model.Auth, error) {
	user, err := uc.Repository.FindByUsername(uc.DB, username)
	if err != nil {
		log.Warnf("failed find by username %v", err)
		return nil, err
	}

	return &model.Auth{ID: user.ID}, nil
}

func (uc *UserUseCase) Current(c *fiber.Ctx) (*model.UserResponse, error) {
	auth := c.Locals("auth").(*model.Auth)

	user := new(entiity.User)
	if err := uc.Repository.FindById(uc.DB, user, auth.ID); err != nil {
		log.Warn("failed get user by id")
		return nil, fiber.ErrUnauthorized
	}

	response := model.UserResponse{
		Username: user.Username,
		Role:     user.Role,
		Avatar:   user.Avatar,
	}
	return &response, nil
}
