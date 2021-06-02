package handler

import (
	"go-bake/auth"
	"go-bake/entity"
	"go-bake/helper"
	"go-bake/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

//show all user in database from route "/users"
func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("successget all user", 200, "status OK", users)
	c.JSON(200, response)
}

//user/external handler for create new user from "/user"
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// if err := helper.ValidatePasswowrd(inputUser.Password); err != nil {
	// 	responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": err.Error()})

	// 	c.JSON(400, responseError)
	// 	return
	// }

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {

		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create user", 201, "status created", newUser)
	c.JSON(200, response)
}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("success get user by id", 200, "succes", user)
	c.JSON(200, response)

}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{
			"errors": splitError})

		c.JSON(400, responseError)

		return
	}

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "User id not authorize"})

		c.JSON(401, responseError)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("update user succeed", 200, "success", user)
	c.JSON(200, response)
}

func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")
	user, err := h.userService.DeleteByUserID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "succes", user)
	c.JSON(200, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data required", http.StatusUnauthorized, "bad request", gin.H{"errors": err})

		c.JSON(400, responseError)
		return
	}

	// token, err := h.authService.GenerateToken(userData.ID)
	token, err := h.authService.GenerateToken(string(userData.Role))

	if err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusUnauthorized, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("succees login user", 200, "success", gin.H{"token": token, "data": userData})
	c.JSON(200, response)
}
