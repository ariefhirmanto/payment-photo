package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"payment/config"
	paymentUsecase "payment/payments/usecase"
	transactionController "payment/transactions/controller"
	transactionRepo "payment/transactions/repository"
	transactionUsecase "payment/transactions/usecase"

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

	// Initialize repository
	transactionRepo := transactionRepo.NewTransactionRepository(s.db)
	// initialize usecase
	paymentUC := paymentUsecase.NewPaymentMidtrans(s.cfg.Midtrans.ClientKey, s.cfg.Midtrans.ServerKey, s.cfg.Midtrans.APIEnv)
	transactionUC := transactionUsecase.NewTransactionUsecase(transactionRepo, paymentUC)
	// initialize router
	transactionController.RegisterHTTPEndpoints(router, transactionUC)

	address := ":9000"
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Listening and serving POST service HTTP on localhost:", address)
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
