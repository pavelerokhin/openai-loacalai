const startButton = document.getElementById("generate");
const resultDiv = document.getElementById("result");

const maxTokens = document.getElementById('max-tokens');
const temperature = document.getElementById('temperature');
const prompt = document.getElementById('prompt');
const model = document.getElementById('model');
const requestType = document.getElementById('request-type');
const stream = document.getElementById('stream');


startButton.addEventListener("click", start);

async function start() {
    let apiUrl = 'http://localhost:8080/'
    switch (requestType.value) {
        case 'chat':
            apiUrl += 'v1/chat/completions';
            break;
        case 'completions':
            apiUrl += 'v1/completions';
            break;
        case 'embeddings':
            apiUrl += 'v1/embeddings';
            break;
    }

    const headers = {
        'Content-Type': 'application/json',
    };

    const requestBody = JSON.stringify({
        prompt: prompt.value,
        max_tokens: parseInt(maxTokens.value),
        model: model.value,
        temperature: parseFloat(temperature.value),
        stream: stream.checked,
    });

    appendUserMessage(prompt.value)

    let response;
    try {
        response = await fetch(apiUrl, {
            method: 'POST',
            headers: headers,
            body: requestBody,
        });
    } catch (error) {
        appendError(error);
        return;
    }

    if (!response.ok) {
        const errorMessage = await response.json();
        appendError(errorMessage["message"]);
        return;
    }

    const e = document.createElement("div")
    e.setAttribute("class", "message")
    const timestamp = document.createElement("div")
    const timestempText = document.createElement("span")
    timestempText.textContent = new Date().toLocaleTimeString();
    timestamp.append(timestempText)
    const message = document.createElement("div")
    message.setAttribute("class", "bot-message")
    e.append(message)
    timestamp.setAttribute("class", "timestamp")
    timestamp.textContent = new Date().toLocaleTimeString();
    e.append(timestamp)
    resultDiv.append(e)

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
    for (let i = 0; i < chunks.length; i++) {
        if (chunks[i] === "") {
            continue;
        }
        if (i === chunks.length - 1) {
            chunk = chunks[i]
        } else {
            chunk = chunks[i] + "}"
        }
        const json = JSON.parse(chunk);
        if (json["choices"]) {
            e.textContent += json["choices"][0]["text"];
        }
    }
}

function appendError(error) {
    const e = document.createElement("div")
    e.setAttribute("class", "error")
    e.textContent = error.toString();
    resultDiv.append(e)
}

function appendUserMessage(text) {
    const e = document.createElement("div")
    e.setAttribute("class", "message")
    const timestamp = document.createElement("div")
    timestamp.setAttribute("class", "timestamp")
    const timestempText = document.createElement("span")
    timestempText.textContent = new Date().toLocaleTimeString();
    timestamp.append(timestempText)
    e.append(timestamp)
    const message = document.createElement("div")
    message.setAttribute("class", "user-message")
    e.append(message)
    resultDiv.append(e)

    message.textContent = text;
}
