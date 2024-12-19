package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/reddtsai/goBedrockREPL/bedrock"
)

/*
Hello, what's your name?
Tell me about the weather in taipei.
*/

type Config struct {
	AWS struct {
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		DefaultRegion   string `yaml:"default_region"`
	} `yaml:"aws"`
}

var (
	cfg Config
)

func main() {
	ctx := context.Background()
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	bedrock.LoadConfig(cfg.AWS.DefaultRegion, cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey)
	// c := bedrock.NewBedrockClient()
	// c.ListFoundationModels(ctx)
	c3 := bedrock.NewClaude3Sonnet()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[33mClaude 3 Sonnet\033[0m")
	fmt.Println("\033[33m傳送訊息至 Bedrock\033[0m")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the newline character from the input
		input = input[:len(input)-1]

		if input == "exit" {
			fmt.Println("Exiting Chat...")
			break
		}

		// output := c3.InvokeModel(ctx, input)
		output := c3.Converse(ctx, input)
		// output := c3.ConverseWithTool(ctx, input)
		// output := c3.ConverseImage(ctx, input)

		fmt.Println("\033[32m", output, "\033[0m")
	}

}
