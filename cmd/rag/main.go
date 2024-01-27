package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigName("ollama")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // or viper.SetConfigType("YAML")
	viper.AddConfigPath("configs") // optionally look for config in the working directory
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	// Read configuration
	model := viper.GetString("ollama.model")
	temperature := viper.GetFloat64("ollama.temperature")

	fmt.Printf("Model: %s, Temperature: %f\n", model, temperature)

	r := mux.NewRouter()

	// // Define endpoints
	r.HandleFunc("/embed-pdf", EmbedPDFHandler).Methods("POST")
	r.HandleFunc("/embed-text", EmbedTextHandler).Methods("POST")
	r.HandleFunc("/semantic-search", SemanticSearchHandler).Methods("GET")
	r.HandleFunc("/qa", QAHandler).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

	gen := NewGenerator()

	body := []byte(`{"model":"mistral"}`)
	responseData, err := gen.Generate(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

func EmbedPDFHandler(w http.ResponseWriter, r *http.Request) {

}

func SemanticSearchHandler(w http.ResponseWriter, r *http.Request) {

}

func QAHandler(w http.ResponseWriter, r *http.Request) {

}

func EmbedTextHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	// Sent to the Ollama Endpoint
	apiURL := "http://example.com/api" // Replace with your API URL
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(fileBytes))
	if err != nil {
		http.Error(w, "Error creating request to external API", http.StatusInternalServerError)
		return
	}

	// Set appropriate headers, if needed
	req.Header.Set("Content-Type", "application/text")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request to external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Optionally, handle the response from the API
	// ...

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File processed and sent to API successfully"))
}
