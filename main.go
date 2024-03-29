package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

// Variables used for command line parameters.

var (
	Token, IexToken string
)

type DogResp struct {
	Success string
	Message string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("DISCORD_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	IexToken = os.Getenv("IEX_TOKEN")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	routing(s, m)
}

// Need much better scaling and maintainability of request routing
func routing(s *discordgo.Session, m *discordgo.MessageCreate) {
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "I'm backkkkkkk!")
	}
	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	if m.Content == "dog" {
		s.ChannelMessageSend(m.ChannelID, getDog())
	}
	if m.Content == "cat" {
		s.ChannelMessageSend(m.ChannelID, getCat())
	}

	stockMatcher, _ := regexp.Compile("\\$([A-Z]|[a-z]){1,4}")
	if stockMatcher.MatchString(m.Content) {
		stockFinder, _ := regexp.Compile("([A-Z]|[a-z]){1,4}")
		stock := stockFinder.FindString(m.Content)
		s.ChannelMessageSend(m.ChannelID, "Attempting to look up stock "+stock)

		quote := quote(IexToken, stock)
		fmt.Println(quote)
		s.ChannelMessageSend(m.ChannelID, "$"+quote)
	}
}
