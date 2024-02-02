package ollama

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Generate(ollamaGeneration *viper.Viper, modelName string, inputText string) string {

	// if modelName is not provided, use the default model
	if modelName == "" {
		modelName = ollamaGeneration.GetString("model")
	}

	llm, err := ollama.New(ollama.WithModel(modelName))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	sb := strings.Builder{}
	completion, err := llm.Call(ctx, inputText,
		llms.WithTemperature(ollamaGeneration.GetFloat64("temperature")),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// here call function for server side events
			sb.Write(chunk)
			fmt.Println(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return completion
}
