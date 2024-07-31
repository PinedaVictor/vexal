package pkg

import (
	"context"
	"fmt"
	"vx/config"
	"vx/internal/keys"
	"vx/internal/secrets"

	"github.com/sashabaranov/go-openai"
)

var (
	openAiCtx    context.Context
	openAiClient *openai.Client
)

func initOpenAI() {
	rCfg, _ := keys.KeySrc()
	openAiClient = openai.NewClient(rCfg.Openai_key)
}

// TODO: implement auth init config
func initWithAuth() string {
	user, _ := config.LoadAuth()
	return secrets.GetSecretVersion(fmt.Sprintf("%s_openai", user.UID))
}

func GenerateReponse(prompt string) string {
	initOpenAI()
	openAiCtx = context.Background()
	resp, err := openAiClient.CreateChatCompletion(
		openAiCtx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		errStr := fmt.Sprintf("ChatCompletion error: %v\n", err)
		return errStr
	}
	// fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}
