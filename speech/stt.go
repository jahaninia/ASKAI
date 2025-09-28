package speech

import (
	"log"
)

// func TranscribeAudio(filePath string) (string, error) {
// 	log.Printf("Transcribing audio file: %s", filePath)
// 	// اینجا باید به سرویس STT مثل Google یا Whisper وصل بشی
// 	// فعلاً خروجی تستی:
// 	return "سلام، قبض من مشکل دارد", nil
// }

// use Whisper 
func TranscribeAudio(filePath string) (string, error) {
	cmd := exec.Command("whisper", filePath, "--model", "base", "--language", "fa", "--output_format", "txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Whisper error: %v", err)
		return "", err
	}
	return string(output), nil
}


