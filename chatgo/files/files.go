package files

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func Load(ctx context.Context, name string) []schema.Document {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var loader documentloaders.Loader

	info, _ := f.Stat()

	fileType := filepath.Ext(info.Name())
	switch fileType {
	case ".pdf":
		loader = documentloaders.NewPDF(f, info.Size())

	case ".txt":
		loader = documentloaders.NewText(f)

	case ".csv":
		loader = documentloaders.NewCSV(f)

	default:
		log.Panicf("unknown file type: %s", fileType)
	}

	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(300),
		textsplitter.WithChunkOverlap(10),
	)

	dd, err := loader.LoadAndSplit(ctx, splitter)
	if err != nil {
		panic(err)
	}

	return dd
}
