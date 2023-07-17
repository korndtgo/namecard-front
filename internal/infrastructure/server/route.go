package server

import (
	controller_v1 "card-service/internal/controller/v1"

	"go.uber.org/dig"
)

// GRPCGateway ...
type GRPCGateway struct {
	dig.In

	GatewayV1 controller_v1.Gateway
}

//func (gs *Server) configRoute(r *gin.Engine) {
//	authGroup := r.Group("/auth")
//	cardGroup := r.Group("/companies/:COMPANY_ID/cards")
//	cardInfoGroup := r.Group("/cards")
//	companyGroup := r.Group("/companies")
//	logsGroup := r.Group("/logs")
//	userGroup := r.Group("/users")
//
//	SetupCompanyRouter(companyGroup)
//	SetupCardRouter(cardGroup)
//	SetupCardInfoRouter(cardInfoGroup)
//	SetupLogsRouter(logsGroup)
//	SetupUserRouter(userGroup)
//	SetupAuthRouter(authGroup)
//}
//
//func SetupAuthRouter(router *gin.RouterGroup) {
//	router.GET("/profile")
//	router.PUT("/change-password")
//	router.PUT("/reset-password")
//}
//
//func SetupCompanyRouter(router *gin.RouterGroup) {
//	router.GET("")
//	router.POST("")
//	router.PUT("/:COMPANY_ID")
//	router.DELETE("/:COMPANY_ID")
//}
//
//func SetupCardRouter(router *gin.RouterGroup) {
//	router.GET("")
//	router.POST("")
//	router.PUT("/:CARD_ID")
//	router.DELETE("/:CARD_ID")
//	router.GET("/import-template-cards")
//	router.POST("/validate-import-cards")
//	router.POST("/import-cards")
//}
//
//func SetupCardInfoRouter(router *gin.RouterGroup) {
//	router.GET("/:CARD_UUID")
//	router.GET("/vcard-th/:CARD_UUID")
//	router.GET("/vcard-en/:CARD_UUID")
//}
//
//func SetupLogsRouter(router *gin.RouterGroup) {
//	router.GET("")
//}
//
//func SetupUserRouter(router *gin.RouterGroup) {
//	router.GET("")
//	router.POST("")
//	router.PUT("/:UUID")
//	router.DELETE("/:UUID")
//	router.GET("/roles")
//	router.GET("/profile")
//	router.GET("/import-template-users")
//	router.POST("/validate-import-users")
//	router.POST("/import-users")
//}
