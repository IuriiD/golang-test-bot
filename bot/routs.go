package bot

import (
	"fmt"

	gbl "github.com/calebhiebert/gobbl"
	fb "github.com/calebhiebert/gobbl/messenger"
)

// HGetStarted is called when user first interacts with the bot
// and also in response to greeting phrases
func getStarted(c *gbl.Context) {
	r := fb.CreateResponse(c)
	r.Text(fmt.Sprintf(say("greeting"), say("botName")))
}
