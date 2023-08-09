# openai-loacalai

Examples:

## No streaming examples
1. chat
```bash
curl -X POST http://localhost:8080/v1/chat/completions \
    -H 'Content-Type: application/json' \
    -d '{
        "model": "gpt-3.5-turbo",
        "messages": [
          {"role": "system", "content": "You are a hasidic rabbi."},
          {"role": "user", "content": "Tell me a random fact about Judaism."}
        ],
        "max_tokens": 50,
        "temperature": 0.7
    }'
```
2. completions
```bash
curl -X POST http://localhost:8080/v1/completions \
    -H "Content-Type: application/json" \
    -d '{
        "model": "text-davinci-003",
        "prompt": "Say you like to be in streaming mode",
        "max_tokens": 7,
        "temperature": 0
    }'
```
3. embeddings
```bash
curl -X POST 127.0.0.1:8080/v1/embeddings \
    -H "Content-Type: application/json" \
    -d '{
      "input": "The food was delicious and the waiter...",
      "model": "text-embedding-ada-002"
    }'
```

## Streaming examples
1. chat
```bash
curl -X POST http://localhost:8080/v1/chat/completions \
    -H 'Content-Type: application/json' \
    -d '{
        "model": "gpt-3.5-turbo",
        "messages": [
        {"role": "system", "content": "You are a reform rabbi."},
        {"role": "user", "content": "Tell me a random fact about Judaism."}
        ],
        "max_tokens": 50,
        "temperature": 0.7,
        "stream": true
    }'
```

2. completions
```bash
curl -X POST http://localhost:8080/v1/completions \
    -H "Content-Type: application/json" \
    -d '{
        "model": "text-davinci-003",
        "prompt": "Say something funny",
        "max_tokens": 7,
        "temperature": 0,
        "stream": true
    }'
```
