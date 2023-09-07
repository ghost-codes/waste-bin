package server

// import (
// 	"fmt"
// 	"ghost-codes/waste-bin/config"
// 	"ghost-codes/waste-bin/db"
// 	models "ghost-codes/waste-bin/db/models"
// 	"ghost-codes/waste-bin/middleware"

// 	"firebase.google.com/go/auth"
// 	"github.com/gin-gonic/gin"
// 	"github.com/kardianos/service"
// )

// type Server struct {
// 	router       *gin.Engine
// 	config       config.Config
// 	store        *models.Store
// 	firebaseAuth *auth.Client
// }

// func NewServer(config config.Config, firebaseAuth *auth.Client) (*Server, error) {
// 	dbIns, err := db.NewGorm(config.DBSource())
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to initialize db,%s", err)
// 	}
// 	server := Server{
// 		config:       config,
// 		firebaseAuth: firebaseAuth,
// 		store: &models.Store{
// 			DB: dbIns,
// 		},
// 	}

// 	server.setupRouter()

// 	return &server, nil
// }

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	router.Use(middleware.Errors("./config/.env", "rollbackToken", service.ConsoleLogger))
// 	router.Use(middleware.UserAuthMiddleware(*server.store, server.firebaseAuth))

// 	// task repository
// 	// taskRepo := todoRepository{}
// 	// router.GET("/todoclear
// 	server.router = router
// }

// func (server *Server) Start(addr string) error {
// 	fmt.Println("======= Server has Started ==============")
// 	return server.router.Run(addr)

// }

// type errorJson struct {
// 	Message string `json:"message"`
// }

// func NewErrorJson(err error) gin.H {
// 	return gin.H{
// 		"message": err.Error(),
// 	}
// }

// func GenericSuccesMsg(msg string) gin.H {
// 	return gin.H{
// 		"message": msg,
// 	}
// }
