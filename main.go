package main

import (
	"fmt"
	"net/http"
	"os"

	fb "github.com/calebhiebert/gobbl/messenger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/iuriid/stickerbot/bot"
)

func main() {
	// Getting credentials
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to read .env file")
	}

	// Bot setup
	_, fbmint, err := bot.Setup()
	if err != nil {
		panic(err)
	}

	// Gin setup
	r := gin.Default()

	// Facebook webhook verification
	r.GET("/webhook", func(c *gin.Context) {
		mode := c.Query("hub.mode")
		token := c.Query("hub.verify_token")
		challenge := c.Query("hub.challenge")

		fmt.Println(os.Getenv("FB_VERIFY_TOKEN"))

		if mode == "subscribe" && token == os.Getenv("FB_VERIFY_TOKEN") {
			c.String(http.StatusOK, challenge)
		} else {
			c.AbortWithStatus(401)
		}
	})

	// For checking purposes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to StickerBot!")
	})

	// Facebook webhook post queries processing
	r.POST("/webhook", func(c *gin.Context) {
		var webhookRequest fb.WebhookRequest

		err := c.ShouldBindJSON(&webhookRequest)
		if err != nil {
			fmt.Println("WEBHOOK PARSE ERR:", err)
			c.JSON(500, gin.H{"error": "Invalid json"})
		} else {
			fbmint.ProcessWebhookRequest(&webhookRequest)
			c.JSON(200, gin.H{"o": "k"})
		}
	})

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	r.Run("0.0.0.0:" + port)
}
