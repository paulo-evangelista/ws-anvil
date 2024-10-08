import WebSocket from 'ws';

const socket = new WebSocket('ws://localhost:8080');

socket.on('open', () => {
    console.log('Connected to server');
    }
);

socket.on('message', (message) => {
    console.log('Received message:', message);
    }
);

socket.on('close', () => {
    console.log('Disconnected from server');
    }
);

socket.on('error', (error) => {
    console.error('An error occurred:', error);
    }
);