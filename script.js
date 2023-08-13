const startButton = document.getElementById("generate");
const resultDiv = document.getElementById("result");

const maxTokens = document.getElementById('max-tokens');
const temperature = document.getElementById('temperature');
const prompt = document.getElementById('prompt');
const model = document.getElementById('model');
const requestType = document.getElementById('request-type');
const stream = document.getElementById('stream');


startButton.addEventListener("click", go);

async function go() {
    let apiUrl = 'http://localhost:8080/'
    let requestBody = null;

    appendUserMessage(prompt.value)

    switch (requestType.value) {
        case 'chat':
            apiUrl += 'v1/chat/completions';
            requestBody = JSON.stringify({
                model: model.value,
                messages: getMessages(),
                max_tokens: parseInt(maxTokens.value),
                temperature: parseFloat(temperature.value),
                stream: stream.checked,
            });
            break;
        case 'completions':
            apiUrl += 'v1/completions';
            requestBody = JSON.stringify({
                model: model.value,
                prompt: prompt.value,
                max_tokens: parseInt(maxTokens.value),
                temperature: parseFloat(temperature.value),
                stream: stream.checked,
            });
            break;
        case 'embeddings':
            apiUrl += 'v1/embeddings';
            requestBody = JSON.stringify({
                input: prompt.value,
                model: model.value,
            });
            break;
    }

    const headers = {
        'Content-Type': 'application/json',
    };

    const response = await fetchResponse(apiUrl, headers, requestBody)

    appendBotMessage(response)
}

async function fetchResponse(apiUrl, headers, requestBody) {
    let response;
    try {
        response = await fetch(apiUrl, {
            method: 'POST',
            headers: headers,
            body: requestBody,
        });

        if (!response.ok || response.status !== 200) {
            const errorMessage = await response.text();
            appendError(errorMessage.replaceAll(/"/g, ''));
        }
    } catch (error) {
        debugger;
        appendError(error.toString());
    }

    return response
}

async function appendBotMessage(response) {
    const b = makeMessageContainer();
    const message = document.createElement("div");
    message.setAttribute("class", "bot-message");
    b.append(message);
    resultDiv.append(b);

    if (stream.checked) {
        const reader = response.body.getReader();
        let chunks = '';

        while (true) {
            const { done, value } = await reader.read();

            if (done) {
                break;
            }

            // decode the Uint8Array into a string
            let chunk = new TextDecoder("utf-8").decode(value);

            chunks = chunk.split("}\n")
            // for range of chunks
            appendBotMessageFromChunks(message, chunks)
        }
    } else {
        const responseData = await response.json();
        const json = JSON.parse(JSON.stringify(responseData, null, 2));
        message.textContent += json["choices"][0]["text"];
    }
}

function appendBotMessageFromChunks(e, chunks) {
    debugger;
    let chunk = "";
    for (let i = 0; i < chunks.length; i++) {
        if (chunks[i] === "") {
            continue;
        }
        if (i === chunks.length - 1) {
            chunk = chunks[i];
        } else {
            chunk = chunks[i] + "}";
        }

        const json = JSON.parse(chunk);

        let t = "";

        switch (json["object"]) {
            case "chat.completion.chunk":
                t = json["choices"][0]["delta"]["content"];
                if (t) {
                    e.textContent += t;
                }
                break;
            case "text_completion":
                t = json["choices"][0]["text"];
                if (t) {
                    e.textContent += t;
                }
                break;
            case "embedding": // TODO: CONTROL THIS
                t = json["choices"][0]["embedding"];
                if (t) {
                    e.textContent += t;
                }
                break;
        }
    }
}

function getMessages() {
    let messages = [
        {"role": "system", "content": "You are a helpful assistant."}
    ]

    // serialize resultDiv
    const inputs = resultDiv.querySelectorAll(".message")
    inputs.forEach(input => {
        const userMessage = input.querySelector(".user-message")
        if (userMessage) {
            const message = {"role": "user", "content": userMessage.textContent}
            messages.push(message)
            return;
        }

        const botMessage = input.querySelector(".bot-message")
        if (botMessage) {
            const message = {"role": "assistant", "content": botMessage.textContent}
            messages.push(message)
        }
    })

    return messages
}

function appendError(error) {
    const e = makeMessageContainer()
    const errorMessage = document.createElement("div")
    errorMessage.setAttribute("class", "error")
    errorMessage.textContent = error.toString();
    e.append(errorMessage)

    resultDiv.append(e)
}

function appendUserMessage(text) {
    const e = makeMessageContainer()
    const message = document.createElement("div")
    message.setAttribute("class", "user-message")
    e.append(message)
    message.textContent = text;

    resultDiv.append(e)
}

function makeMessageContainer() {
    const e = document.createElement("div")
    e.setAttribute("class", "message")
    const timestamp = document.createElement("div")
    timestamp.setAttribute("class", "timestamp")
    const timestempText = document.createElement("span")
    timestempText.textContent = new Date().toLocaleTimeString();
    timestamp.append(timestempText)
    e.append(timestamp)

    return e
}
