package config

import (
    "fmt"
    "strconv"
    "strings"
    "os"
    "github.com/BurntSushi/toml"
    "github.com/bwmarrin/discordgo"
)

type Config struct {
    Channels         []string
    Message          string
    NumberOfWarnings int
    BanTextChannels	 bool
    BanDuration 	 int
    Ip               string
    Database         string
}

var GlobalConfig Config

// Reads info from config file
func ReadConfig(conf string) {
    var BotConfig Config
    if _, err := toml.DecodeFile(conf, &BotConfig); err != nil {
		fmt.Println(err)
	}

  GlobalConfig = BotConfig
}

func Get() Config {
  return GlobalConfig
}

//This will build default message if Message is commented out in file
func UseDefaultMess(s *discordgo.Session) {
    //might want to add a strip just in case
    if(GlobalConfig.Message == "") {
        var buildMess string
        var warnings = strconv.Itoa(GlobalConfig.NumberOfWarnings)
        var formatMess string
        for _,channelID := range GlobalConfig.Channels {
            channel,err:=s.Channel(channelID)
            if(err!=nil) {
                fmt.Println("Please check Channels ids in botconfig.toml. \nERROR: ",err)
                os.Exit(1)
            }

            buildMess += channel.Name+", "
            formatMess = strings.TrimRight(buildMess,", ")
            GlobalConfig.Message = "Do not paste links in `[ "+formatMess+" ]` chat!!!! You will be banned if you do this "+warnings+" more times!"
        }
    }
}
