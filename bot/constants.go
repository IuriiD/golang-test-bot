package bot

// CONTEXT VARIABLES
// Name of context we will store our dialog state
var CBotState = "BOT_STATE"

// CPromptedPhoto value for context "BOT_STATE" is set after user was asked to upload a photo
var CPromptedPhoto = "PHOTO_PROMPT_FOLLOWUP"

// CPromptedText value for context "BOT_STATE" is set after user was asked to send text for the sticker
var CPromptedText = "TEXT_PROMPT_FOLLOWUP"

// TEXT CONSTANTS
// TCGetStarted will show the user the welcome message
var TCGetStarted = "GETTING_STARTED"

// TCReplaceImage will prompt user to send an image
var TCReplaceImage = "REPLACE_IMAGE"

// TCImageReceived, temp
var TCImageReceived = "IMAGE_RECEIVED"

// TCShowFlag, temp
var TCShowFlag = "FLAG"

// OTHER CONSTANTS
// User's cusom greetings and restarting commands
var OCGreetings = []string{"hello", "hi", "hey", "aloha", "greetings", "good day", "good morning", "good evening", "restart", "start over", "reload", "begin", "start afresh", "start", "get started", "menu", "show menu", "go back", "do it again", "default", "getting started", "main menu"}
