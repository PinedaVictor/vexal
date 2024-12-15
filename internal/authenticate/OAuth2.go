// Package authenticate handles any authentication flows needed to secure user data
package authenticate

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"vx/config"
	"vx/pkg/exe"
)

var server *http.Server

// Key for storing the config type in context
type contextKey string

const configTypeKey contextKey = "configType"

func RunAuthServer(configType string) {
	env, _ := config.LoadEnvironment()
	server = &http.Server{Addr: env.SERVER_REDIRECT_ADDR}
	http.Handle("/callback", withConfigType(configType, http.HandlerFunc(callbackHandler)))

	err := server.ListenAndServe()
	if err != nil {
		// NOTE: Theres not an issue when authenticating into jira but this will print an error
		fmt.Println("", err)
	}
}

func ShutdownServer() {
	// Shutdown the server after handling the callback
	go func() {
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shut down server: %v", err)
		}
		os.Exit(0)
	}()
}

func Login() {
	env, _ := config.LoadEnvironment()
	OAuthEndpoint := fmt.Sprintf("%s/login", env.APP_URL)
	exe.OpenURL(OAuthEndpoint)
	// Start the HTTP server
	RunAuthServer("vx-user")
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token and user ID from the callback URL parameters
	token := r.URL.Query().Get("token")
	userID := r.URL.Query().Get("userId")

	// Retrieve the config type from context
	configType := r.Context().Value(configTypeKey).(string)
	// Dynamically update the appropriate config based on configType
	// FIXME: This works but the strings must match in calling function
	switch configType {
	case "vx-user":
		config.SetUserCfg(userID, token)
	case "jira":
		config.SetJiraCfg(userID, token)
	default:
		http.Error(w, "Invalid configuration type", http.StatusBadRequest)
		return
	}

	// Respond to the callback request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authenticated Successfully"))
	ShutdownServer()
}

// Middleware to inject the config type into the context
func withConfigType(configType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// func WithValue(parent Context, key, val any) Context
		ctx := context.WithValue(r.Context(), configTypeKey, configType)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
