package rest

import (
	"final-project/pkg/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register is a handler for user registration end point
func (h *UserHandler) Register(c *gin.Context) {
	// Bind request body to RegisterRequest struct
	// TODO: Add validation
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Call service to register user
	user, err := h.userService.Register(&domain.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
		Password: req.Password,
	})
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	})
}

// Login is a handler for user login end point
func (h *UserHandler) Login(c *gin.Context) {
	// Bind request body to LoginRequest struct
	// TODO: Add validation
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	token, err := h.userService.Login(&domain.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// UpdateUser is a handler for updating user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(uint(userID), &domain.UpdateUserRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Status:  "success",
		Message: "user updated",
		Data:    user,
	})

}
