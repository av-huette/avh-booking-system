let  ws = null as WebSocket | null; // WebSocket connection
let isConnected = false // Connection status
const retryInterval = 3000; // Time interval for retrying connection (ms)
let retryTimeout = null as number | null; // Timeout reference for reconnection
import { useSocketStore } from "../store/socketStore";

// Function to initiate WebSocket connection
export function connectWebSocket(){
  let socketStore = useSocketStore();
  ws = new WebSocket("wss://echo.websocket.org/");
  socketStore.socket = ws;

  ws.onopen = () => {
    console.log("WebSocket connected.");
    isConnected = true;
    if (retryTimeout) {
      clearTimeout(retryTimeout); // Clear the retry timeout on successful connection
      retryTimeout = null;
    }
  };

  ws.onmessage = (event) => {
    console.log("Message received:", event.data);
    socketStore.addNotification(event.data);
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
    isConnected = false;
  };

  ws.onclose = () => {
    console.log("WebSocket closed. Retrying...");
    isConnected = false;
    attemptReconnect(); // Trigger reconnection process when the connection closes
  };
};


// Retry connection mechanism
function attemptReconnect() {
  if (!isConnected) {
    retryTimeout = setTimeout(() => {
      console.log("Attempting to reconnect...");
      connectWebSocket();
    }, retryInterval);
  }
};

export function closeConnection() {
  let socketStore = useSocketStore();
  let ws = socketStore.socket;
  if (ws) {
    console.log("Closing connection.");
    ws.close();
  }
  if (retryTimeout) {
    clearTimeout(retryTimeout);
  }
}