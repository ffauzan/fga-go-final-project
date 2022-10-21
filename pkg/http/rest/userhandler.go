package rest

import (
	"final-project/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gt=8"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
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
	// Bind request body to UpdateUserRequest struct
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get userID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Call service to update user
	user, err := h.userService.UpdateUser(currentUserID, &domain.UpdateUserRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"email":    user.Email,
		"username": user.Username,
	})

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Get userID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	err := h.userService.DeleteUser(uint(currentUserID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Your account has been successfully deleted",
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Get userID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	user, err := h.userService.GetUserByID(uint(currentUserID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
