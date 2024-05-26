package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

func initSecretManger() {
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
	log.Println("Passing creds json")
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

func CreateSecet() {
	initSecretManger()
	env, _ := config.LoadEnvironment()
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", env.GCP_PROJECT_ID),
		SecretId: "129873287UID_SDKSUFFIX",
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
				"sdk": "sdk",
			},
		},
	}

	result, err := secretClient.CreateSecret(secretsCtx, req)
	if err != nil {
		log.Println("error creating secret")
	}
	log.Println(result.Name)
	defer secretClient.Close()
}
