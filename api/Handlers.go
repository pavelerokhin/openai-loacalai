package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"

	"OpenAI-api/api/local"
)

const (
	LLMatteo = "llm-matteo"
)

var (
	client *openai.Client
)

func Init() {
	key := viper.GetString("open-ai-api-key")
	if key != "" {
		client = openai.NewClient(key)
	}
}

func HandleChat(c echo.Context) error {
	// get the chat completion request from Echo request body
	var request openai.ChatCompletionRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()

	if request.Stream {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		enc := json.NewEncoder(c.Response())

		if request.Model == LLMatteo {
			response := make(chan openai.ChatCompletionResponse)
			errors := make(chan error)
			go local.CreateChatCompletionStream(ctx, request, response, errors)
			for {
				select {
				case resp, ok := <-response:
					if !ok {
						break
					}
					if err := enc.Encode(resp); err != nil {
						return c.JSON(http.StatusInternalServerError, err.Error())
					}
					c.Response().Flush()
				case err := <-errors:
					return c.JSON(http.StatusInternalServerError, err.Error())
				}
			}
		} else {
			if client == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "OpenAI client is not initialized")
			}
			stream, err := client.CreateChatCompletionStream(ctx, request)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			for {
				resp, e := stream.Recv()
				if e != nil {
					break
				}
				if err := enc.Encode(resp); err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}
				c.Response().Flush()
			}
		}

		return nil
	}

	var response openai.ChatCompletionResponse
	var err error

	if request.Model == LLMatteo {
		response, err = local.CreateChatCompletion(ctx, request)
	} else {
		if client == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "OpenAI client is not initialized")
		}
		response, err = client.CreateChatCompletion(ctx, request)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func HandleCompletions(c echo.Context) error {
	var request openai.CompletionRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()

	if request.Stream {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		enc := json.NewEncoder(c.Response())

		if request.Model == LLMatteo { // local model
			response := make(chan openai.CompletionResponse)
			errors := make(chan error)
			go local.CreateCompletionStreaming(ctx, request, response, errors)
			for {
				select {
				case resp, ok := <-response:
					if !ok {
						break
					}
					if err := enc.Encode(resp); err != nil {
						return c.JSON(http.StatusInternalServerError, err.Error())
					}
					c.Response().Flush()
				case err := <-errors:
					return c.JSON(http.StatusInternalServerError, err.Error())
				}
			}
		} else { // openAI model
			if client == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "OpenAI client is not initialized")
			}
			stream, err := client.CreateCompletionStream(ctx, request)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			for {
				resp, e := stream.Recv()
				if e != nil {
					break
				}
				if err := enc.Encode(resp); err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
				}
				c.Response().Flush()
			}
		}

		return nil
	}

	var response openai.CompletionResponse
	var err error

	if request.Model == LLMatteo {
		response, err = local.CreateCompletion(ctx, request)
	} else {
		if client == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "OpenAI client is not initialized")
		}
		response, err = client.CreateCompletion(ctx, request)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func HandleEmbeddings(c echo.Context) error {
	var request openai.EmbeddingRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()

	var response openai.EmbeddingResponse
	var err error

	if request.Model.String() == LLMatteo {
		response, err = local.CreateEmbeddings(ctx, request)
	} else {
		if client == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "OpenAI client is not initialized")
		}
		response, err = client.CreateEmbeddings(ctx, request)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
