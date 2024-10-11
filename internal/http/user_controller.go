package http

import (
	"boilerplate/internal/model"
	"boilerplate/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(uc *usecase.UserUseCase) *UserController {
	return &UserController{UseCase: uc}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON((&model.ApiResponse[any]{}).BadRequest(ctx, err))
	}

	loginResponse, err := c.UseCase.Login(request)
	if err != nil {
		return ctx.JSON((&model.ApiResponse[any]{}).BadRequest(ctx, err))
	}

	return ctx.JSON((&model.ApiResponse[*model.LoginResponse]{}).Success("Login berhasil", loginResponse))
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed parse request body %v", err)
		return fiber.ErrBadRequest
	}

	return ctx.JSON((&model.ApiResponse[any]{}).Success("berhasil menyimpan data", nil))
}
