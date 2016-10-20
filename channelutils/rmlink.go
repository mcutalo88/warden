package channelutils
import (
	"github.com/bwmarrin/discordgo"
	"regexp"
)
var Reg = regexp.MustCompile(`https?`)

/*Ok so your looking at the function names and asking WHY ARE THEY CAPITAL?
  Go needs them to be capital to be exportable.*/ 
//Check if the message content is a link
func IsLink(s *discordgo.Session, m *discordgo.MessageCreate) {
	if Reg.MatchString(m.Content) {
		RemoveLink(s,m)
		WarnUser(s,m)
	}
}
// Sends the user a private message saying fuck off man
func WarnUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	 createChat,_:=s.UserChannelCreate(m.Author.ID)
	 //channel,_:=s.Channel(m.ChannelID)
	 message:= "Do not paste link outside of pics_vids channel!!!! This is a warning!"
	 _, _ = s.ChannelMessageSend(createChat.ID,message)

}
// This pillages the link from chat get wrecked son
func RemoveLink(s *discordgo.Session, m *discordgo.MessageCreate) {
	//THIS IS FOR TESTING REMOVE WHEN DONE!!!!
	//time.Sleep(5 *time.Second)
	s.ChannelMessageDelete(m.ChannelID,m.ID)
}
