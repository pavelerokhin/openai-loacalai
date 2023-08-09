package local

import "github.com/sashabaranov/go-openai"

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	N                int                     `json:"n,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
	Stop             []string                `json:"stop,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
	// LogitBias is must be a token id string (specified by their token ID in the tokenizer), not a word string.
	// incorrect: `"logit_bias":{"You": 6}`, correct: `"logit_bias":{"1639": 6}`
	// refs: https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias
	LogitBias map[string]int `json:"logit_bias,omitempty"`
	User      string         `json:"user,omitempty"`
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`

	// This property isn't in the official documentation, but it's in
	// the documentation for the official library for python:
	// - https://github.com/openai/openai-python/blob/main/chatml.md
	// - https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
	Name string `json:"name,omitempty"`
}

// CompletionRequest represents a request structure for completion API.
type CompletionRequest struct {
	Model            string   `json:"model"`
	Prompt           any      `json:"prompt,omitempty"`
	Suffix           string   `json:"suffix,omitempty"`
	MaxTokens        int      `json:"max_tokens,omitempty"`
	Temperature      float32  `json:"temperature,omitempty"`
	TopP             float32  `json:"top_p,omitempty"`
	N                int      `json:"n,omitempty"`
	Stream           bool     `json:"stream,omitempty"`
	LogProbs         int      `json:"logprobs,omitempty"`
	Echo             bool     `json:"echo,omitempty"`
	Stop             []string `json:"stop,omitempty"`
	PresencePenalty  float32  `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32  `json:"frequency_penalty,omitempty"`
	BestOf           int      `json:"best_of,omitempty"`
	// LogitBias is must be a token id string (specified by their token ID in the tokenizer), not a word string.
	// incorrect: `"logit_bias":{"You": 6}`, correct: `"logit_bias":{"1639": 6}`
	// refs: https://platform.openai.com/docs/api-reference/completions/create#completions/create-logit_bias
	LogitBias map[string]int `json:"logit_bias,omitempty"`
	User      string         `json:"user,omitempty"`
}

type EmbeddingRequest struct {
	Input any    `json:"input"`
	User  string `json:"user"`
}

// CompletionResponse represents a response structure for completion API.
type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
}

// CompletionChoice represents one of possible completions.
type CompletionChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
	LogProbs     LogprobResult `json:"logprobs"`
}

// LogprobResult represents logprob result of Choice.
type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

// EmbeddingResponse represents a response structure for embedding API.
type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Index   int                   `json:"index"`
	Message ChatCompletionMessage `json:"message"`
}

func ToLocalChatCompletionRequest(req openai.ChatCompletionRequest) ChatCompletionRequest {
	return ChatCompletionRequest{
		Model:            req.Model,
		Messages:         ToLocalMessages(req.Messages),
		MaxTokens:        req.MaxTokens,
		Temperature:      req.Temperature,
		TopP:             req.TopP,
		N:                req.N,
		Stream:           req.Stream,
		Stop:             req.Stop,
		PresencePenalty:  req.PresencePenalty,
		FrequencyPenalty: req.FrequencyPenalty,
		LogitBias:        req.LogitBias,
		User:             req.User,
	}
}

func ToLocalMessages(messages []openai.ChatCompletionMessage) []ChatCompletionMessage {
	var localMessages []ChatCompletionMessage
	for _, message := range messages {
		localMessages = append(localMessages, ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
			Name:    message.Name,
		})
	}

	return localMessages
}

func ToOpenAIChatCompletionResponse(resp ChatCompletionResponse) openai.ChatCompletionResponse {
	return openai.ChatCompletionResponse{
		ID:      resp.ID,
		Object:  resp.Object,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: ToOpenAIChatCompletionChoices(resp.Choices),
	}
}

func ToOpenAIChatCompletionChoices(choices []ChatCompletionChoice) []openai.ChatCompletionChoice {
	var openAIChoices []openai.ChatCompletionChoice
	for _, choice := range choices {
		openAIChoices = append(openAIChoices, openai.ChatCompletionChoice{
			Index:   choice.Index,
			Message: ToOpenAIChatCompletionMessage(choice.Message),
		})
	}

	return openAIChoices
}

func ToOpenAIChatCompletionMessage(message ChatCompletionMessage) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    message.Role,
		Content: message.Content,
		Name:    message.Name,
	}
}

func ToLocalCompletionRequest(request openai.CompletionRequest) CompletionRequest {
	return CompletionRequest{
		Model:            request.Model,
		Prompt:           request.Prompt,
		Suffix:           request.Suffix,
		MaxTokens:        request.MaxTokens,
		Temperature:      request.Temperature,
		TopP:             request.TopP,
		N:                request.N,
		Stream:           request.Stream,
		LogProbs:         request.LogProbs,
		Echo:             request.Echo,
		Stop:             request.Stop,
		PresencePenalty:  request.PresencePenalty,
		FrequencyPenalty: request.FrequencyPenalty,
		BestOf:           request.BestOf,
		LogitBias:        request.LogitBias,
		User:             request.User,
	}
}

func ToLocalEmbeddingRequest(request openai.EmbeddingRequest) EmbeddingRequest {
	return EmbeddingRequest{
		Input: request.Input,
		User:  request.User,
	}
}

func ToLocalCompletionResponse(response openai.CompletionResponse) CompletionResponse {
	return CompletionResponse{
		ID:      response.ID,
		Object:  response.Object,
		Created: response.Created,
		Model:   response.Model,
		Choices: ToLocalCompletionChoices(response.Choices),
	}
}

func ToLocalCompletionChoices(choices []openai.CompletionChoice) []CompletionChoice {
	var localChoices []CompletionChoice
	for _, choice := range choices {
		localChoices = append(localChoices, CompletionChoice{
			Text:         choice.Text,
			Index:        choice.Index,
			FinishReason: choice.FinishReason,
			LogProbs:     ToLocalLogprobResult(choice.LogProbs),
		})
	}

	return localChoices
}

func ToLocalLogprobResult(logprobs openai.LogprobResult) LogprobResult {
	return LogprobResult{
		Tokens:        logprobs.Tokens,
		TokenLogprobs: logprobs.TokenLogprobs,
		TopLogprobs:   logprobs.TopLogprobs,
		TextOffset:    logprobs.TextOffset,
	}
}

func ToLocalEmbeddingResponse(response openai.EmbeddingResponse) EmbeddingResponse {
	return EmbeddingResponse{
		Object: response.Object,
		Data:   ToLocalEmbeddings(response.Data),
	}
}

func ToLocalEmbeddings(embeddings []openai.Embedding) []Embedding {
	var localEmbeddings []Embedding
	for _, embedding := range embeddings {
		localEmbeddings = append(localEmbeddings, Embedding{
			Object:    embedding.Object,
			Embedding: embedding.Embedding,
			Index:     embedding.Index,
		})
	}

	return localEmbeddings
}

// Embedding is a special format of data representation that can be easily utilized by machine
// learning models and algorithms. The embedding is an information dense representation of the
// semantic meaning of a piece of text. Each embedding is a vector of floating point numbers,
// such that the distance between two embeddings in the vector space is correlated with semantic similarity
// between two inputs in the original format. For example, if two texts are similar,
// then their vector representations should also be similar.
type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

func ToOpenAICompletionResponse(resp CompletionResponse) openai.CompletionResponse {
	return openai.CompletionResponse{
		ID:      resp.ID,
		Object:  resp.Object,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: ToOpenAICompletionChoices(resp.Choices),
	}
}

func ToOpenAICompletionChoices(choices []CompletionChoice) []openai.CompletionChoice {
	var openAIChoices []openai.CompletionChoice
	for _, choice := range choices {
		openAIChoices = append(openAIChoices, openai.CompletionChoice{
			Text:         choice.Text,
			Index:        choice.Index,
			FinishReason: choice.FinishReason,
			LogProbs:     ToOpenAILogprobResult(choice.LogProbs),
		})
	}

	return openAIChoices
}

func ToOpenAILogprobResult(logprobs LogprobResult) openai.LogprobResult {
	return openai.LogprobResult{
		Tokens:        logprobs.Tokens,
		TokenLogprobs: logprobs.TokenLogprobs,
		TopLogprobs:   logprobs.TopLogprobs,
		TextOffset:    logprobs.TextOffset,
	}
}

func ToOpenAIEmbeddingsResponse(response EmbeddingResponse) openai.EmbeddingResponse {
	return openai.EmbeddingResponse{
		Object: response.Object,
		Data:   ToOpenAIEmbeddings(response.Data),
	}
}

func ToOpenAIEmbeddings(embeddings []Embedding) []openai.Embedding {
	var openAIEmbeddings []openai.Embedding
	for _, embedding := range embeddings {
		openAIEmbeddings = append(openAIEmbeddings, openai.Embedding{
			Object:    embedding.Object,
			Embedding: embedding.Embedding,
			Index:     embedding.Index,
		})
	}

	return openAIEmbeddings
}
