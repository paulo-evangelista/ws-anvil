import WebSocket from 'ws';

const socket = new WebSocket('wss://blockchain.vert-tech.dev/server/');

socket.on('open', () => {
    console.log('Connected to server');
    }
);

socket.on('message', (message) => {
    console.log('>', message.toString());
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