# Ethereum Event Listener

This project is a standalone Go module that listens to Ethereum smart contract events, specifically focusing on ERC-20 Transfer events. It outputs logs to a localhost server for real-time verification.

## Project Structure

```
eth-event-listener
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── listener
│   │   └── listener.go  # Implementation of the event listener
│   └── types
│       └── event.go     # Data structures related to events
├── go.mod                # Go module configuration
└── README.md             # Project documentation
```

## Setup Instructions

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine.

   ```bash
   git clone <repository-url>
   cd eth-event-listener
   ```

3. **Install Dependencies**: Navigate to the project directory and run the following command to install the required dependencies.

   ```bash
   go mod tidy
   ```

## Usage

1. **Configure Ethereum Client**: Update the Ethereum client connection details in `cmd/main.go` to point to your desired Ethereum node (e.g., Infura, Alchemy, or a local node).

2. **Run the Application**: Start the application by executing the following command:

   ```bash
   go run cmd/main.go
   ```

3. **View Logs**: Open your web browser and navigate to `http://localhost:8080` to view the real-time logs of ERC-20 Transfer events.

## Event Listener Functionality

The application listens for ERC-20 Transfer events emitted by smart contracts on the Ethereum blockchain. It captures these events and outputs relevant information to a local server for easy monitoring and verification.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.