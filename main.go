package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Email    string
	Password string
	Token    string
	BotID    string
)
var Reg = regexp.MustCompile(`https?`)

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New(Email, Password, Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	islink:= isLink(m)

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	if islink {
		removeLink(s,m)
		//warnUser
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, strconv.FormatBool(islink))
}

//Check if the message content is a link
func isLink(m *discordgo.MessageCreate) bool {
	if Reg.MatchString(m.Content) {
		return true
	} else {
		return false
	}
}

// func warnUser() {
//
// }
//
func removeLink(s *discordgo.Session, m *discordgo.MessageCreate) {
	//time.Sleep(5 *time.Second) THIS IS FOR TESTING REMOVE WHEN DONE!!!!
	s.ChannelMessageDelete(m.ChannelID,m.ID)
}
