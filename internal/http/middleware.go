package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
)

func MiddlewareRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				stack := make([]byte, 1<<16)
				length := runtime.Stack(stack, true)

				// Convert the stack trace to a string
				stackStr := string(stack[:length])

				// Log the error and the stack trace
				slog.Log(context.Background(), slog.LevelError, "recovered from panic", slog.String("error", fmt.Sprintf("%v", r)), slog.String("stack", stackStr))
				fmt.Println(fmt.Sprintf("%v", stackStr))

				// Return internal server error
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func MiddlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allow these HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allow these headers

		// Handle preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
