package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// этот контекст будет передаваться во все процессы
	_, cancel := context.WithCancel(context.Background())
	mainSrv := http.Server{Addr: ":5000"}

	sigquit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigquit, syscall.SIGINT, syscall.SIGTERM)
	stopAppCh := make(chan struct{})
	go func() {
		log.Println("Captured signal: ", <-sigquit)
		log.Println("Gracefully shutting down server...")

		cancel()

		if err := mainSrv.Shutdown(context.Background()); err != nil {
			log.Println("Can't shutdown main server: ", err.Error())
		}
		stopAppCh <- struct{}{}
	}()

	<-stopAppCh
}
