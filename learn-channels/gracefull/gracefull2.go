package gracefull

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/mazxcv/custom-tcp-server/gracefull/closer"
)

const (
	listenAddr      = "127.0.0.1:3000"
	shutdownTimeout = 5 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		log.Fatal(err)
	}

}

func gracefull(ctx context.Context) error {
	var (
		mux = http.NewServeMux()
		srv = &http.Server{
			Addr:    listenAddr,
			Handler: mux,
		}
		c = &closer.Closer{}
	)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}))

	c.Add(func(ctx context.Context) error {
		time.Sleep(2 * time.Second)
		return nil
	})

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

	if err := c.Close(shutdownCtx); err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}

func stop() {}
