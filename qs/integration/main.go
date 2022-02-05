package main

import (
	"context"
	"io/ioutil"
	"log"
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

	resp, err := client.Post(createDocUrl, "application/json", strings.NewReader("{\n    \"name\": \"fibre-product:new-year-special\",\n    \"type\": \"offer\"\n}"))
	logger.Printf("main: API-Calling : Client is calling the document api: url %v", createDocUrl)

	if err != nil {
		logger.Printf("main: API-Calling : Error: %v", err.Error())
		panic(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	logger.Printf("main: API-Calling : Client response: %v", string(bodyBytes))
}
