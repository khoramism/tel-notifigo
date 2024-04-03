# Telegram Message Sender

This Go application serves as a simple HTTP server that allows you to send messages to a specified Telegram chat or group through a straightforward HTTP POST request. Utilizing the Telegram Bot API, it offers an easy way to integrate message notifications or alerts into your systems.

## Features

- Send messages to Telegram chats or groups using HTTP POST requests.
- Easy configuration via environment variables.
- Basic error handling and logging.

## Requirements

- Go (version 1.11 or newer)
- A Telegram bot token and chat ID

## Setup

### Getting Started

1. **Create a Telegram Bot**: Follow the [official Telegram instructions](https://core.telegram.org/bots#6-botfather) to create a bot and obtain your bot token.

2. **Find Your Chat ID**: You need the chat ID of the group or chat where messages will be sent. [Here's how you can find it](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id).

### Installation

Clone this repository and navigate into the project directory:

```bash
git clone https://github.com/khoramism/tel-notifigo
cd tel-notifigo
```

### Configuration

Set the necessary environment variables:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot's token.
- `DEFAULT_CHAT_ID`: The default chat ID where messages will be sent if none is specified in the request.

You can set these variables in your shell or use a `.env` file in your project directory.

## Usage
Your README provides a good overview of your Go application, detailing its purpose, setup, and how to use it. Below, I've added a section on how to use your application with Docker, which you can incorporate into your existing README. This new section covers building a Docker image from your Dockerfile and running your application using Docker Compose, based on the Dockerfile and `docker-compose.yml` you've provided.

---

### Docker Usage

This section describes how to use Docker to build and run the Telegram Message Sender application. Docker simplifies deployment and ensures consistency across different environments.

#### Building the Docker Image
##### NOTE: (image has already been pushed to dockerhub if that way suits you better: `khoramism/tel-notifigo`
Before running the application with Docker, you need to build the Docker image. You can do this by running the following command in the project directory:

```bash
docker build -t khoramism/tel-notifigo .
```

This command builds a Docker image named `khoramism/tel-notifigo` based on the instructions in your Dockerfile.

#### Running with Docker Compose

To run the application using Docker Compose, ensure you have `docker-compose` installed and then use the following command:

```bash
docker compose up
```

This command starts the application as defined in your `docker-compose.yml` file. The service is named `telnotif`, and it exposes the application on port `8787` of your localhost. It also uses an environment file, `.env`, where you should specify your `TELEGRAM_BOT_TOKEN` and `DEFAULT_CHAT_ID` among any other environment variables you might need.

#### Docker Compose File Explanation

Your `docker-compose.yml` version is set to `'3'`, which specifies the version of the Docker Compose file format. The `services` section defines the services your application uses (in this case, just one service named `telnotif`). The `image` specifies the Docker image to use for the container, and `ports` maps port 8787 inside the container to port 8787 on your host, allowing you to access the application via `localhost:8787`. The `env_file` specifies the location of your environment file, which Docker Compose will use to set environment variables inside the container.

By following these steps, you can easily build and run your Telegram Message Sender application in a Dockerized environment, making it more portable and easier to deploy across different systems.

---

Start the server using bash:

```bash
go run .
```

The server will listen on port `8787`. To send a message, make an HTTP POST request to `http://localhost:8787/` with the following JSON payload:

```json
{
  "text": "Hello, World!",
  "chat_id": "[optional_chat_id]"
}
```

If `chat_id` is omitted, the message will be sent to the `DEFAULT_CHAT_ID` defined in your environment variables.

### Example CURL Command

```bash
curl -X POST http://localhost:8787/ -H "Content-Type: application/json" -d '{"text":"Hello, World!", "chat_id":"YOUR_CHAT_ID_HERE"}'
```

