package cmd

import (
	"bufio"
	"chatgo/llm"
	"context"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/vectorstores"
)

var (
	threshold       float32 = 0.1
	maxDocsResponse int     = 50
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query the model",
	Long:  `query the model using a prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		loopQuery(ctx)
	},
}

// loopQuery is a loop that keeps asking for input from the user
// and querying the model. It exits when the user types "exit".
func loopQuery(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)
	for {
		io.WriteString(os.Stdout, "Enter prompt: ")
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, "\n", "", -1)

		if query == "exit" {
			log.Println("exiting...")
			return
		}

		if err := execQuery(ctx, query); err != nil {
			log.Println(err)
		}
	}
}

func execQuery(ctx context.Context, query string) error {
	dd, err := store.SimilaritySearch(
		ctx,
		query,
		maxDocsResponse,
		vectorstores.WithScoreThreshold(threshold),
	)

	if err != nil {
		return err
	}

	llm.Query(ctx, model, dd, query)
	return nil
}
