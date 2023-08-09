package local

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func CreateEmbeddings(ctx context.Context, req openai.EmbeddingRequest) (openai.EmbeddingResponse, error) {
	return openai.EmbeddingResponse{}, nil
}

func callModelEmbeddings(request EmbeddingRequest, response chan EmbeddingResponse) {
	// call the local model
}
