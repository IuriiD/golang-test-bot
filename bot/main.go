package bot

import (
	"fmt"
	"os"

	gbl "github.com/calebhiebert/gobbl"
	fb "github.com/calebhiebert/gobbl/messenger"
)

// Setup will setup the bot instance
func Setup() (*gbl.Bot, *fb.MessengerIntegration, error) {

	// Create a new bot
	gobblr := gbl.New()

	// Use this middleware to make sure the bot responds to requests
	gobblr.Use(gbl.ResponderMiddleware())

	// Get user ID
	gobblr.Use(gbl.UserExtractionMiddleware())

	// Extract what the user has said from the context
	gobblr.Use(gbl.RequestExtractionMiddleware())

	// Show that user's message was seen by the bot
	gobblr.Use(fb.MarkSeenMiddleware())

	// Router setup
	textRouter := gbl.TextRouter()
	ictxRouter := bctx.ContextIntentRouter()
	intentRouter := gbl.IntentRouter()
	customRouter := gbl.CustomRouter()

	// Custom simple echo middleware
	gobblr.Use(func(c *gbl.Context) {
		// Create response object in context object
		r := fb.CreateResponse(c)

		// Add text response
		fmt.Printf("c.User.FirstName %s\n", c.User.FirstName)
		fmt.Printf("c.User %s\n", c.User)
		fmt.Printf("User said: %s\n", c.Request.Text)
		r.Text(fmt.Sprintf(say("greeting"), c.User.FirstName, say("botName")))
	})

	// FACEBOOK MESSENGER SETUP
	mapi := fb.CreateMessengerAPI(os.Getenv("PAGE_ACCESS_TOKEN"))
	messengerIntegration := fb.MessengerIntegration{
		API:            mapi,
		Bot:            gobblr,
		VerifyToken:    os.Getenv("FB_VERIFY_TOKEN"),
		DevMode:        true,
		EnableRecovery: false,
		Always200:      true,
	}

	return gobblr, &messengerIntegration, nil
}

/*
	// Create a new bot
	gobblr := gbl.New()

	// Use this middleware to make sure the bot responds to requests
	gobblr.Use(gbl.ResponderMiddleware())

	// Use the request extraction middleware
	// to extract what the user has said from the context
	gobblr.Use(gbl.RequestExtractionMiddleware())

	// Add a simple middleware that will send an echo response
	gobblr.Use(func(c *gbl.Context) {

		// When using the console integration, the context R (response) object
		// we need to cast it so we can use it's functions
		basicResponse := c.R.(*gbl.BasicResponse)

		// Add a text message to the output
		basicResponse.Text(fmt.Sprintf("Echo: %s", c.Request.Text))
		basicResponse.Text("I am another line of text")
	})

	// Create a new console integration
	ci := gbl.ConsoleIntegration{}

	// Start listening to the console input
	ci.Listen(gobblr)*/

/*
func main() {
	// Create new bot
	gobblr := gbl.New()

	// Create Messenger API
	mapi := fb.CreateMessengerAPI(os.Getenv("PAGE_ACCESS_TOKEN"))

	// Create the integration
	fbmint := fb.MessengerIntegration{
		API: mapi,
		VerifyToken: os.Getenv("FB_VERIFY_TOKEN")
	}

	// Middleware to extract user from FB Messenger webhooks
	gobblr.Use(gbl.UserExtractionMiddleware())

	// Middleware to extract user's text requests
	gobblr.Use(gbl.RequestExtractionMiddleware())

	// Middleware to display that messages were seen by bot
	gobblr.Use(gbl.MarkSeenMiddleware())

	// Custom simple echo middleware
	gobblr.Use(func(c *gbl.Context) {
		// Create response object in context object
		r := fb.CreateResponse(c)

		// Add text response
		r.Text(fmt.Sprintf("Echo: %s", c.Request.Text))
	})
}

*/
