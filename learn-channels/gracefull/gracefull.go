package gracefull

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// const (
// 	listenAddr      = "127.0.0.1:3000"
// 	shutdownTimeout = 5 * time.Second
// )

// func gracefull() {
// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	defer stop()

// 	if err := runServer(ctx); err != nil {
// 		log.Fatal(err)
// 	}

// }

func runServer(ctx context.Context) error {
	var (
		mux = http.NewServeMux()
		srv = &http.Server{
			Addr:    listenAddr,
			Handler: mux,
		}
	)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server started on %s\n", listenAddr)
	<-ctx.Done()
	fmt.Println("gracefullyshutting down the server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}

	longShutdown := make(chan struct{}, 1)

	go func() {
		time.Sleep(10 * time.Second)
		longShutdown <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("shutdown timed out: %w", shutdownCtx.Err())
	case <-longShutdown:
		log.Println("long shutdown")
	}

	return nil
}

// func stop() {}
