package bot

// Phrases used by bot, all in one place
func say(key string) string {
	phrases := make(map[string]string)
	phrases["greeting_v1"] = "Hi! 👋 I'm a %s and together we can create some nice stickers!"
	phrases["greeting_v2"] = "Hi! 👋 I'm a %s, here to make custom stickers"
	phrases["greeting_v3"] = "Hi! 👋 I'm a %s, chatbot for making custom stickers"
	phrases["to_start_send_image"] = "To start making a sticker like the one shown below please send me a photo (jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max)"
	phrases["botName"] = "StickerBot"
	phrases["perfect_one"] = "This one is perfect. I will save it in Templates in case you'll want to reuse it."
	phrases["replace_photo_or_provide_text"] = "If you'ld like to replace the image please click \"Replace image\". If it's Ok, please send me the text for your sticker (%s symbols max)."
	phrases["replace_image"] = "Replace image"
	phrases["resend_image"] = "Ok, please send me another one. Remember that it should be in jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max."
	phrases["bad_image"] = "Sorry but this won't work. Please send me an image in jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max."
	phrases["bad_phrase"] = "Sorry but your phrase should be longer than %s symbols. Please try again."
	phrases["didnt_get_that"] = "Sorry but I didn't get that."

	return phrases[key]
}
