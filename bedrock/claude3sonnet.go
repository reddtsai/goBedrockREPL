package bedrock

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type IClaude3Sonnet interface {
}

type Claude3Request struct {
	AnthropicVersion string    `json:"anthropic_version"`
	MaxTokens        int       `json:"max_tokens"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature,omitempty"`
	TopP             float64   `json:"top_p,omitempty"`
	TopK             int       `json:"top_k,omitempty"`
	StopSequences    []string  `json:"stop_sequences,omitempty"`
	SystemPrompt     string    `json:"system,omitempty"`
}

type Content struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Message struct {
	Role    string    `json:"role,omitempty"`
	Content []Content `json:"content,omitempty"`
}

type Claude3Response struct {
	ID              string            `json:"id,omitempty"`
	Model           string            `json:"model,omitempty"`
	Type            string            `json:"type,omitempty"`
	Role            string            `json:"role,omitempty"`
	ResponseContent []ResponseContent `json:"content,omitempty"`
	StopReason      string            `json:"stop_reason,omitempty"`
	StopSequence    string            `json:"stop_sequence,omitempty"`
	Usage           Usage             `json:"usage,omitempty"`
}

type ResponseContent struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens,omitempty"`
	OutputTokens int `json:"output_tokens,omitempty"`
}

type Claude3Sonnet struct {
	ModelID          string
	AnthropicVersion string
	MaxTokens        int
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

func (c *Claude3Sonnet) Claude3SonnetInvokeModel(ctx context.Context, msg string) string {
	payload := Claude3Request{
		AnthropicVersion: c.AnthropicVersion,
		MaxTokens:        c.MaxTokens,
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
	var resp Claude3Response
	err = json.Unmarshal(output.Body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	return resp.ResponseContent[0].Text
}
