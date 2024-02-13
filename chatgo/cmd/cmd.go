package cmd

import (
	"chatgo/chromastore"
	"chatgo/llm"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores/chroma"
)

var model *ollama.LLM
var store chroma.Store

func init() {
	var err error

	model, err = llm.New()
	if err != nil {
		log.Panic(err)
	}

	store, err = chromastore.New(model)
	if err != nil {
		log.Panic(err)
	}

	// add commands

	rootCmd.AddCommand(addFileCmd)
	rootCmd.AddCommand(queryCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gochat",
	Short: "gochat is a llm chatbot that can query documents and answer questions.",
	Long: `gochat is a llm chatbot that can query documents and answer questions, using a session chat-history.
It can also be used to add documents to the model.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
