package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var discordToken = os.Getenv("DISCORD_TOKEN")
var commandRegex, _ = regexp.Compile(`^!c +\d+ +\d+ +\d+ +\d+$`)

func parseCommandArgs(command string) (int, int, int, int, error) {
	if commandRegex.MatchString(command) {
		command = trimDuplicatedWhiteSpaces(command)
		args := strings.Split(command[3:], " ")
		if len(args) == 4 {
			// initial resources
			ir, _ := strconv.Atoi(args[0])
			// number of resources required per product
			nr, _ := strconv.Atoi(args[1])
			// crafting price per product
			cp, _ := strconv.Atoi(args[2])
			// resource return rate - per-mille
			rrr, _ := strconv.Atoi(args[3])

			return ir, nr, cp, rrr, nil
		}
	}

	return 0, 0, 0, 0, errors.New("Invalid command")
}

func main() {
	log.Println("CRAFETA DISCORD BOT STARTING...")

	bot, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Panic("[ERROR]: creating Discord session,", err)
		return
	}
	bot.AddHandler(messageCreate)
	err = bot.Open()
	if err != nil {
		log.Panic("[ERROR]: opening connection,", err)
		return
	}
	defer bot.Close()

	log.Println("CRAFETA DISCORD BOT STARTED")

	sc := make(chan struct{})
	<-sc

	log.Println("CRAFETA DISCORD BOT STOPED")
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		// ignore if the message came from this bot itself
		return
	}
	if strings.HasPrefix(msg.Content, "!c") {
		command := msg.Content
		initialResources, numberOfResourcesPerProduct, craftingPricePerProduct, resourceReturnRate, err := parseCommandArgs(command)
		if err == nil {
			totalProducts, totalPrice, resourcesRemaining := getTotalCraftedProducts(initialResources, numberOfResourcesPerProduct, craftingPricePerProduct, resourceReturnRate)

			session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("TOTAL CRAFTED PRODUCTS: %d\nTOTAL CRAFTING COST: %d\nRESOURCES REMAINING: %d", totalProducts, totalPrice, resourcesRemaining))
			return
		}

		// wrong command format, shown command usage
		session.ChannelMessageSend(msg.ChannelID, "Hi!")
	}
}
