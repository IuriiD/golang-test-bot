package bot

import (
	"os"
	s "strings"

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
	//ictxRouter := bctx.ContextIntentRouter()
	//intentRouter := gbl.IntentRouter()
	//customRouter := gbl.CustomRouter()

	// Adding router middleware to the bot
	gobblr.Use(textRouter.Middleware())
	//gobblr.Use(ictxRouter.Middleware())
	//gobblr.Use(customRouter.Middleware())
	//gobblr.Use(intentRouter.Middleware())

	// Routes setup
	// We'll be using Flags to store dialog state at the moment >> store to DB
	// Response to GETTING_STARTED button payload
	textRouter.Text(TCGetStarted, getStarted)

	// Response to greeting and restarting messages
	gobblr.Use(func(c *gbl.Context) {
		userSaid := s.ToLower(c.Request.Text)
		if contains(OCGreetings, userSaid) {
			getStarted(c)
		}
	})

	// Response to IMAGE_RECEIVED payload, temp
	textRouter.Text(TCImageReceived, provideText)

	// Response to REPLACE_IMAGE payload
	textRouter.Text(TCReplaceImage, replaceImage)

	// Default fallback response
	//gobblr.Use(defaultFallback)

	// Displays value in "botState" flag, temp
	textRouter.Text(TCShowFlag, showFlag)

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
