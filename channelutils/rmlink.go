package channelutils
import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"Jeffs-bot/config"
)
var Reg = regexp.MustCompile(`https?`)

/*Ok so your looking at the function names and asking WHY ARE THEY CAPITAL?
  Go needs them to be capital to be exportable.*/
//Check if the message content is a link
func IsLink(s *discordgo.Session, m *discordgo.MessageCreate) {
	if Reg.MatchString(m.Content) {
		RemoveLink(s,m)
	}
}

// Sends the user a private message saying fuck off man
func WarnUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	 createChat,_:=s.UserChannelCreate(m.Author.ID)
	 _, _ = s.ChannelMessageSend(createChat.ID,config.GetMessage())
}

// This pillages the link from chat get wrecked son
func RemoveLink(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	var channels = config.GetChannelIDs()
	for _,channelID := range channels {
		if(m.ChannelID == channelID) {
			s.ChannelMessageDelete(m.ChannelID,m.ID)
			WarnUser(s,m)
			return true
		}
	}
	return false
}
