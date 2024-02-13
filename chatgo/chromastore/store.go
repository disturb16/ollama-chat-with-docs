package chromastore

import (
	"chatgo/documents"
	"context"
	"log"

	chroma_go "github.com/amikos-tech/chroma-go"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/chroma"
)

const (
	namespace = "my-docs"
	serverURL = "http://localhost:8000"
	apiKey    = "A035C8C19219BA821ECEA86B64E628F8D684696D"
)

type Params struct {
	LLM      *ollama.LLM
	Store    vectorstores.VectorStore
	Query    string
	FilePath string
}

func New(llm *ollama.LLM) (chroma.Store, error) {
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		panic(err)
	}

	return chroma.New(
		chroma.WithChromaURL(serverURL),
		chroma.WithEmbedder(e),
		chroma.WithNameSpace(namespace),
		chroma.WithOpenAiAPIKey(apiKey),
		chroma.WithDistanceFunction(chroma_go.COSINE),
	)
}

func AddDoc(ctx context.Context, s *chroma.Store, path string) {
	docs := documents.Load(ctx, path)

	if s == nil {
		panic("store is nil")
	}

	log.Printf("adding %d documents for %s\n", len(docs), path)
	err := s.AddDocuments(ctx, docs)
	if err != nil {
		panic(err)
	}
}
