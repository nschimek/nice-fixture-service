package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/handler"
	"github.com/nschimek/nice-fixture-service/repository"
	"github.com/nschimek/nice-fixture-service/service"
)

const (
	configFile = "./config/default.yaml"
)

var (
	services *service.ServiceRegistry
	router *gin.Engine
)

func main() {
	core.Log.Info("Starting server...")

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", core.Cfg.Host, core.Cfg.Port),
		Handler: router,
	}

	// Graceful server shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				core.Log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	core.Log.Infof("Server up and listening at %s", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	core.Log.Warnf("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		core.Log.Fatalf("Server forced to shutdown due to error: %v", err)
	}
}

func init() {
	core.SetupViper()
	core.Setup(configFile)

	// setup repositories, services, and handlers: DB->repos->services->handlers (added to router)
	repos := repository.Setup(core.DB)
	services = service.Setup(repos)
	router = handler.CreateRouter(services)
}