package speech

import (
	"log"
)

func TranscribeAudio(filePath string) (string, error) {
	log.Printf("Transcribing audio file: %s", filePath)
	// اینجا باید به سرویس STT مثل Google یا Whisper وصل بشی
	// فعلاً خروجی تستی:
	return "سلام، قبض من مشکل دارد", nil
}