package main

import (
	"chatgo/chromastore"
	"chatgo/files"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores"
)

var (
	scoreThreshold = vectorstores.WithScoreThreshold(0.3)
	temperature    = chains.WithTemperature(0.5)
)

const (
	modelName  string = "llama2:latest"
	collection string = "mydocs"
	ollamaURL  string = "http://localhost:11434"
)

type Params struct {
	LLM      *ollama.LLM
	Store    vectorstores.VectorStore
	Query    string
	FilePath string
}

func main() {
	now := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("received %s, exiting\n", sig)
			cancel()
		}
	}()

	query, filePath := flagVars()
	llm := LLM()
	store := Store(ctx, llm)

	params := Params{
		LLM:      llm,
		Store:    store,
		Query:    query,
		FilePath: filePath,
	}

	if filePath != "" {
		addDoc(ctx, params)
	}

	if query != "" {
		queryLLM(ctx, params)
	}

	log.Printf("finished in %f mins\n", time.Since(now).Minutes())
}

func flagVars() (query, filePath string) {
	query = ""
	filePath = ""
	flag.StringVar(&query, "q", "", "prompt to use")
	flag.StringVar(&filePath, "f", "", "file path to add")
	flag.Parse()

	return query, filePath
}

func LLM() *ollama.LLM {
	llm, err := ollama.New(
		ollama.WithModel(modelName),
		ollama.WithServerURL(ollamaURL),
	)

	if err != nil {
		panic(err)
	}

	return llm
}

func Store(ctx context.Context, llm *ollama.LLM) vectorstores.VectorStore {
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		panic(err)
	}

	store, err := chromastore.New(e)
	if err != nil {
		panic(err)
	}

	return store
}

func queryLLM(ctx context.Context, params Params) {
	fmt.Printf("\n")
	dd, err := params.Store.SimilaritySearch(ctx, params.Query, 50, scoreThreshold)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")

	// exit early if no docs found
	fmt.Printf("found %d documents\n", len(dd))
	if len(dd) == 0 {
		return
	}

	log.Printf("%s \n\n", params.Query)

	chain := chains.LoadStuffQA(params.LLM)

	input := map[string]any{
		"question":        params.Query,
		"input_documents": dd,
	}

	response, err := chains.Call(ctx, chain, input, temperature)
	if err != nil {
		panic(err)
	}

	log.Println(response["text"])
}

func addDoc(ctx context.Context, params Params) {

	docs := files.Load(ctx, params.FilePath)

	log.Printf("adding %d documents for %s\n", len(docs), params.FilePath)
	err := params.Store.AddDocuments(ctx, docs)
	if err != nil {
		panic(err)
	}
}
