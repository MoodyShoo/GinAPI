package http

import (
	"net/http"

	"github.com/MoodyShoo/GinAPI/internal/http/handler"
	"github.com/MoodyShoo/GinAPI/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter(userService service.UserService, log *logrus.Logger) *gin.Engine {
	r := gin.New()

	r.Use(func(c *gin.Context) {
		log.Infof("=> %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		status := c.Writer.Status()
		log.Infof("<= %d %s", status, http.StatusText(status))
	})

	userHandler := handler.NewUserHandler(userService)

	api := r.Group("/api/users")
	{
		api.GET("/:id", userHandler.GetUser)
		api.GET("/:id/:field", userHandler.GetUserField)
	}

	testGroup := r.Group("/test")
	{
		testGroup.GET("/fill", func(c *gin.Context) {
			count, err := userService.SeedTestUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Тестовые пользователи добавлены", "count": count})
		})

		testGroup.GET("/erase", func(c *gin.Context) {
			count, err := userService.EraseTestUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Тестовые пользователи удалены", "count": count})
		})
	}

	return r
}
