package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"vx/config"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Vexal Secret Manager context
	secretsCtx context.Context
	// Vexal Secret Manager client instance
	secretClient *secretmanager.Client
	once         sync.Once
	initError    error
)

// strings.Replace(env.FB_ADMIN_PRIVATE_KEY, "\\n", "\n", -1))
func initSecretManager() {
	env, _ := config.LoadEnvironment()
	creds := map[string]string{
		"project_id":   env.GCP_PROJECT_ID,
		"type":         env.GCP_ACCOUNT_TYPE,
		"private_key":  strings.Replace(env.SECRETS_MAN_PRIVATE_KEY, "\\n", "\n", -1),
		"client_email": env.SECRETS_MAN_EMAIL,
	}
	credsJSON, err := json.Marshal(creds)
	if err != nil {
		log.Fatalf("error marshaling credentials: %v", err)
	}
	secretsCtx = context.Background()
	opt := option.WithCredentialsJSON(credsJSON)
	once.Do(func() {
		secretClient, err = secretmanager.NewClient(secretsCtx, opt)
		initError = err
		if err != nil {
			log.Fatalf("failed to create secret manager client: %v", err)
		}
	})
}

func AddSecret(secretId string, secretValue string) bool {
	initSecretManager()
	defer secretClient.Close()
	env, _ := config.LoadEnvironment()
	user, _ := config.LoadAuth()
	secret := fmt.Sprintf("%s_%s", user.UID, secretId)
	parent := fmt.Sprintf("projects/%s/secrets/%s", env.GCP_PROJECT_ID, secret) // "projects/my-project/secrets/my-secret"
	return addSecretVersion(parent, secretValue)
}

func CreateSecret(secretId string) bool {
	initSecretManager()
	env, _ := config.LoadEnvironment()
	user, _ := config.LoadAuth()
	secretName := fmt.Sprintf("%s_%s", user.UID, secretId) // "UID_SDKSUFFIX"
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", env.GCP_PROJECT_ID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_UserManaged_{
					UserManaged: &secretmanagerpb.Replication_UserManaged{
						Replicas: []*secretmanagerpb.Replication_UserManaged_Replica{
							{
								Location: "us-central1",
							},
						},
					},
				},
			},
			Labels: map[string]string{
				"sdk": secretId,
			},
		},
	}
	result, err := secretClient.CreateSecret(secretsCtx, req)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			fmt.Println("Secret version already exists.")
			fmt.Println("Try running: vx config set -k [API] -v [YOUR API KEY]")
			return false
		}
		return false
	}
	defer secretClient.Close()
	nameResult := strings.Split(result.Name, "/")
	return nameResult[len(nameResult)-1] == secretName
}

func addSecretVersion(secretId string, secretValue string) bool {
	payload := []byte(secretValue)
	req := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secretId,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
			// DataCrc32C: &checksum,
		},
	}
	_, err := secretClient.AddSecretVersion(secretsCtx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			fmt.Println("API key not found or does not exist.")
		} else if status.Code(err) == codes.Canceled {
			fmt.Println("Request canceled. Client connection might be closing.")
		} else {
			fmt.Printf("An unexpected error occurred")
		}
		return false
	}
	defer secretClient.Close()
	return true
}

func GetSecretVersion(secretId string) string {
	initSecretManager()
	env, _ := config.LoadEnvironment()
	// "projects/my-project/secrets/my-secret"
	parent := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", env.GCP_PROJECT_ID, secretId)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: parent,
	}
	result, err := secretClient.AccessSecretVersion(secretsCtx, req)
	/*
		TODO: Use or don't use defer statement
		The defer statement is used to close the secret manager client. The issue arises when
		there are 2 calls from the calling function.
		FIXME: You can either think of a different way of managing the client or accept an array od secretIds
	*/
	// defer secretClient.Close()
	if err != nil {
		fmt.Println("error: secret not found or does not exists.")
		fmt.Println("Make sure the API you are trying to use is enabled and configured.")
		os.Exit(0)
	}
	if err != nil {
		fmt.Printf("Error accessing secret version")
		if status.Code(err) == codes.NotFound {
			fmt.Println("Secret not found or does not exist.")
		} else if status.Code(err) == codes.Canceled {
			fmt.Println("Request canceled. gRPC client connection might be closing.")
		} else {
			fmt.Printf("An unexpected error occurred")
		}
		os.Exit(1)
	}
	return string(result.Payload.Data)
}
