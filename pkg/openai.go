package pkg

import (
	"context"
	"fmt"
	"vx/config"
	"vx/internal/secrets"

	"github.com/sashabaranov/go-openai"
)

var (
	openAiCtx    context.Context
	openAiClient *openai.Client
)

func initOpenAI() {
	user, _ := config.LoadAuth()
	secret := secrets.GetSecretVersion(fmt.Sprintf("%s_openai", user.UID))
	fmt.Println("openai key:", secret)
	openAiClient = openai.NewClient(secret)
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
