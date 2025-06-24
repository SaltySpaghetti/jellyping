# jellyping

jellyping is a Go-based application that integrates with Jellyfin and Telegram to provide notifications and user management features.

## Features
- Telegram bot integration for notifications and commands
- Jellyfin user import and management

## Prerequisites
- Docker & Docker Compose
- Telegram Bot Token
- Jellyfin server URL and API key

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/SaltySpaghetti/jellyping.git
   cd jellyping
   ```
2. **Copy and edit the environment variables:**
   ```sh
   cp .env.example .env
   ```
3. **Run with Docker Compose:**
   ```sh
   docker-compose up --build
   ```
   Or run locally:
   ```sh
   go run main.go
   ```

## Usage
Contact your bot directly on Telegram sending the `/username` command followed by your usename on Jellyfin. If the provided username doesn't match you'll get an error.

## Customization
| Variable           | Required | Description                        |
|--------------------|:--------:|------------------------------------|
| TELEGRAM_BOT_TOKEN |   Yes    | Telegram bot token                 |
| JELLYFIN_URL       |   Yes    | Jellyfin server URL                |
| JELLYFIN_API_KEY   |   Yes    | Jellyfin API key                   |
| PORT               |    No    | Port for the application           |
| POSTGRES_USER      |    No    | PostgreSQL username                |
| POSTGRES_PASSWORD  |    No    | PostgreSQL password                |
| POSTGRES_DB        |    No    | PostgreSQL database name           |

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.