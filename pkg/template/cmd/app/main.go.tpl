package main

import (
	"context"
	goflag "flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/pkg/ctrl"
	"{{ .ProjectName }}/pkg/logi"
	flag "github.com/spf13/pflag"
)

var log = logi.Log.Sugar()

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	development := flag.Bool("development", true, "development mod")
	flag.Parse()

	if *development {
		logi.SetLogger(logi.Logger(true))
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	pprof.Register(router) // default is "debug/pprof"
	ctrl.Register(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Infof("server start: %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown:", err)
		os.Exit(1)
	}
	log.Info("server exiting")
}
