package main

import (
    "context"
    "fmt"
    "log"
	"github.com/spf13/viper"
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/ollama"
)

type Ollama struct {
    llm         *LLM
    model       string
    temperature float64
}

func NewOllama() (*Ollama, error) {

    v := viper.New()

    v.SetConfigName("ollama")  // name of config file (without extension)
    v.SetConfigType("yaml")    // or viper.SetConfigType("YAML")
    v.AddConfigPath("configs") // optionally look for config in the working directory

    err := v.ReadInConfig()    // Find and read the config file
    if err != nil {            // Handle errors reading the config file
        return nil, err
    }

    // Read configuration
    model := v.GetString("ollama.model")
    temperature := v.GetFloat64("ollama.temperature")

       llm, err := ollama.New(ollama.WithModel(model))
    if err != nil {
        log.Fatal(err)
    }
    ctx := context.Background()
    completion, err := llm.Call(ctx, "Human: Who was the first man to walk on the moon?\nAssistant:",
        llms.WithTemperature(temperature),
        llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
            fmt.Print(string(chunk))
            return nil
        }),
    )
    if err != nil {
        log.Fatal(err)
    }

    _ = completion
}