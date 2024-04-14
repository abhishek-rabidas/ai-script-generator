package views

type TranscriptionResponse struct {
	Language string  `json:"language"`
	Duration float64 `json:"duration"`
	Text     string  `json:"text"`
}
