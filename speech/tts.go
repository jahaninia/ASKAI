package speech

import (
	"log"
)

// func SynthesizeSpeech(text string) (string, error) {
// 	log.Printf("Synthesizing speech for: %s", text)
// 	// اینجا باید به سرویس TTS مثل Google یا Coqui وصل بشی
// 	// فعلاً خروجی تستی:
// 	return "response.wav", nil
// }
// use Coqui 
func SynthesizeSpeech(text string) (string, error) {
	cmd := exec.Command("tts", "--text", text, "--model_name", "tts_models/fa/fast_pitch", "--out_path", "response.wav")
	err := cmd.Run()
	if err != nil {
		log.Printf("TTS error: %v", err)
		return "", err
	}
	return "response.wav", nil
}
