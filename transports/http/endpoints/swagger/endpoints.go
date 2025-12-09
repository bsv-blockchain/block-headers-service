package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/bsv-blockchain/block-headers-service/config"
	"github.com/bsv-blockchain/block-headers-service/docs"
	"github.com/bsv-blockchain/block-headers-service/service"
	router "github.com/bsv-blockchain/block-headers-service/transports/http/endpoints/routes"
)

// NewHandler creates new endpoint handler.
func NewHandler(_ *service.Services, apiURLPrefix string) router.RootEndpoints {
	return router.RootEndpointsFunc(func(router *gin.RouterGroup) {
		docs.SwaggerInfo.BasePath = apiURLPrefix
		docs.SwaggerInfo.Version = config.Version()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	})
}
