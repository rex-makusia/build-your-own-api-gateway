# build-your-own-api-gateway

## Overview
This project is a basic HTTP API gateway that routes requests to various backend services. It is designed to be lightweight and easily extensible.

## Features
- HTTP request handling
- Proxying requests to backend services
- Middleware for authentication
- Configuration management
- Docker support for containerization

## Project Structure
```
build-your-own-api-gateway
├── cmd
│   └── gateway
│       └── main.go          # Entry point of the application
├── internal
│   ├── gateway
│   │   ├── handler.go       # HTTP handlers for the API gateway
│   │   ├── proxy.go         # Proxy logic for forwarding requests
│   │   └── server.go        # HTTP server setup
│   └── middleware
│       └── auth.go          # Authentication middleware
├── pkg
│   └── config
│       └── config.go        # Configuration management
├── api
│   └── routes.go            # API route definitions
├── configs
│   └── config.yaml          # Configuration settings in YAML format
├── scripts
│   ├── build.sh             # Script to build the application
│   └── run.sh               # Script to run the application
├── Dockerfile                # Docker image build instructions
├── Makefile                  # Build commands and tasks
├── go.mod                   # Go module definition
└── README.md                 # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd build-your-own-api-gateway
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   ./scripts/build.sh
   ```

4. Run the application:
   ```
   ./scripts/run.sh
   ```

## Usage
Once the application is running, you can send HTTP requests to the API gateway, which will route them to the appropriate backend services based on the defined routes.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.