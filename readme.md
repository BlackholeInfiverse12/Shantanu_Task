# Go Blockchain Bridge

This project is a blockchain bridge  that connects Ethereum and Solana blockchains. It consists of a backend written in Go and a frontend built with React and Vite.

## Features

- **Backend**:
  - Fetches Ethereum block data using the Infura API.
  - Fetches Solana account balance using the Solana SDK.
  - Implements a simple blockchain with proof-of-work consensus.

- **Frontend**:
  - Displays blockchain data in a user-friendly interface.
  - Built with React and styled using Tailwind CSS.

## Getting Started

### Backend

1. Install Go (version 1.24.2 or higher).
2. Navigate to the `backend` folder:
   cd backend
  
  installing dependencies:
  - go mod tidy

3. run the backend server:
    go run main.go
    - The server will start at http://localhost:8080/blocks

###Frontend

1. Install dependencies:
    cd frontend
    npm i
    npm install -D tailwindcss postcss autoprefixer

2. Initialize the Tailwind CSS configuration:
    npx tailwindcss init

3. Start the development server:
    npm run dev
    - - The server will start at http://localhost:5173

### Working

1. As per given task, after running the backend, it will appear the real time Ethereum - solana transactions on stored in local         blockchain.
2. The data stored in custom blockchain is fetched and appear on webpage where Etherem and transactions will be stored and displayed   simultaneously.