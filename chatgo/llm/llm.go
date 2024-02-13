package llm

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
)

const (
	modelName string = "mistral:7b"
	ollamaURL string = "http://localhost:11434"
)

var (
	temperature = chains.WithTemperature(0.4)
	chatMemory  = memory.NewConversationBuffer()
)

func init() {
	chatMemory.InputKey = "question"
	chatMemory.OutputKey = "text"
	chatMemory.MemoryKey = "context"
}

func New() (*ollama.LLM, error) {
	return ollama.New(
		ollama.WithModel(modelName),
		ollama.WithServerURL(ollamaURL),
	)
}

func streamingCallback(ctx context.Context, chunk []byte) error {
	log.Println(string(chunk))
	return nil
}

// Query executes a query using:
// the provided LLM model, list of documents and query string.
func Query(ctx context.Context, model *ollama.LLM, dd []schema.Document, query string) {
	log.Printf("\nusing %d documents\n", len(dd))
	log.Printf("%s \n\n", query)

	chain := chains.LoadStuffQA(model)
	// chain.LLMChain.Memory = chatMemory

	input := map[string]any{
		"question":        query,
		"input_documents": dd,
	}

	response, err := chains.Call(
		ctx,
		chain,
		input,
		temperature,
		// chains.WithStreamingFunc(streamingCallback),
	)

	if err != nil {
		panic(err)
	}

	log.Println(response["text"])
}
