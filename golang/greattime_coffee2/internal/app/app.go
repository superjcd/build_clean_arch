package app

import (
	"fmt"
	v1 "greattime_coffee/internal/controller/http"
	"greattime_coffee/internal/usecase"
	"greattime_coffee/internal/usecase/coupon"
	"greattime_coffee/internal/usecase/repo"
	"greattime_coffee/pkg/httpserver"
	"greattime_coffee/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run() {
	l := logger.New("info")
	coffeeServiceUsecase := usecase.New(
		repo.NewInMemoryRepo(),
		coupon.NewRandomCuppon(),
	)

	handler := gin.New()

	v1.NewRouter(handler, l, coffeeServiceUsecase)
	httpServer := httpserver.NewServer(handler, httpserver.Port("80"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
