let socket;

const connectWebSocket = () => {
    socket = new WebSocket("ws://" + window.location.host + "/ws");

    socket.onmessage = (event) => {
        const data = JSON.parse(event.data);
        createElement(data.username, data.message, new Date(data.timestamp));
    };

    socket.onerror = (e) => {
        console.error("WebSocket error:", e);
        document.getElementById("notification").textContent = "Socket error!";
    };

    socket.onclose = (e) => {
        e.wasClean ? console.log(`Closed cleanly, code=${e.code}, reason=${e.reason}`) : console.error(`Connection died (not clean)`);
    };

    socket.onopen = (e) => {
        console.log("WebSocket connection opened:", e);
        document.getElementById("notification").textContent = "Connected!";
    };
}

const sendMessage = (username) => {
    const input = document.getElementById("input");

    if (socket.readyState === WebSocket.OPEN) {
        const data = JSON.stringify({ username, message: input.value, timestamp: new Date() });
        socket.send(data);
        input.value = "";
    } else {
        console.error("Websocket is not connected!");
    }
};

connectWebSocket();
