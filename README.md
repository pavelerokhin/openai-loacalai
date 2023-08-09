# openai-loacalai

Examples:

no streaming
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
streaming
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
