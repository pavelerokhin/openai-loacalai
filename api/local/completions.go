package local

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func CreateCompletion(ctx context.Context, req openai.CompletionRequest) (openai.CompletionResponse, error) {
	return openai.CompletionResponse{}, nil
}

func CreateCompletionStreaming(ctx context.Context, req openai.CompletionRequest, resp chan openai.CompletionResponse, errors chan error) {
	// cast the request to the local model
	request := ToLocalCompletionRequest(req)
	response := make(chan CompletionResponse)

	// call the local model
	go func() {
		defer close(response)
		<-ctx.Done()
	}()

	go callModelCompletion(request, response)

	for r := range response {
		resp <- ToOpenAICompletionResponse(r)
	}
}

func callModelCompletion(request CompletionRequest, response chan CompletionResponse) {
	// call the local model
}
