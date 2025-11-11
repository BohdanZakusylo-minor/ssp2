# South Park Messaging System (SSP2)


This project implements **Hexagonal Architecture (Ports and Adapters)** with the following components:

- **Go HTTP API** - RESTful service that accepts messages
- **RabbitMQ Message Broker** - Asynchronous message queue
- **Python Consumer** - Service that consumes and displays messages

### Hexagonal Architecture Structure

```
go/internal/core/
├── domain/          # Domain entities (Message)
├── ports/           # Interfaces/contracts (MessagePublisher)
├── app/             # Application services/use cases
└── adapters/
    ├── http/        # Inbound adapter (HTTP handlers)
    └── rabbitmq/    # Outbound adapter (RabbitMQ implementation)
```

**Flow:**
```
HTTP Request → HTTP Adapter → Application Service → Port Interface → RabbitMQ Adapter → RabbitMQ Queue
                                                                                          ↓
                                                                                    Python Consumer
```

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.25+ (for local development)
- Python 3.12+ (for local development)

### Running with Docker Compose

1. **Start all services:**
   ```bash
   docker compose up --build
   ```

2. **Services will be available at:**
   - Go API: http://localhost:8080
   - RabbitMQ Management UI: http://localhost:15672 (user: `user`, password: `password`)

3. **View logs:**
   ```bash
   # View all logs (Python consumer logs have timestamps and levels)
   docker compose logs -f
   
   # View only Python consumer logs (with timestamps and formatted output)
   docker compose logs -f python-consumer
   
   # View only Go service logs
   docker compose logs -f go-service
   ```

4. **Stop services:**
   ```bash
   docker compose down
   ```

## 📡 API Usage

### Send a Message

**Endpoint:** `POST /messages`

**Request:**
```bash
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{
    "author": "Cartman",
    "body": "Respect my authoritah!"
  }'
```

**Response:**
```json
{
  "status": "message received and published"
}
```

### Example Messages

```bash
# Cartman
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Cartman", "body": "Respect my authoritah!"}'

# Stan
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Stan", "body": "Oh my God, they killed Kenny!"}'

# Kyle
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Kyle", "body": "You bastards!"}'
```

## 🐍 Python Consumer

The Python consumer service listens to the `southpark_messages` queue and displays messages to the console with **enhanced logging** (timestamps, log levels, and formatted output) for better visibility in Docker.

**The consumer runs automatically in Docker Compose**, but you can also run it locally:

```bash
cd python
pip install -r requirements.txt
RABBITMQ_HOST=localhost python main.py
```

**Expected log output format:**
```
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [INFO] === South Park Messages Consumer Starting ===
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [INFO] Initializing connection to RabbitMQ...
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [INFO] Connecting to RabbitMQ at rabbit-mq:5672...
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [SUCCESS] ✓ Successfully connected to RabbitMQ
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [SUCCESS] ✓ Queue 'southpark_messages' ready
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:45] [INFO] ⏳ Waiting for messages...
python-service  | 
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE] 📨 NEW MESSAGE from 'Cartman'
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE]    Content: Respect my authoritah!
python-service  | [PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE] ────────────────────────────────────────────────────────────
```

**Note:** All Python consumer logs are prefixed with `[PYTHON-CONSUMER]` and include timestamps with log levels (`[INFO]`, `[SUCCESS]`, `[MESSAGE]`, `[ERROR]`, `[WARNING]`) to stand out in Docker logs. Use `docker compose logs -f python-consumer` to view only consumer logs.

## 📁 Project Structure

```
ssp2/
├── docker-compose.yml          # Docker orchestration
├── test-services.sh           # Script to send test messages
├── go/                         # Go API service
│   ├── Dockerfile
│   ├── go.mod
│   ├── main.go                # Application entry point
│   └── internal/core/
│       ├── domain/            # Domain entities
│       │   └── message.go
│       ├── ports/             # Port interfaces
│       │   └── message_publisher.go
│       ├── app/               # Application services
│       │   └── message_service.go
│       └── adapters/          # Infrastructure adapters
│           ├── http/          # HTTP handlers
│           │   └── message_handler.go
│           └── rabbitmq/     # RabbitMQ implementation
│               └── rabbitmq_service.go
└── python/                    # Python consumer service
    ├── Dockerfile
    ├── main.py
    ├── requirements.txt
    └── rabbit_mq_reader/
        └── consumer.py
```

## 🔧 Development

### Running Go Service Locally

```bash
cd go
go mod download
go run main.go
```

### Running Python Consumer Locally

```bash
cd python
pip install -r requirements.txt
python main.py
```

**Note:** Make sure RabbitMQ is running (via Docker Compose) before starting services locally.

## ✅ Features

- ✅ Hexagonal Architecture (Ports and Adapters)
- ✅ RESTful HTTP API
- ✅ Message validation
- ✅ Asynchronous message processing via RabbitMQ
- ✅ Python consumer with JSON parsing and enhanced logging (timestamps, levels)
- ✅ Docker Compose orchestration
- ✅ Error handling for invalid JSON and missing fields
- ✅ Automated message sending script for testing

## 📜 Message Sending Script

A bash script is provided to automatically send multiple messages to the API. This is perfect for testing the system or demonstrating the messaging flow.

### Usage

```bash
bash test-services.sh "Stan" "Oh my God, they killed Kenny!"
```

**What it does:**
- Sends 10 messages with a 1-second delay between each
- Automatically appends a message number (#1, #2, etc.) to track the sequence
- Perfect for testing the consumer and seeing messages flow through RabbitMQ

**Example output:**
The script will send 10 messages like:
- `{"author":"Cartman","body":"Respect my authoritah! #1"}`
- `{"author":"Cartman","body":"Respect my authoritah! #2"}`
- ... and so on

**Note:** Make sure the Go API service is running before executing the script.

## 🧪 Testing

### Test the API

```bash
# Valid message
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Cartman", "body": "Test message"}'

# Invalid JSON (should return 400)
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Cartman"}'

# Missing fields (should return 400)
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": ""}'
```

### Test with the Script

1. Start all services: `docker compose up --build`
2. In another terminal, watch the Python consumer logs: `docker compose logs -f python-consumer`
3. In a third terminal, run the script: `./test-services.sh "Cartman" "Respect my authoritah!"`
4. Watch formatted messages with timestamps appear in the Python consumer logs!

**Expected output in consumer logs:**
```
[PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE] 📨 NEW MESSAGE from 'Cartman'
[PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE]    Content: Respect my authoritah! #1
[PYTHON-CONSUMER] [2025-11-11 21:32:46] [MESSAGE] ────────────────────────────────────────────────────────────
[PYTHON-CONSUMER] [2025-11-11 21:32:47] [MESSAGE] 📨 NEW MESSAGE from 'Cartman'
[PYTHON-CONSUMER] [2025-11-11 21:32:47] [MESSAGE]    Content: Respect my authoritah! #2
[PYTHON-CONSUMER] [2025-11-11 21:32:47] [MESSAGE] ────────────────────────────────────────────────────────────
```
