package main

import (
	"bytes"
	"encoding/json"
	"io"
    log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mfmezger/go-ollama-rag/internal/ollama"
	"github.com/spf13/viper"
)


var config *viper.Viper

func init() {
	var err error
	config, err = loadConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
    log.Info("Config loaded successfully!")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/embed-pdf", EmbedPDFHandler).Methods(http.MethodPost)
	r.HandleFunc("/embed-text", EmbedTextHandler).Methods(http.MethodPost)
	r.HandleFunc("/semantic-search", SemanticSearchHandler).Methods(http.MethodGet)
	r.HandleFunc("/qa", QAHandler).Methods(http.MethodPost)
	r.HandleFunc("/generate/model/{modelName}", GenerateHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", r))

    log.Info("Server started successfully!")
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("ollama")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

// generate a generate request that takes a prompt and returns a response
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	ollamaGeneration := config.Sub("ollama_generation")

	modelName := mux.Vars(r)["modelName"]

	content, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	text := string(content)

	answer := ollama.Generate(ollamaGeneration, modelName, text)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": answer})
}

func EmbedPDFHandler(w http.ResponseWriter, r *http.Request) {

}

func SemanticSearchHandler(w http.ResponseWriter, r *http.Request) {

}

func QAHandler(w http.ResponseWriter, r *http.Request) {

}

func EmbedTextHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	apiURL := "http://example.com/api"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, apiURL, bytes.NewBuffer(fileBytes))
	if err != nil {
		http.Error(w, "Error creating request to external API", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/text")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request to external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File processed and sent to API successfully"))
}
