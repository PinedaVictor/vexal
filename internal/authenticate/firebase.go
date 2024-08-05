package authenticate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"vx/config"

	firebase "firebase.google.com/go"
	"github.com/fatih/color"
	"google.golang.org/api/option"
)

var (
	// Vexal Firebase context - use this it access firebase functions
	firebaseCtx context.Context
	// Vexal Firebase app
	firebaseApp *firebase.App
	once        sync.Once
)

func initFirebase() {
	env, _ := config.LoadEnvironment()
	creds := map[string]string{
		"type":         env.GCP_ACCOUNT_TYPE,
		"project_id":   env.GCP_PROJECT_ID,
		"private_key":  strings.Replace(env.FB_ADMIN_PRIVATE_KEY, "\\n", "\n", -1),
		"client_email": env.FB_ADMIN_CLIENT_EMAIL,
	}
	credsJSON, err := json.Marshal(creds)
	if err != nil {
		log.Fatalf("error marshaling credentials: %v", err)
	}
	firebaseCtx = context.Background()
	cfg := &firebase.Config{
		ProjectID: env.GCP_PROJECT_ID,
	}
	opt := option.WithCredentialsJSON(credsJSON)
	once.Do(func() {
		firebaseApp, err = firebase.NewApp(firebaseCtx, cfg, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	})
}

func ValidateToken() (bool, string) {
	initFirebase()
	user, _ := config.LoadAuth()
	auth, err := firebaseApp.Auth(firebaseCtx)
	if err != nil {
		log.Println("error initializing Firebase", err)
		return false, "error"
	}
	validToken, tokenErr := auth.VerifyIDToken(firebaseCtx, user.Token)
	if tokenErr != nil {
		return false, "Unauthorized: Run vx login"
	}
	return validToken.UID == user.UID, "Authorized"
}

// TODO: Implement refresh token flow
func RequireAuth() {
	initFirebase()
	auth, _ := ValidateToken()
	authMsg := "The command you're tyring to run either requires authentication or repository configuration."
	if !auth {
		fmt.Println(authMsg)
		os.Exit(0)
	}
}

func RootAuthStatus() string {
	initFirebase()
	auth, msg := ValidateToken()
	authMsg := ""
	if auth {
		c := color.New(color.FgGreen)
		authMsg = c.Sprint(msg)
		return authMsg
	}
	red := color.New(color.FgRed)
	authMsg = red.Sprint(msg)
	return authMsg
}
