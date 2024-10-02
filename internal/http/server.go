package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Middleware func(http.Handler) http.Handler

type ApiServer struct {
	server     *http.Server
	mux        *http.ServeMux
	middleware []Middleware
}

func NewApiServer(port int) *ApiServer {
	mux := http.NewServeMux()
	return &ApiServer{
		mux:        mux,
		middleware: []Middleware{},
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (a *ApiServer) Use(mw Middleware) {
	a.middleware = append(a.middleware, mw)
}

func (a *ApiServer) applyMiddleware(h http.Handler) http.Handler {
	for i := len(a.middleware) - 1; i >= 0; i-- {
		h = a.middleware[i](h)
	}
	return h
}

func (a *ApiServer) AddHandler(path string, hFunc func(w http.ResponseWriter, r *http.Request)) {
	finalHandler := a.applyMiddleware(http.HandlerFunc(hFunc))
	a.mux.Handle(path, finalHandler)
}

func (a *ApiServer) AddStaticHandler(urlPath string, dirPath string) {
	// FileServer to serve static files
	fileServer := http.FileServer(http.Dir(dirPath))

	// Custom handler to serve .wasm files with correct Content-Type
	a.mux.Handle(urlPath, http.StripPrefix(urlPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all static file requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if the request is for a .wasm file
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		fileServer.ServeHTTP(w, r)
	})))
}

func (a *ApiServer) ListenAndServe() {

	// Channel to listen for interrupt signals (e.g., Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		slog.Log(context.Background(), slog.LevelInfo, "Server is running on port 8001...")
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Log(context.Background(), slog.LevelError, "Could not listen on :8001")
			os.Exit(1)
		}
	}()

	// Block until we receive a signal
	<-stop
	slog.Log(context.Background(), slog.LevelInfo, "Shutting down server...")

	// Create a deadline to wait for the server to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Attempt to gracefully shut down the server
	if err := a.server.Shutdown(ctx); err != nil {
		slog.Log(ctx, slog.LevelError, "server forced to shutdown")
		os.Exit(1)
	}

	slog.Log(ctx, slog.LevelInfo, "Server gracefully stopped")
}
