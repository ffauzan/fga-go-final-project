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

func NewRouter(
	userService *domain.UserService,
	authService *domain.AuthService,
	photoService *domain.PhotoService,
) *gin.Engine {
	r := gin.Default()

	// User handler routes
	userHandler := NewUserHandler(*userService)
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)

		protectedUserRouter := userRouter.Group("/")
		{
			protectedUserRouter.Use(AuthMiddleware(*authService))
			protectedUserRouter.PUT("/:id", userHandler.UpdateUser)
			protectedUserRouter.DELETE("/", userHandler.DeleteUser)
			protectedUserRouter.GET("/", userHandler.GetUser)
		}
	}

	// Photo handler routes
	photoHandler := NewPhotoHandler(*photoService, *userService)
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(AuthMiddleware(*authService))
		photoRouter.POST("/", photoHandler.AddPhoto)
		photoRouter.GET("/", photoHandler.GetPhotos)
		photoRouter.PUT("/:id", photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:id", photoHandler.DeletePhoto)
	}

	return r
}

// Function to send error response
func SendErrorResponse(c *gin.Context, err error, code int) {
	c.JSON(code, BaseResponse{
		Status:  "error",
		Message: err.Error(),
		Data:    nil,
	})
}
