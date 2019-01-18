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
	intentRouter := gbl.IntentRouter()
	customRouter := gbl.CustomRouter()

	// Adding router middleware to the bot
	gobblr.Use(textRouter.Middleware())
	gobblr.Use(customRouter.Middleware())
	gobblr.Use(intentRouter.Middleware())

	// Route setup
	textRouter.Text(TCGetStarted, getStarted)

	// Custom simple echo middleware
	gobblr.Use(func(c *gbl.Context) {
		// Create response object in context object
		r := fb.CreateResponse(c)

		// Add text response
		fmt.Printf("c.User.FirstName %s\n", c.User.FirstName)
		fmt.Printf("c.User %s\n", c.User)
		fmt.Printf("User said: %s\n", c.Request.Text)
		r.Text(fmt.Sprintf(say("greeting"), say("botName")))
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
