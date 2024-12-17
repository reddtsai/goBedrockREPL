package bedrock

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	bt "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

type IClaude3Sonnet interface {
}

type Claude3Sonnet struct {
	ModelID          string
	AnthropicVersion string
	MaxTokens        int32
	Client           *bedrockruntime.Client
}

func NewClaude3Sonnet() *Claude3Sonnet {

	return &Claude3Sonnet{
		ModelID:          "apac.anthropic.claude-3-sonnet-20240229-v1:0",
		AnthropicVersion: "bedrock-2023-05-31",
		MaxTokens:        1024,
		Client:           bedrockruntime.NewFromConfig(awsCfg),
	}
}

func (c *Claude3Sonnet) InvokeModel(ctx context.Context, msg string) string {
	payload := InvokeModelRequest{
		AnthropicVersion: c.AnthropicVersion,
		MaxTokens:        1024,
		SystemPrompt:     WEATHER_SYSTEM_PROMPT,
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: msg,
					},
				},
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	output, err := c.Client.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		Body:        payloadBytes,
		ModelId:     aws.String(c.ModelID),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		log.Fatal(err)
	}
	var resp InvokeModelResponse
	err = json.Unmarshal(output.Body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	return resp.ResponseContent[0].Text
}

func (c *Claude3Sonnet) Converse(ctx context.Context, msg string) string {
	var sysPrompt []bt.SystemContentBlock
	sysPrompt = append(sysPrompt, &bt.SystemContentBlockMemberText{Value: WEATHER_SYSTEM_PROMPT})
	input := &bedrockruntime.ConverseInput{
		ModelId: aws.String(c.ModelID),
		System:  sysPrompt,
		InferenceConfig: &bt.InferenceConfiguration{
			MaxTokens: aws.Int32(c.MaxTokens),
		},
	}
	message := bt.Message{
		Role: bt.ConversationRoleUser,
		Content: []bt.ContentBlock{
			&bt.ContentBlockMemberText{
				Value: msg,
			},
		},
	}
	input.Messages = append(input.Messages, message)
	resp, err := c.Client.Converse(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := resp.Output.(*bt.ConverseOutputMemberMessage)
	responseContentBlock := result.Value.Content[0]
	text, _ := responseContentBlock.(*bt.ContentBlockMemberText)

	return text.Value
}
