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

func RunAuthServer() {
	server = &http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/callback", callbackHandler)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
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
	OAuthEndpoint := fmt.Sprintf("%s/login", env.API_URL)
	exe.OpenURL(OAuthEndpoint)
	// Start the HTTP server
	RunAuthServer()
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token and user ID from the callback URL parameters
	token := r.URL.Query().Get("token")
	userID := r.URL.Query().Get("userId")
	config.SetUserCfg(userID, token)

	// Respond to the callback request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authenticated Successfully"))
	ShutdownServer()
}
