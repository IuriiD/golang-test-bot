package bot

// Phrases used by bot, all in one place
func say(key string) string {
	phrases := make(map[string]string)
	phrases["greeting"] = "Hi! 👋\nI'm a %s and together we can create some nice stickers!\n\nTo start making a sticker like the one shown below please send me a photo (jpeg or png format, not smaller than 150x150px and not bigger than 3500x2400px, 5Mb max)"
	phrases["botName"] = "StickerBot"

	return phrases[key]
}
