package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"payment/auth"
	"payment/config"
	frameController "payment/frame/controller"
	frameRepository "payment/frame/repository"
	frameUsecase "payment/frame/usecase"
	categoryRepository "payment/frame_category/repository"
	categoryUsecase "payment/frame_category/usecase"
	paymentUsecase "payment/payments/usecase"
	promoController "payment/promo/controller"
	promoRepository "payment/promo/repository"
	promoUsecase "payment/promo/usecase"
	transactionController "payment/transactions/controller"
	transactionRepo "payment/transactions/repository"
	transactionUsecase "payment/transactions/usecase"
	userController "payment/users/controller"
	userRepository "payment/users/repository"
	userUsecase "payment/users/usecase"

	webHandler "payment/web/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	httpServer *http.Server
	cfg        *config.MainConfig
	db         *gorm.DB
}

func NewServer(cfg *config.MainConfig, db *gorm.DB) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Run() error {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(CORSMiddleware())
	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("startup", cookieStore))
	router.HTMLRender = loadTemplates(isWebAdminRunOnServer(s.cfg.Server.Env, "web/templates"))
	router.Static("/images", isWebAdminRunOnServer(s.cfg.Server.Env, "images"))
	router.Static("/css", isWebAdminRunOnServer(s.cfg.Server.Env, "web/assets/css"))
	router.Static("/js", isWebAdminRunOnServer(s.cfg.Server.Env, "web/assets/js"))
	router.Static("/webfonts", isWebAdminRunOnServer(s.cfg.Server.Env, "web/assets/webfonts"))
	// Initialize repository
	transactionRepo := transactionRepo.NewTransactionRepository(s.db)
	userRepository := userRepository.NewUserRepository(s.db)
	promoRepository := promoRepository.NewPromoRepository(s.db)
	categoryRepository := categoryRepository.NewCategoryRepository(s.db)
	frameRepository := frameRepository.NewFrameRepository(s.db)
	// initialize usecase
	authService := auth.NewService()
	paymentUC := paymentUsecase.NewPaymentMidtrans(s.cfg.Midtrans.ClientKey, s.cfg.Midtrans.ServerKey, s.cfg.Midtrans.APIEnv)
	transactionUC := transactionUsecase.NewTransactionUsecase(transactionRepo, paymentUC)
	userUC := userUsecase.NewUserUsecase(userRepository)
	promoUC := promoUsecase.NewPromoUsecase(promoRepository)
	categoryUC := categoryUsecase.NewCategoryUsecase(categoryRepository)
	frameUC := frameUsecase.NewFrameUsecase(frameRepository, categoryRepository, s.cfg.Product.BaseUrl)
	// initialize router
	transactionController.RegisterHTTPEndpoints(router, transactionUC)
	userController.RegisterHTTPEndpoints(router, userUC, authService)
	promoController.RegisterHTTPEndpoints(router, promoUC)
	frameController.RegisterHTTPEndpoints(router, frameUC, categoryUC)
	// quick fix using middleware, initialize transaction Controller
	transactionController := transactionController.NewTransactionControllers(transactionUC)
	promoController := promoController.NewPromoController(promoUC)
	api := router.Group("api/v1")
	api.POST("/transaction/bypass", authMiddleware(authService, userUC, s.cfg.Product.LocalEvent), transactionController.BypassNormalFlow)
	api.POST("/promo", authMiddleware(authService, userUC, s.cfg.Product.AdminSwitch), promoController.CreatePromoCode)

	// Web Handler
	promoWebHandler := webHandler.NewPromoHandler(promoUC, userUC)
	sessionHandler := webHandler.NewSessionHandler(userUC)
	categoryWebHandler := webHandler.NewCategoryHandler(categoryUC, userUC)
	frameWebHandler := webHandler.NewFrameHandler(frameUC, userUC, categoryUC)
	// Dashboard Promo Router
	router.GET("/login", sessionHandler.New)
	router.POST("/session", sessionHandler.Create)
	router.GET("/logout", sessionHandler.Destroy)
	router.GET("/promo", authAdminMiddleware(), promoWebHandler.Index)
	router.GET("/promo/new", authAdminMiddleware(), promoWebHandler.New)
	router.POST("/promo", authAdminMiddleware(), promoWebHandler.Create)
	router.GET("/promo/status/:id", authAdminMiddleware(), promoWebHandler.ActivatePage)
	router.POST("/promo/update/status/:id", authAdminMiddleware(), promoWebHandler.ActivationAction)
	router.GET("/promo/edit/:id", authAdminMiddleware(), promoWebHandler.Edit)
	router.POST("/promo/update/:id", authAdminMiddleware(), promoWebHandler.Update)
	router.GET("/category", authAdminMiddleware(), categoryWebHandler.Index)
	router.GET("/category/new", authAdminMiddleware(), categoryWebHandler.New)
	router.POST("/category", authAdminMiddleware(), categoryWebHandler.Create)
	router.POST("/category/update/:id", authAdminMiddleware(), categoryWebHandler.Update)
	router.GET("/category/edit/:id", authAdminMiddleware(), categoryWebHandler.Edit)
	router.GET("/category/delete/:id", authAdminMiddleware(), categoryWebHandler.DeleteRoute)
	router.POST("/category/delete/:id", authAdminMiddleware(), categoryWebHandler.Delete)
	router.GET("/frame", authAdminMiddleware(), frameWebHandler.Index)
	router.GET("/frame/new", authAdminMiddleware(), frameWebHandler.New)
	router.POST("/frame", authAdminMiddleware(), frameWebHandler.Create)
	router.GET("/frame/delete/:id", authAdminMiddleware(), frameWebHandler.DeleteRoute)
	router.POST("/frame/delete/:id", authAdminMiddleware(), frameWebHandler.Delete)
	router.GET("/frame/change/:id", authAdminMiddleware(), frameWebHandler.ChangeStatusRoute)
	router.POST("/frame/change/:id", authAdminMiddleware(), frameWebHandler.ChangeStatus)

	address := s.cfg.Server.Port
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Listening and serving POST service HTTP on localhost", address)
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}

func isWebAdminRunOnServer(env string, pathWebAdmin string) string {
	url := pathWebAdmin
	if env != "local" {
		url = "/app/" + pathWebAdmin
	}

	return url
}
