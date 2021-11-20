package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/themobilecoder/ocm-meta/meta"
)

type Config struct {
	DISCORD_API_KEY string `envconfig:"DISCORD_API_KEY" required:"true"`
}

var monkeys []meta.Monkey

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "meta",
		Description: "Show meta traits of OCM. Enter valid OCM ID (1 to 10000)",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "OCM Id (1 - 10000)",
				Required:    true,
			},
		},
	},
}
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"meta": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		margs := []interface{}{
			i.ApplicationCommandData().Options[0].StringValue(),
		}
		ocmId, err := strconv.Atoi(margs[0].(string))
		if err != nil || ocmId < 1 || ocmId > 10000 {
			sendInvalidCommandMessage(s, i)
			return
		}
		idx := ocmId - 1
		if len(monkeys) < idx {
			sendInvalidCommandMessage(s, i)
			return
		}
		monkey := monkeys[idx]
		idString := strconv.Itoa(ocmId)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "OCM #" + idString,
						URL:   "https://opensea.io/assets/0x960b7a6bcd451c9968473f7bbfd9be826efd549a/" + idString,
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "https://d3q7x2s6555pey.cloudfront.net/png/" + idString + ".png",
						},
						Type: "rich",
						Footer: &discordgo.MessageEmbedFooter{
							Text:    "ocmetabot",
							IconURL: "https://d3q7x2s6555pey.cloudfront.net/png/4642.png",
						},
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "Hat",
								Value:  monkey.Hat,
								Inline: true,
							},
							{
								Name:   "Fur",
								Value:  monkey.Fur,
								Inline: true,
							},

							{
								Name:   "Clothes",
								Value:  monkey.Clothes,
								Inline: true,
							},
							{
								Name:   "Mouth",
								Value:  monkey.Mouth,
								Inline: true,
							},
							{
								Name:   "Eyes",
								Value:  monkey.Eyes,
								Inline: true,
							},
							{
								Name:   "Earring",
								Value:  monkey.Earring,
								Inline: true,
							},
							{
								Name:   "Background",
								Value:  monkey.Background,
								Inline: false,
							},
							{
								Name:   "Trait Count",
								Value:  monkey.Trait_count,
								Inline: true,
							},
							{
								Name:   "Color Match",
								Value:  monkey.Color_match,
								Inline: true,
							},
							{
								Name:   "Mouth Match",
								Value:  monkey.Mouth_match,
								Inline: true,
							},
							{
								Name:   "Zeroes",
								Value:  monkey.Zeros,
								Inline: true,
							},
							{
								Name:   "Nips",
								Value:  monkey.Nips,
								Inline: true,
							},
							{
								Name:   "Poker Hands",
								Value:  flattenOrNone(monkey.Poker_hands),
								Inline: true,
							},
							{
								Name:   "Twins",
								Value:  flattenOrNone(monkey.Xplets),
								Inline: true,
							},
						},
					},
				},
			},
		})
	},
}

func main() {
	var cfg Config
	err := envconfig.Process("METABOT", &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s, err := discordgo.New("Bot " + cfg.DISCORD_API_KEY)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	monkeys = meta.GetOnChainMonkeys()
	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")
}

func flattenOrNone(ss []string) string {
	if len(ss) == 0 {
		return "None"
	}
	var sb = strings.Builder{}
	for i, s := range ss {
		if i < len(ss)-1 {
			sb.WriteString(s + ", ")

		} else {
			sb.WriteString(s)
		}
	}
	return sb.String()
}

func sendInvalidCommandMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: "Invalid command. Make sure id is within 1 to 10000",
				},
			},
		},
	})
}
