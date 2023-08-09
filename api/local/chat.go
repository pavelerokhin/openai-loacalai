package local

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	// cast the request to the local model

	// call the local model

	// cast the response to the openai model

	return openai.ChatCompletionResponse{}, nil
}

func CreateChatCompletionStream(ctx context.Context, req openai.ChatCompletionRequest, resp chan openai.ChatCompletionResponse, errors chan error) {
	// cast the request to the local model
	request := ToLocalChatCompletionRequest(req)
	response := make(chan ChatCompletionResponse)
	// call the local model
	go func() {
		defer close(resp)
		<-ctx.Done()
	}()

	go callModelChat(request, response, errors)

	for r := range response {
		resp <- ToOpenAIChatCompletionResponse(r)
	}
}

func callModelChat(request ChatCompletionRequest, response chan ChatCompletionResponse, errors chan error) {
	// call the local model
}
