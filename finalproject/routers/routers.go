package routers

import (
	"finalproject/database"
	"finalproject/docs"
	"finalproject/handlers"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup() *gin.Engine {

	docs.SwaggerInfo.Title = "Example Swagger User Rest API"
	docs.SwaggerInfo.Description = "Documentation of User Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	re := r.Group("/todos")
	{
		re.GET("", api.GetTodos)
		re.GET("/:id", api.GetTodo)
		re.POST("", api.CreateTodo)
		re.PUT("/:id", api.UpdateTodo)
		re.DELETE("/:id", api.DeleteTodo)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
