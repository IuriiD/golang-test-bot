package bot

import (
	"fmt"
	"os"

	gbl "github.com/calebhiebert/gobbl"
	bctx "github.com/calebhiebert/gobbl/context"
	fb "github.com/calebhiebert/gobbl/messenger"
)

// Function getStarted is called when user first interacts with the bot
// and also in response to greeting and restarting phrases. Sets bot's
// state to "AWAITING_IMAGE"
func getStarted(c *gbl.Context) {
	r := fb.CreateResponse(c)

	phrase1 := fmt.Sprintf(say("greeting_v1"), say("botName"))
	phrase2 := fmt.Sprintf(say("greeting_v2"), say("botName"))
	phrase3 := fmt.Sprintf(say("greeting_v3"), say("botName"))
	r.RandomText(phrase1, phrase2, phrase3)

	r.Text(fmt.Sprintf(say("to_start_send_image")))

	r.AttachmentByID("image", os.Getenv("EXAMPLE_STICKER_ID"))
	fmt.Printf("Setting context %s with lifespan %d\n", CPromptedPhoto, 1)

	// Create context and store current dialog step (awaiting image from user)
	bctx.Add(c, CBotState, 2)
	//bctx.Set(c, CBotState, "next", CPromptedPhoto)
}

// Function provideText is called when user provided a photo
func provideText(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(fmt.Sprintf(say("perfect_one")))
	r.Text(fmt.Sprintf(say("replace_photo_or_provide_text"), os.Getenv("MAX_STICKER_PHRASE_LENGTH")))
	r.QR(
		fb.QRText(say("replace_image"), TCReplaceImage),
	)
}

// Function replaceImage is called when user has uploaded image but
// then clicke QR button "Replace image"
func replaceImage(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(fmt.Sprintf(say("resend_image")))
}

// Default fallback response function. Is called in case no other functions were triggered
func defaultFallback(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(say("didnt_get_that"))
	r.Text(say("to_start_send_image"))
	r.AttachmentByID("image", os.Getenv("EXAMPLE_STICKER_ID"))
}

func showFlag(c *gbl.Context) {
	currState := "Context not found"
	if bctx.Get(c, CBotState, "next") != "" {
		currState = bctx.Get(c, CBotState, "next")
	}
	r := fb.CreateResponse(c)
	r.Text(currState)
}
