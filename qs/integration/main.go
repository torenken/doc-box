package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	logger := log.New(os.Stdout, "API-TEST:Docbox : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Printf("main: Started: Application initializing: version %q", "0.0.1")

	ctx := context.TODO()
	createDocUrl := "todo"

	config := clientcredentials.Config{
		ClientID:     "todo",
		ClientSecret: "todo",
		Scopes:       []string{"subscriber/docbox"},
		TokenURL:     "todo",
		AuthStyle:    oauth2.AuthStyleInParams,
	}
	client := config.Client(ctx)

	req, err := http.NewRequest("POST", createDocUrl, strings.NewReader("{\n    \"name\": \"fibre-product:new-year-special\",\n    \"type\": \"offer\"\n}"))
	if err != nil {
		logger.Printf("main: API-Request : Error: %v", err.Error())
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "todo")

	resp, err := client.Do(req)
	if err != nil {
		logger.Printf("main: API-Calling : Error: %v", err.Error())
		panic(err)
	}

	logger.Printf("main: API-Calling : Client is calling the document api: url %v", createDocUrl)

	bodyBytes, _ := io.ReadAll(resp.Body)
	logger.Printf("main: API-Calling : Client response: %v", string(bodyBytes))
}
