package ai

import (
	"log"
)

func GenerateResponse(prompt string) (string, error) {
	log.Printf("Generating AI response for: %s", prompt)
	// اینجا می‌تونی به OpenAI یا Llama وصل بشی
	// فعلاً خروجی تستی:
	return "لطفاً شماره اشتراک را بفرمایید", nil
}