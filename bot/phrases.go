package bot

// Phrases used by bot, all in one place
func say(key string) string {
	phrases := make(map[string]string)
	phrases["greeting"] = "Hi! 👋\nI'm a %s and together we can create some nice stickers!\n\nTo start making a sticker like the one shown below please send me a photo (jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max)"
	phrases["botName"] = "StickerBot"
	phrases["perfect_one"] = "This one is perfect. I will save it in Templates in case you'll want to reuse it."
	phrases["replace_photo_or_provide_text"] = "If you'ld like to replace the image please click \"Replace image\". If it's Ok, please send me the text for your sticker (60 symbols max)"
	phrases["replace_image"] = "Replace image"
	phrases["resend_image"] = "Ok, please send me another one. Remember that it should be in jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max."

	return phrases[key]
}
