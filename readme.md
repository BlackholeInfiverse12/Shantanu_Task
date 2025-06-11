Go Blockchain Bridge Documentation
=====================================

### Overview
--------
Go Blockchain Bridge is a cross-chain bridge connecting Ethereum and Solana, featuring a Go backend and a React + Vite frontend. It standardizes cross-chain messages, validates them, prevents replay attacks, and provides a real-time dashboard for monitoring transactions.

### 1. Setup
Prerequisites
Go: v1.24.2 or higher
Node.js: v18+ and npm
Git: For cloning the repository

 **Directory Structure**:
  
    TTCont/
    |-- backend/
        |
        |-- core/
            |-- blockchain/ 
            |
            |-- bridge/ 
            |       |-- blockchain/
            |       |-- cmd/
            |       |-- internal/   
            |       |-- Web/ 
            |           |-- static/
            |           |-- templates/ 
            |
            |-- ethereum/
            |       |-- eth-event-listener/
            |                   |-- cmd/
            |                   |-- internal/
            |
            |-- solana/
            |       |-- sol-event-listener/
            |     
            |-- go-bridge-server/
                    |-- cmd/ 
                    |       |-- server/
                    |-- internal/

    |-- frontend/


**----------------------------------------**

### 2. Configuration
    Environment Variables
    For flexibility and security, use a .env file or config.json for sensitive and environment-specific parameters.

    Create a .env file in both backend and frontend directories as needed.

**backend/.env**:
    ETH_RPC_URL=wss://mainnet.infura.io/ws/v3/688f2501b7114913a6b23a029bd43c9d
    SOL_RPC_URL=https://api.mainnet-beta.solana.com
    ETH_ERC20_CONTRACT=0xdAC17F958D2ee523a2206206994597C13D831ec7
    BACKEND_PORT=8083
    ETH_LISTENER_PORT=8081
    SOL_LISTENER_PORT=8082

**frontend/.env**:

    VITE_BACKEND_URL=http://localhost:8083
    VITE_ETH_EVENTS_URL=http://localhost:8081/events
    VITE_SOL_EVENTS_URL=http://localhost:8082/events
    
**config.json**:

    {
        "ethRpcUrl": "wss://mainnet.infura.io/ws/v3/688f2501b7114913a6b23a029bd43c9d",
        "solRpcUrl": "https://api.mainnet-beta.solana.com",
        "ethErc20Contract": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
        "backendPort": 8083
    }

**-------------------------------------------**

### 3. Modules:-
**Backend**
    1. Install Go dependencies:
    
    cd backend
    go mod tidy

**-------------------------------------**

    2. Running Core Blockchain:

    go run main.go

**-------------------------------------**

    3. Module - Core:
            It includes the blockchain, bridge, Ethereum, Solana, and Go-Bridge-Server packages
    
**-------------------------------------**

    4. Module - Bridge:

            Bridge is the central module that ties everything together.
            It is responsible for handling transactions, events, and communication between the blockchain and the Go-Bridg and relay servers.

             cd bridge/cmd
            go run main.go

         Output: 
            Server started at http://localhost:8083/events

         This will start the relay server that will capture Ethereum and Solana Transactions and forward events from the blockchain to the Go-Bridge.
    
**---------------------------------**

    5. Module - Ethereum:
    It includes the Ethereum package that handles Ethereum transactions and events.

        cd ethereum/eth-event-listener/cmd
        go run main.go

**---------------------------------**

    6. Module - Solana:
    It includes the Solana package that handles Solana transactions and events.

        cd solana/sol-event-listener
        go run solListener.go

**-------------------------------------**

    7. Module - Go-Bridge-Server:

    It includes the Go-Bridge-Server package that handles communication between the blockchain and the Go-Bridge
    Main function: 
    Working bridge server in /cmd/server.go with Postman or cURL Test.

        cd go-bridge-server/cmd/server
        go run main.go

**----------------------------------------**

    8.  Bridge Dashboard (CLI or Minimal UI):

    It shows recent events, relayed messages, validation results
    
        Server started at http://localhost:8083/  (Recent Events)
        Server started at http://localhost:8083/validation  (Validation results)

**--------------------------------------------**

### 4. Final Run
 
**a. Installing Dependencies**:

    run the following command in the each module:
        go mod tidy

 
**Run the following commands in separate terminals**:

**b. Run the Core Blockchain** :

    cd backend
    go run main.go

        server start at http://localhost:8080/blocks (blockchain)

**c. Run the Relay Server and Dashboard**:

    cd bridge
    go run main.go

    Expected output:
    
     server start at http://localhost:8083/events (Relay server capturing events)

     server start at http://localhost:8083/api/bridge-validation (Validation tests)

     server start at http://localhost:8083/validation (Validation results Dashboard)

     server start at http://localhost:8083/ (Dashboard)


     ( Note: This will start the Ethereum and Solana Event Listeners relaying events to the Go-Bridge )

**d. REST API testing**:


    cd core/go-bridge-server/cmd/server
    go run main.go

    With POSTMAN:

        - Send a POST request to http://localhost:8084/eth/relay

        - Send a POST request to http://localhost:8084/sol/relay

    ( Note: The Message of Transaction Post method will appear in VS Terminal )


**e. Run simple frontend for Core Blockchain Transactions** :

        cd frontend
        npm install
        
        tailwind:
            npm install -D tailwindcss postcss autoprefixer
            npx tailwindcss init

        Start the development server:
            npm run dev

        The app will be available at http://localhost:5173

        For simple dashboard design:

            http://localhost:5173/dashboard
