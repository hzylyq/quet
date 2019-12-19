package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"quet/routers"
)

func main() {
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())

	s := &http.Server{
		Addr:         ":10000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Println("quet listen server err", err)
			return
		}
	}()

	routers.SetRouter(router)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("shutdown error:", err)
	}
}
