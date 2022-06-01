package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	prefix  string
	replies = []string{
		"it is certain.",
		"it is decidedly so.",
		"without a doubt.",
		"yes definitely.",
		"you may rely on it.",
		"as I see it, yes.",
		"most likely.",
		"outlook good.",
		"yes.",
		"signs point to yes.",
		"reply hazy, try again.",
		"ask again later.",
		"better not tell you now.",
		"cannot predict now.",
		"concentrate and ask again.",
		"don't count on it.",
		"my reply is no.",
		"my sources say no.",
		"outlook not so good.",
		"very doubtful.",
	}
)

func main() {

	//Load .env file
	godotenv.Load()

	if os.Getenv("PREFIX") == "" {
		prefix = "!"
	} else {
		prefix = os.Getenv("PREFIX")
	}

	fmt.Println("Prefix: '" + prefix + "'")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If message does not start with prefix, ignore it
	if len(m.Content) < len(prefix) || m.Content[0:len(prefix)] != prefix {
		return
	}

	// If message is not in a guild, ignore it
	if m.GuildID == "" {
		return
	}

	// Convert message to lowercase
	message := m.Content[len(prefix):]
	message = strings.ToLower(message)

	// Remove all spaces
	message = strings.Replace(message, " ", "", -1)

	// Remove all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
	}

	message = reg.ReplaceAllString(message, "")

	// Convert string containg letters to Int64
	h := md5.New()
	io.WriteString(h, message)
	var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))

	// reply to the user with a random reply and notify user
	s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, "+replies[rand.Intn(len(replies))])
}
