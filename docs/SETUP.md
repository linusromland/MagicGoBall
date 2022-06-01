# Run and Build

## Dependencies

-   Go
-   DiscordGo
-   godotenv

## Install dependencies

```bash
# Navigate to the bot directory
cd bot

# Install dependencies
go get
```

## Environment variables

Create a `.env` file in the `bot` directory of the project.

Add the following lines to the `.env` file:

```bash
DISCORD_BOT_TOKEN=TOKEN
```

Replacing `TOKEN` with your bot's token.

## Start the bot

```bash
go run main.go
```
