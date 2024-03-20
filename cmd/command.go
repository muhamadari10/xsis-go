package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	env "movie-app/config"
	httpv1 "movie-app/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CLI interface {
	Start()
	Error() error
}

type Command struct {
	di   *env.Dependency
	args []string
	err  error
}

func NewCLI(params *env.Dependency, args []string) *Command {
	return &Command{params, args, nil}
}

func (cmd *Command) Start() {
	if cmd.validateCmd(); cmd.err != nil {
		return
	}
	switch cmd.args[1] {
	case "movie":
		go cmd.runApiGin()
		cmd.gracefulShutdownAll()
		return
	default:
		flag.PrintDefaults()
		return
	}
}

func (cmd *Command) validateCmd() {
	// Can be modified to get how many commands we want as a default
	if len(cmd.args) < 2 {
		cmd.err = errors.New("--help for options")
		return
	}
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <additional>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.String("point-gateway", "", "./main start")
	flag.Parse()
}

func (cmd *Command) runApiGin() {
	serveRoutes := &http.Server{
		Addr:           ":" + cmd.di.Params.Ports.Gin,
		Handler:        httpv1.InitRoutesGin(cmd.di),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("listen: %s\n", cmd.di.Params.Ports.Gin)
	go func() {
		if err := serveRoutes.ListenAndServe(); cmd.err != nil && errors.Is(cmd.err, http.ErrServerClosed) {
			cmd.err = err
			log.Printf("listen: %s\n", cmd.err)
			return
		}
	}()
	cmd.gracefulShutdown(serveRoutes)
	return
}

func (cmd *Command) gracefulShutdownAll() {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func (cmd *Command) gracefulShutdown(serverRoutes *http.Server) {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serverRoutes.Shutdown(ctx); err != nil {
		cmd.err = err
		return
	}
}

func (cmd *Command) Error() error {
	return cmd.err
}
