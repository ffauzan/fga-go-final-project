package rest

import (
	"final-project/pkg/domain"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRouter(userService *domain.UserService) *gin.Engine {
	r := gin.Default()

	// User handler routes
	userHandler := NewUserHandler(*userService)
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	return r
}

// // Function to send error response
func SendErrorResponse(c *gin.Context, err error, code int) {
	c.JSON(code, BaseResponse{
		Status:  "error",
		Message: err.Error(),
		Data:    nil,
	})
}
