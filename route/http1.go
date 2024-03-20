package httpv1

import (
	"log"
	"movie-app/app/presenter"
	env "movie-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutesGin(env *env.Dependency) *gin.Engine {
	rD := gin.New()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("[INFO] endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	rD.Use(CORSMiddleware())

	rC := rD.Group("api/", gin.Logger())
	{
		movieGroup(rC.Group("/movies"), env)
	}

	return rD
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Client-ID, Client-Module, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

func movieGroup(group *gin.RouterGroup, env *env.Dependency) {
	serviceRunning("movie")
	group.GET("/", presenter.List(env))
	group.GET("/:id", presenter.Detail(env))
	group.PATCH("/", presenter.Updated(env))
	group.DELETE("/:id", presenter.Remove(env))
	group.POST("/add", presenter.Add(env))
}

func serviceRunning(m string) {
	log.Printf("[INFO] %s service is running ...... ", m)
}
