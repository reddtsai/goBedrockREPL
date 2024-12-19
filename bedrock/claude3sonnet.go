package bedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/document"
	bt "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

var (
	tool = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"location": map[string]interface{}{
				"type":        "string",
				"description": "The city and state, e.g. San Francisco, CA",
			},
			"unit": map[string]interface{}{
				"type":        "string",
				"enum":        []string{"celsius", "fahrenheit"},
				"description": "The unit of temperature, either 'celsius' or 'fahrenheit'",
			},
		},
	}

	stop_reason = "tool_use"
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
		// InferenceConfig: &bt.InferenceConfiguration{
		// 	MaxTokens: aws.Int32(c.MaxTokens),
		// },
	}

	message := bt.Message{
		Role: bt.ConversationRoleUser,
		Content: []bt.ContentBlock{
			&bt.ContentBlockMemberText{
				Value: msg,
			},
			// &bt.ContentBlockMemberText{
			// 	Value: `
			// 	You should answer based on the user's question, generate a JSON object with the following structure:
			// 	{
			// 		"Content": [
			// 		{
			// 			"Value": "<Provide a concise answer to the user's question here>",
			// 			"Temperatures": "<temperatures in °C (°F)>",
			// 			"Wind": {
			// 				"Speed":"<wind speed in km/h (mph)>",
			// 				"Desc":"<wind direction>"
			// 			},
			// 			"Conditions": "<Keep weather reports concise. Sparingly use emojis where appropriate.>"
			// 		}]
			// 	}`,
			// },
		},
	}
	input.Messages = append(input.Messages, message)
	resp, err := c.Client.Converse(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	// jsonResp, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(jsonResp))

	var txt string
	result := resp.Output.(*bt.ConverseOutputMemberMessage).Value
	if len(result.Content) > 0 {
		if resp.StopReason == "tool_use" {
			for _, tc := range result.Content {
				if _, ok := tc.(*bt.ContentBlockMemberToolUse); !ok {
					continue
				}
				var toolUse bt.ToolUseBlock = tc.(*bt.ContentBlockMemberToolUse).Value
				txt = *toolUse.Name
				break
			}
		} else {
			responseContentBlock := result.Content[0]
			text, _ := responseContentBlock.(*bt.ContentBlockMemberText)
			txt = text.Value
		}
	}

	return txt
}

func (c *Claude3Sonnet) ConverseWithTool(ctx context.Context, msg string) string {
	var sysPrompt []bt.SystemContentBlock
	sysPrompt = append(sysPrompt, &bt.SystemContentBlockMemberText{Value: WEATHER_SYSTEM_PROMPT})
	// var toolSpecs []bt.ToolSpecification
	toolSpec := bt.ToolSpecification{
		Name:        aws.String("get_weather"),
		Description: aws.String("Get the current weather in a given location"),
		InputSchema: &bt.ToolInputSchemaMemberJson{
			Value: document.NewLazyDocument(tool),
		},
	}
	var tools []bt.Tool
	tools = append(tools, &bt.ToolMemberToolSpec{
		Value: toolSpec,
	})
	toolConfig := &bt.ToolConfiguration{
		Tools: tools,
	}
	input := &bedrockruntime.ConverseInput{
		ModelId: aws.String(c.ModelID),
		System:  sysPrompt,
		// InferenceConfig: &bt.InferenceConfiguration{
		// 	MaxTokens: aws.Int32(c.MaxTokens),
		// },
		ToolConfig: toolConfig,
	}

	message := bt.Message{
		Role: bt.ConversationRoleUser,
		Content: []bt.ContentBlock{
			&bt.ContentBlockMemberText{
				Value: msg,
			},
			// &bt.ContentBlockMemberText{
			// 	Value: `
			// 	You should answer based on the user's question, generate a JSON object with the following structure:
			// 	{
			// 		"Content": [
			// 		{
			// 			"Value": "<Provide a concise answer to the user's question here>",
			// 			"Temperatures": "<temperatures in °C (°F)>",
			// 			"Wind": {
			// 				"Speed":"<wind speed in km/h (mph)>",
			// 				"Desc":"<wind direction>"
			// 			},
			// 			"Conditions": "<Keep weather reports concise. Sparingly use emojis where appropriate.>"
			// 		}]
			// 	}`,
			// },
		},
	}
	input.Messages = append(input.Messages, message)
	resp, err := c.Client.Converse(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	// jsonResp, _ := json.MarshalIndent(resp, "", "  ")
	// fmt.Println(string(jsonResp))

	var txt string
	result := resp.Output.(*bt.ConverseOutputMemberMessage).Value
	if len(result.Content) > 0 {
		if string(resp.StopReason) == stop_reason {
			for _, tc := range result.Content {
				if _, ok := tc.(*bt.ContentBlockMemberToolUse); !ok {
					continue
				}
				var toolUse bt.ToolUseBlock = tc.(*bt.ContentBlockMemberToolUse).Value
				txt = *toolUse.Name
				toolInput := make(map[string]interface{})
				toolUse.Input.UnmarshalSmithyDocument(&toolInput)
				fmt.Println(toolInput)
				break
			}
		} else {
			responseContentBlock := result.Content[0]
			text, _ := responseContentBlock.(*bt.ContentBlockMemberText)
			txt = text.Value
		}
	}

	return txt
}

func (c *Claude3Sonnet) ConverseImage(ctx context.Context, msg string) string {
	input := &bedrockruntime.ConverseInput{
		ModelId: aws.String(c.ModelID),
		InferenceConfig: &bt.InferenceConfiguration{
			Temperature: aws.Float32(0.8),
		},
	}

	bytes, err := os.ReadFile("cat.png")
	if err != nil {
		log.Fatal(err)
	}

	message := bt.Message{
		Role: bt.ConversationRoleUser,
		Content: []bt.ContentBlock{
			&bt.ContentBlockMemberText{
				Value: "Can you help me?",
			},
			&bt.ContentBlockMemberImage{
				Value: bt.ImageBlock{
					Format: bt.ImageFormatPng,
					Source: &bt.ImageSourceMemberBytes{
						Value: bytes,
					},
				},
			},
		},
	}
	input.Messages = append(input.Messages, message)
	resp, err := c.Client.Converse(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	jsonResp, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonResp))

	var txt string
	result := resp.Output.(*bt.ConverseOutputMemberMessage).Value
	if len(result.Content) > 0 {
		responseContentBlock := result.Content[0]
		txt = responseContentBlock.(*bt.ContentBlockMemberText).Value
	}

	return txt
}
