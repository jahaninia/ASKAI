package ari

import (
	"log"
)

func HandleCallEvent(msg []byte) {
	log.Printf("Received ARI event: %s", string(msg))
	// در مراحل بعدی: پردازش تماس، STT، AI، TTS، پخش صوت
}