package main

import (
	"github.com/picatz/go-av"
)

func main() {
	text := "Hello world! This is Arthur."

	av.TextToSpeechWithVoice(text, "com.apple.ttsbundle.siri_Arthur_en-GB_premium")
}
