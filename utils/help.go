/*
 * Help output for warden
 * Will handle all output related to -h, with later support
 * for specific command help.
 *
 * eg: ./warden -h # prints all commands
 * eg: ./warden ban -h # would print out help for ban command
 *
 */
package utils

import (
  "github.com/bwmarrin/discordgo"
)

// Render config
func RenderConfig(s *discordgo.Session, m *discordgo.MessageCreate) {
  _, _ = s.ChannelMessageSend(m.ChannelID, "Choice")
}
