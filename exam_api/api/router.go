package api

import (
	v1 "examLast/exam_api/api/handlers/v1"
	"examLast/exam_api/config"
	"examLast/exam_api/pkg/logger"
	"examLast/exam_api/services"

	_ "examLast/exam_api/api/docs" //swag

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/custumer/create", handlerV1.CreateCustumer)
	api.GET("/custumer/get/:id", handlerV1.GetCustumer)
	api.DELETE("/custumer/delete/:id", handlerV1.DeleteCustumer)
	api.POST("/custumer/update", handlerV1.UpdateCustumer)
	api.POST("/post/update",handlerV1.UpdatePost)
	api.POST("/post/create", handlerV1.CreatePost)
	api.POST("/reating/create",handlerV1.CreateReating)
	api.POST("/reating/update",handlerV1.UpdateReating)
	api.DELETE("/reating/delete/:id",handlerV1.DeleteReating)
	api.DELETE("/post/delet/:id",handlerV1.DeletePost)
	api.GET("/post/allInfo/:id",handlerV1.GetPostAllInfo)
	api.GET("/reating/get/:id",handlerV1.GetReating)
	api.GET("/post/get/reatings/:id",handlerV1.GetPostReating)
	api.GET("/custumer/getList",handlerV1.GetListCustumers)
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler, url))
	return router

}
