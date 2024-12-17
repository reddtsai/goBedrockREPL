package bedrock

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
)

var (
	awsCfg aws.Config
)

func LoadConfig(region, accessKeyID, secretAccessKey string) {
	var err error
	awsCfg, err = config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
}

type BedrockClient struct {
	Client *bedrock.Client
}

func NewBedrockClient() *BedrockClient {
	return &BedrockClient{
		Client: bedrock.NewFromConfig(awsCfg),
	}
}

func (b *BedrockClient) ListFoundationModels(ctx context.Context) {
	result, err := b.Client.ListFoundationModels(ctx, &bedrock.ListFoundationModelsInput{})
	if err != nil {
		log.Fatalf("failed to list foundation models, %v", err)
	}

	// Print the model IDs
	for _, model := range result.ModelSummaries {
		s := fmt.Sprintf("Name: %s | Provider: %s | Id: %s | Modality: %s", *model.ModelName, *model.ProviderName, *model.ModelId, model.OutputModalities)
		fmt.Println(s)
	}
}
