package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func getResponse(m string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(m))
	if err != nil {
		log.Fatal(err)
	}
	if resp != nil {
		fmt.Println("Token Count", resp.Candidates[0].TokenCount)
		fmt.Println("AI Response")
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	} else {
		fmt.Println("Empty response")
	}
}
