package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	uuid "github.com/satori/go.uuid"
	"github.com/twoflower3/interview-service/docs"

	log "github.com/sirupsen/logrus"
	"github.com/twoflower3/interview-service/pkg/handlers"
	"github.com/twoflower3/interview-service/pkg/server/routes"
)

// CreateRouter create router
func createRouter(s *Server) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	{
		e.Use(ginLogger(log.New()))
		e.Use(Recovery())
		e.Use(requestID())
		initCORS(e)
		initRoutes(e, s)
	}
	return e
}

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewV4()
		c.Set("Request-ID", id.String())
		c.Next()
	}
}

func initCORS(e *gin.Engine) {
	cfg := cors.DefaultConfig()
	{
		cfg.AllowAllOrigins = true
	}

	e.Use(cors.New(cfg))
}

func initRoutes(e gin.IRouter, srv *Server) {

	docs.SwaggerInfo.Title = "Interview service API"
	docs.SwaggerInfo.Description = "Interview service API"
	docs.SwaggerInfo.Version = "0.0.1-dev"
	docs.SwaggerInfo.Host = "localhost" + srv.Addr
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/", handlers.MainPage)
	e.GET("/swagger.json", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=\"swagger.json\"")
		c.File("./docs/swagger.json")
	})
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.API(e)
}
