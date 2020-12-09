package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serverApp(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	s := &http.Server{
		Addr: addr,
		Handler: mux,
	}
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello "))
	})
	go func() {
		<- ctx.Done()
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func serverDebug(ctx context.Context, addr string) error {
	s := http.Server{
		Addr: addr,
		Handler: http.DefaultServeMux,
	}
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func serverSignal(ctx context.Context, ch chan os.Signal) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-ch:
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				return errors.New(s.String())
			case syscall.SIGHUP:
			default:
				return errors.New("undefined signal")
			}
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return serverApp(ctx, "127.0.0.1:8000")
	})
	g.Go(func() error {
		return serverDebug(ctx, "127.0.0.1:8081")
	})
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	g.Go(func() error {
		return serverSignal(ctx, ch)
	})
	if err:=g.Wait(); err != nil {
		log.Fatal(err)
	}
}
