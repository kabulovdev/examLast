package v1

import (
	"swagger/examLast/exam_api/config"
	"fmt"
	jwthandler "swagger/examLast/exam_api/api/tokens"
	"swagger/examLast/exam_api/pkg/logger"
	t "swagger/examLast/exam_api/api/tokens"
	"swagger/examLast/exam_api/storage/repo"
	"swagger/examLast/exam_api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redis          repo.InMemoryStorageI
	jwthandler     t.JWTHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.InMemoryStorageI
	JWTHandler     t.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redis:          c.Redis,
		jwthandler:     c.JWTHandler,
		
	}
}
	func GetClaims(h handlerV1, c *gin.Context) (*jwthandler.CustomClaims, error) {

		var (
			claims = jwthandler.CustomClaims{}
		)
		
		strToken := c.GetHeader("Authorization")
		fmt.Println(h.cfg.SignKey)
		
		token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {return []byte(h.cfg.SignKey), nil})
		
		if err != nil {
			fmt.Println(err)
			h.log.Error("invalid access token")
			return nil, err
		}
		// rawClaims := token.Claims.(jwt.MapClaims)
	
		// claims.Sub = rawClaims["sub"].(string)
		// claims.Exp = rawClaims["exp"].(float64)
		// fmt.Printf("%T type of value in map %v",rawClaims["exp"],rawClaims["exp"])
		// fmt.Printf("%T type of value in map %v",rawClaims["iat"],rawClaims["iat"])
	
		claims.Token = token
		return &claims, nil
	
	}
