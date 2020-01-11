<img align="right" src="http://bwmarrin.github.io/discordgo/img/discordgo.png">

## Beginning of Discord bot for stocks

### Install dependencies

From within the project root, run the following to install all of the go gettable dependencies.

```sh
go get -d ./
```

### Build

From within the project root, run the below command to compile the source code into an executable.

```sh
go build
```

### Usage

This example uses bot tokens for authentication only. While user/password is
supported by DiscordGo, it is not recommended for bots.

```
./stonks --help
Usage of ./stonks:
  -t string
        Bot Token
```

The below example shows how to start the bot

```sh
./stonks -t YOUR_BOT_TOKEN
Bot is now running.  Press CTRL-C to exit.
```

### Add bot to Discord server

Navigate to:

```
https://discordapp.com/oauth2/authorize?client_id=%Client_ID%&scope=bot&permissions=0
```

### Bot commands

Currently getting everything set up still.

```
- ping
- pong
- cat
- dog
```