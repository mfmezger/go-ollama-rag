package ollama

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Generate(ollamaGeneration *viper.Viper, modelName string, inputText string, w http.ResponseWriter, r *http.Request) string {

	// if modelName is not provided, use the default model
	if modelName == "" {
		modelName = ollamaGeneration.GetString("model")
	}

	llm, err := ollama.New(ollama.WithModel(modelName))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llm.Call(ctx, inputText,
		llms.WithTemperature(ollamaGeneration.GetFloat64("temperature")),
		llms.WithStreamingFunc(StreamWrapper(w, r)),
	)
	if err != nil {
		log.Fatal(err)
	}

	return completion
}

// StreamWrapper is a function that takes a http.ResponseWriter and a http.Request as parameters.
// It returns a function that takes a context and a byte slice, and returns an error.
func StreamWrapper(w http.ResponseWriter, r *http.Request) func(context context.Context, chunk []byte) error {
	// The returned function
	return func(context context.Context, chunk []byte) error {
		// clientGone is the context of the request. It will be done if the client disconnects.
		clientGone := r.Context()
		// Infinite loop
		for {
			// Select statement is used for choosing which operation to perform based on the state of multiple channels.
			select {
			// If the client has disconnected, return an error
			case <-clientGone.Done():
				return fmt.Errorf("client gone")
			// If the client is still connected, write the chunk to the response writer as a JSON object
			// and flush the response writer to send the data to the client immediately.
			default:
				fmt.Fprintf(w, "{\"content\"=\"%s\"}\n", string(chunk))
				w.(http.Flusher).Flush()
				return nil
			}
		}
	}
}
