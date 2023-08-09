const startButton = document.getElementById("generate");
const promptInput = document.getElementById("prompt");
const resultDiv = document.getElementById("result");

startButton.addEventListener("click", startStreaming);

async function startStreaming() {
    const promptText = promptInput.value;

    const requestData = {
        model: "text-davinci-003",
        prompt: promptText,
        max_tokens: 50,
        temperature: 0.7,
        stream: true
    };

    try {
        const response = await fetch("http://localhost:8080/v1/completions", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(requestData)
        });

        if (response.ok) {
            let streamReader = response.body.getReader();
            readStream(streamReader);
        } else {
            resultDiv.innerHTML = "<p>Error starting stream.</p>";
        }
    } catch (error) {
        resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    }
}

async function readStream(streamReader) {
        while (true) {
            try {
                const { done, value } = await streamReader.read();

                if (done) {
                    resultDiv.innerHTML += "<p>Stream completed.</p>";
                    break;
                }

                let c = new TextDecoder().decode(value);
                // decode text to json
                let json = JSON.parse(c);
                console.log(json);

                if (json.choices[0].text === promptInput.value) {
                    continue;
                }
                resultDiv.innerHTML += json.choices[0].text;
            } catch (error) {
                debugger;
                resultDiv.innerHTML += `<p>Error reading stream: ${error.message}</p>`;
            }
        }
}
