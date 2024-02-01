package main

import (
    "context"
    "fmt"
    "log"

    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/ollama"
)

func main() {
    llm, err := ollama.New(ollama.WithModel("zephyr"))
    if err != nil {
        log.Fatal(err)
    }
    ctx := context.Background()
    completion, err := llm.Call(ctx, "Human: Who was the first man to walk on the moon?\nAssistant:",
        llms.WithTemperature(0.8),
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

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	"github.com/spf13/viper"
// 	"io"
// 	"log"
// 	"net/http"
// )


// // func main() {
// //     v, err := loadConfig()
// //     if err != nil {
// //         log.Fatalf("Fatal error config file: %s \n", err)
// //     }

// //     model := v.GetString("ollama.model")
// //     temperature := v.GetFloat64("ollama.temperature")

// //     fmt.Printf("Model: %s, Temperature: %f\n", model, temperature)

// //     r := mux.NewRouter()

// //     r.HandleFunc("/embed-pdf", EmbedPDFHandler).Methods(http.MethodPost)
// //     r.HandleFunc("/embed-text", EmbedTextHandler).Methods(http.MethodPost)
// //     r.HandleFunc("/semantic-search", SemanticSearchHandler).Methods(http.MethodGet)
// //     r.HandleFunc("/qa", QAHandler).Methods(http.MethodPost)
// //     r.HandleFunc("/generate", GenerateHandler).Methods(http.MethodPost)

// //     log.Fatal(http.ListenAndServe(":8000", r))
// // }

// // func loadConfig() (*viper.Viper, error) {
// //     v := viper.New()
// //     v.SetConfigName("ollama")
// //     v.SetConfigType("yaml")
// //     v.AddConfigPath("configs")
// //     err := v.ReadInConfig()
// //     if err != nil {
// //         return nil, err
// //     }
// //     return v, nil
// // }


// // generate a generate request that takes a prompt and returns a response
// func GenerateHandler(w http.ResponseWriter, r *http.Request) {
//     gen := NewGenerator()
//     body := []byte(`{"model":"mistral"}`)
//     responseData, err := gen.Generate(body)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.Write(responseData)
// }

// func EmbedPDFHandler(w http.ResponseWriter, r *http.Request) {

// }

// func SemanticSearchHandler(w http.ResponseWriter, r *http.Request) {

// }

// func QAHandler(w http.ResponseWriter, r *http.Request) {

// }

// func EmbedTextHandler(w http.ResponseWriter, r *http.Request) {
//     err := r.ParseMultipartForm(10 << 20)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     file, _, err := r.FormFile("file")
//     if err != nil {
//         http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
//         return
//     }
//     defer file.Close()

//     fileBytes, err := io.ReadAll(file)
//     if err != nil {
//         http.Error(w, "Error reading the file", http.StatusInternalServerError)
//         return
//     }

//     apiURL := "http://example.com/api"
//     req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, apiURL, bytes.NewBuffer(fileBytes))
//     if err != nil {
//         http.Error(w, "Error creating request to external API", http.StatusInternalServerError)
//         return
//     }

//     req.Header.Set("Content-Type", "application/text")

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         http.Error(w, "Error sending request to external API", http.StatusInternalServerError)
//         return
//     }
//     defer resp.Body.Close()

//     w.WriteHeader(http.StatusOK)
//     w.Write([]byte("File processed and sent to API successfully"))
// }
