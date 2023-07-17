package server

import (
	"card-service/internal/config"
	"card-service/internal/util/log"
	"fmt"
	"net"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Server ...
type Server struct {
	Gateway GRPCGateway

	RouteRestful RouteRestfulService
	config       *config.Configuration
	logger       *log.Logger
	Server       *grpc.Server
	Listener     net.Listener
}

////Start ...
//func (gs *Server) Start() error {
//	gs.configRoute()
//
//	gs.logger.Println("Server listening on GRPC port", gs.config.Port)
//
//	return gs.Server.Serve(gs.Listener)
//}

// StartRestful ...
func (gs *Server) Start() error {
	//TODO: New Gin Service
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With", "Content-Disposition"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gs.ConfigRouteRESTful(app)
	gs.logger.Println("Server listening on RESTful port", gs.config.PortRestful)

	err := app.Run(":" + gs.config.PortRestful)
	if err != nil {
		gs.logger.Println("Server Run on RESTful error: ", err)
		return err
	}

	return err
}

// NewServer ...
func NewServer(
	g GRPCGateway,
	restful RouteRestfulService,
	c *config.Configuration,
	l *log.Logger,
) *Server {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Port))
	if err != nil {
		panic(err)
	}
	return &Server{
		Gateway:      g,
		RouteRestful: restful,
		config:       c,
		logger:       l,
		Server:       grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*64), grpc.MaxSendMsgSize(1024*1024*64)),
		Listener:     listen,
	}
}
