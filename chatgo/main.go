package main

import (
	"bufio"
	"chatgo/chromastore"
	"chatgo/llm"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/chroma"
)

var threshold = vectorstores.WithScoreThreshold(0.2)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			log.Printf("received %s, exiting\n", sig)
			cancel()
		}
	}()

	query, filePath := flagVars()
	model := llm.New()
	store := chromastore.New(model)

	if filePath != "" {
		chromastore.AddDoc(ctx, store, filePath)
	}

	isFirstQuestion := true

	if query != "" {
		reader := bufio.NewReader(os.Stdin)
		for {

			if isFirstQuestion {
				askQuery(ctx, model, store, query)
			}

			// keep getting input from command line until user exits
			// or sends a signal
			io.WriteString(os.Stdout, "Enter query: ")
			query, _ = reader.ReadString('\n')
			query = strings.Replace(query, "\n", "", -1)

			if query == "exit" {
				log.Println("exiting...")
				cancel()
				return
			}

			askQuery(ctx, model, store, query)
			isFirstQuestion = false
		}
	}
}

func flagVars() (query, filePath string) {
	query = ""
	filePath = ""
	flag.StringVar(&query, "q", "", "prompt to use")
	flag.StringVar(&filePath, "f", "", "file path to add")
	flag.Parse()

	return query, filePath
}

func askQuery(ctx context.Context, model *ollama.LLM, store *chroma.Store, query string) {
	dd, err := store.SimilaritySearch(ctx, query, 50, threshold)
	if err != nil {
		panic(err)
	}

	llm.Query(ctx, model, dd, query)
}
