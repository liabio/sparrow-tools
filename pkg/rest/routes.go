package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var router *gin.Engine
var timeoutSimulate time.Duration

type RunOptions struct {
	ListenPort string
	InterPath  string
	Timeout    time.Duration
}

func init() {
	router = gin.New()
	router.Use(gin.Recovery())
}

func getRoutes(opt *RunOptions) {
	router.POST(opt.InterPath, timeout)
}

func Run(opt *RunOptions) {
	timeoutSimulate = opt.Timeout
	getRoutes(opt)
	srv := &http.Server{
		Addr:         opt.ListenPort,
		Handler:      router,
		ReadTimeout:  25 * time.Second,
		//WriteTimeout: 25 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		klog.Infof("starting HTTP server listening at port: %v", opt.ListenPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatalf("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	klog.Info("Shutting down web server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		klog.Fatal("Server forced to shutdown:", err)
	}
	klog.Info("Server exiting")
}

func handleResponse(c *gin.Context, data interface{}, err error) {
	resp := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func timeout(c *gin.Context) {
	start := time.Now()
	time.Sleep(timeoutSimulate)
	klog.Infof("time-spend %.2f seconds", time.Since(start).Seconds())
	handleResponse(c, nil, nil)
}
