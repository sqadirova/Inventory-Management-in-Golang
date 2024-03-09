package user

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"im-backoffice/CONSTANTS"
	myErrors "im-backoffice/errors"
	"im-backoffice/middleware"
	"log"
)

type UserController struct {
	service   IUserService
	validator *validator.Validate
}

func NewUserController(service IUserService) UserController {
	return UserController{
		service:   service,
		validator: validator.New(),
	}
}

// GetAllUsers godoc
// @Summary Get all users.
// @Description get all users from users table in database
// @Tags user
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} UserResp
// @Failure 404 {object} errors.Response
// @Router /users [get]
func (uc UserController) GetAllUsers(ctx *fiber.Ctx) error {
	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	users, err := uc.service.getAllUsers(ctx.Context(), token)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

// GetUserMe godoc
// @Summary Get user info by token in the request header.
// @Description get user info from database by token in the request header
// @Tags user
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} UserMeResp
// @Failure 400 {object} errors.Response
// @Failure 401 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /user/me [get]
func (uc UserController) GetUserMe(ctx *fiber.Ctx) error {
	accessToken := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	if accessToken == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(myErrors.NewResponseByKey("bearer_token_missing", CONSTANTS.LANGUAGE))
	}

	token, _, err := middleware.VerifyAccessToken(accessToken)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(myErrors.NewResponseByKey(err.Error(), CONSTANTS.LANGUAGE))
	}

	userInfo, userRole, err := uc.service.GetUserInfo(ctx.Context(), token)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(userToUserMeResp(userInfo, userRole))
}

// GetAllRoles godoc
// @Summary Get all roles.
// @Description get all roles from roles table in database
// @Tags user
// @Accept */*
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} RolesResp
// @Failure 404 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /user/roles [get]
func (uc UserController) GetAllRoles(ctx *fiber.Ctx) error {
	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	roles, err := uc.service.getRolesFromDB(ctx.Context(), token)

	if err != nil {
		log.Println(err)
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(rolesToRolesRespArr(roles))
}

// GetOneRole godoc
// @Summary Get one role.
// @Description get one role by id from database
// @Tags user
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} RolesResp
// @Failure 404 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Param id path string true "Role ID"
// @Router /user/roles/{id} [get]
func (uc UserController) GetOneRole(ctx *fiber.Ctx) error {
	roleId := ctx.Params("id")

	if roleId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	role, err := uc.service.getRoleByIDFromDB(ctx.Context(), roleId, token)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(roleToRolesResp(role))
}

// CreateUser godoc
// @Summary Create a user.
// @Description Create a user in database
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id body UserDTO true "UserDTO"
// @Success 201 {object} UserResp
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /user [post]
func (uc UserController) CreateUser(ctx *fiber.Ctx) error {
	userReq := UserDTO{}
	err := ctx.BodyParser(&userReq)

	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	user, err := uc.service.createUser(ctx.Context(), userReq, token)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user in database
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param id body user.UpdateUserReq true "UpdateUserReq"
// @Success 200 {object} UserResp
// @Failure 400 {object} errors.Response
// @Failure 304 {object} errors.Response
// @Failure 500 {object} errors.Response
// @Router /user/{id} [put]
func (uc UserController) UpdateUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	if userId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_id", CONSTANTS.LANGUAGE))
	}

	var user UpdateUserReq

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(myErrors.NewResponseByKey("invalid_data", CONSTANTS.LANGUAGE))
	}

	token := middleware.ExtractBearerToken(ctx.GetReqHeaders()["Authorization"])

	updatedUser, err := uc.service.updateUser(ctx.Context(), userId, user, token)

	if err != nil {
		return ctx.Status(err.(myErrors.HttpError).Code).JSON(err.(myErrors.HttpError).Response)
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedUser)
}
