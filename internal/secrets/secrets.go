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
)

var (
	// Vexal Secret Manager context
	secretsCtx context.Context
	// Vexal Secret Manager client instance
	secretClient *secretmanager.Client
	once         sync.Once
	initError    error
)

func initSecretManager() {
	env, _ := config.LoadEnvironment()
	creds := map[string]string{
		"project_id":   env.GCP_PROJECT_ID,
		"type":         env.GCP_ACCOUNT_TYPE,
		"private_key":  env.SECRETS_MAN_PRIVATE_KEY,
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

func AddSecret(secretId string, secretValue string) {
	initSecretManager()
	env, _ := config.LoadEnvironment()
	user, _ := config.LoadAuth()
	secret := fmt.Sprintf("%s_%s", user.UID, secretId)
	parent := fmt.Sprintf("projects/%s/secrets/%s", env.GCP_PROJECT_ID, secret) // "projects/my-project/secrets/my-secret"
	req := &secretmanagerpb.GetSecretRequest{
		Name: parent,
	}
	_, err := secretClient.GetSecret(secretsCtx, req)
	if err != nil {
		secretCreated := createSecet(secretId)
		if secretCreated {
			addSecretVersion(parent, secretValue)
		}
	} else {
		addSecretVersion(parent, secretValue)
	}
	defer secretClient.Close()
}

func createSecet(secretId string) bool {
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
		log.Println("error creating secret")
		return false
	}
	defer secretClient.Close()
	nameResult := strings.Split(result.Name, "/")
	return nameResult[len(nameResult)-1] == secretName
}

func addSecretVersion(secretId string, secretValue string) bool {
	payload := []byte(secretValue)
	fmt.Println(payload)
	req := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secretId,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
			// DataCrc32C: &checksum,
		},
	}
	_, err := secretClient.AddSecretVersion(secretsCtx, req)
	if err != nil {
		log.Println("error adding secret version")
		return false
	}
	defer secretClient.Close()
	return true
}

func GetSecretVersion(secretId string) string {
	initSecretManager()
	env, _ := config.LoadEnvironment()
	parent := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", env.GCP_PROJECT_ID, secretId) // "projects/my-project/secrets/my-secret"
	fmt.Println(parent)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: parent,
	}
	result, err := secretClient.AccessSecretVersion(secretsCtx, req)
	defer secretClient.Close()
	if err != nil {
		fmt.Println("error: secret not found or does not exists", err)
		os.Exit(0)
	}
	return string(result.Payload.Data)

}
