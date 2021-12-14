// Package api provides a Vercel Serverless Function that uses disgoslash
// to serve Discord Slash Command Interactions.
//
// https://vercel.com/docs/serverless-functions/supported-languages#go
//
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
package api

import (
	"net/http"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
)

// GuildIDs holds the list of Guild (server) IDs you would like to register
// a slash command to.
var GuildIDs = []string{"850618557523689502"}

// Global indicates whether or not a slash command should be registered globally
// across all Guilds the bot has access to.
var Global = true

// Credentials holds your Discord Application's secret credentials.
//
// WARNING - Do not track these secrets in version control.
//
// https://discord.com/developers/applications
var Credentials = &discord.Credentials{
	PublicKey: "93abf1cf47ae528fd19ce6eae4fe8bcc61b816892c1b0a7791742aff25134912", // Your Discord Application's Public Key
	ClientID:  "864136726683713536",  // Your Discord Application's Client ID
	Token:     "ODY0MTM2NzI2NjgzNzEzNTM2.YOxECg.1xhEMxv5zQyG2pApVs8rwfTKvW8",      // Your Discord Application's Bot's Token
}

var command = &discord.ApplicationCommand{
	Name:        "hello",
	Description: "Says hello to the user",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Name",
			Description: "Enter your name",
			Required:    true,
		},
	},
	DefaultPermission: true,
}

// hello is where the code of the slash command lives
func hello(request *discord.InteractionRequest) *discord.InteractionResponse {
	// Your custom code goes here!
	name, _ := request.Data.Options[0].StringValue()
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: "Hello " + name + "!",
		},
	}
}

var slashCommand = disgoslash.NewSlashCommand(command, hello, Global, GuildIDs)

// SlashCommandsMap is exported for use with the sync package which will
// register the slash command on Discord.
var SlashCommandMap = disgoslash.NewSlashCommandMap(slashCommand)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	handler := &disgoslash.Handler{SlashCommandMap: SlashCommandMap, Creds: Credentials}
	handler.Handle(w, r)
}