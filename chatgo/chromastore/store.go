package chromastore

import (
	chroma_go "github.com/amikos-tech/chroma-go"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/chroma"
)

const (
	namespace = "my-docs"
	serverURL = "http://localhost:8000"
	apiKey    = "A035C8C19219BA821ECEA86B64E628F8D684696D"
)

func New(e embeddings.Embedder) (*chroma.Store, error) {
	s, err := chroma.New(
		chroma.WithChromaURL(serverURL),
		chroma.WithEmbedder(e),
		chroma.WithNameSpace(namespace),
		chroma.WithOpenAiAPIKey(apiKey),
		chroma.WithDistanceFunction(chroma_go.COSINE),
	)

	return &s, err
}
