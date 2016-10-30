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

const CODE_BLOCK string = "```"

func PrintHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
  _, _ = s.ChannelMessageSend(m.ChannelID, Help())
}

func Help() string {
  return CODE_BLOCK +
    "./warden \n" +
    "\t-h :\tprint help menu \n" +
    CODE_BLOCK
}
