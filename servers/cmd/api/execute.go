package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"servers/internal/api"
	"servers/pkg/rds"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func execute(c *cli.Context) error {
	if c.NArg() != 0 {
		if err := cli.ShowAppHelp(c); err != nil {
			return err
		}
		// Return empty error to set 1 to exit status.
		return errors.New("")
	}

	var shutdown = make(chan struct{})
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT)
		<-sigch
		shutdown <- struct{}{}
	}()

	connCtx, connCtxCanncel := context.WithTimeout(context.Background(), 10*time.Second)
	defer connCtxCanncel()
	db, err := rds.ConnectWithContext(
		connCtx,
		c.String(dbUserFlag.Name),
		c.String(dbPasswordFlag.Name),
		c.String(dbHostFlag.Name),
		c.Uint(dbPortFlag.Name),
		c.String(dbNameFlag.Name),
	)
	if err != nil {
		return err
	}

	var logger = logrus.New()
	logger.Out = os.Stdout

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Uint(portFlag.Name)),
		Handler: application.BuildRouter(db),
	}
	go func() {
		if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("ListenAndServe: %v", err)
		}
	}()

	<-shutdown

	logger.Info("Shutting down...")
	ctx, srcancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer srcancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Shutdown server: %v", err)
	}
	logger.Info("Shutdown")
	return nil
}
