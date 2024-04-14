package views

type NewScriptRequest struct {
	AccountId string `json:"accountId"`
	Topic     string `json:"topic"`
	Platform  string `json:"platform"`
	Duration  int    `json:"duration"`
}
