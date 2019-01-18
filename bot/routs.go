package bot

import (
	"fmt"

	gbl "github.com/calebhiebert/gobbl"
	fb "github.com/calebhiebert/gobbl/messenger"
)

// Function getStarted is called when user first interacts with the bot
// and also in response to greeting phrases
func getStarted(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(fmt.Sprintf(say("greeting"), say("botName")))
}

// Function provideText is called when user provided a photo
func provideText(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(fmt.Sprintf(say("perfect_one")))
	r.Text(fmt.Sprintf(say("replace_photo_or_provide_text")))
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
