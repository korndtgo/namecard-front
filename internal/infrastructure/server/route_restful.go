package server

import (
	controller_v1 "card-service/internal/controller/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RouteRestfulService ...
type RouteRestfulService struct {
	dig.In

	GatewayV1 controller_v1.Gateway
}

func (c *Server) ConfigRouteRESTful(r *gin.Engine) {
	//g := r.Group("api/campaign-service")
	//v1 := g.Group("/v1")
	//
	//v1.GET("/ping", c.Gateway.GatewayV1.PingController.Ping)

	cardGroup := r.Group("/companies/:companyId/cards")
	cardInfoGroup := r.Group("/cards")

	SetupCardRouter(cardGroup, c.Gateway.GatewayV1.CardController)
	SetupCardInfoRouter(cardInfoGroup, c.Gateway.GatewayV1.CardInfoController)
}

func SetupCardRouter(router *gin.RouterGroup, c *controller_v1.CardController) {
	router.GET("", c.FindAll)
	router.POST("", c.Create)
	router.PUT("/:cardId", c.Update)
	router.DELETE("/:cardId", c.Delete)
	router.GET("/import-template-cards", c.ExportCards)
	router.POST("/validate-import-cards", c.ValidateImportCards)
	router.POST("/import-cards", c.ImportCards)
}

func SetupCardInfoRouter(router *gin.RouterGroup, c *controller_v1.CardInfoController) {
	router.GET("/:cardId", c.GetCard)
	router.GET("/vcard-th/:cardId", c.GetVCardTH)
	router.GET("/vcard-en/:cardId", c.GetVCardEN)
}
