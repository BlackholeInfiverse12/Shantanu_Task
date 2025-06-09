# Go Bridge Server

## Overview
The Go Bridge Server is a service designed to facilitate the relaying of messages between Ethereum and Solana blockchains. It provides HTTP endpoints for processing relay requests and integrates with a blockchain core for message handling.

## Project Structure
```
go-bridge-server
├── internal
│   ├── handler.go        # HTTP handlers for processing relay messages
│   └── server.go         # Main server setup and routing logic
├── core
│   └── blockchain
│       └── blockchain.go  # Blockchain logic for message handling
├── go.mod                # Module definition file
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Installation
To get started with the Go Bridge Server, clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd go-bridge-server
```

Ensure you have Go installed on your machine. You can download it from the official Go website.

## Dependencies
Run the following command to download the necessary dependencies:

```bash
go mod tidy
```

## Running the Server
To start the server, execute the following command:

```bash
go run internal/server.go
```

The server will start and listen for incoming relay requests.

## API Endpoints
### Ethereum Relay
- **Endpoint:** `/eth/relay`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
      "data": "your_message_here"
  }
  ```
- **Response:** 
  - Status: 202 Accepted
  - Body: 
  ```json
  {
      "status": "success",
      "message": "Ethereum relay processed"
  }
  ```

### Solana Relay
- **Endpoint:** `/sol/relay`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
      "data": "your_message_here"
  }
  ```
- **Response:** 
  - Status: 202 Accepted
  - Body: 
  ```json
  {
      "status": "success",
      "message": "Solana relay processed"
  }
  ```

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.