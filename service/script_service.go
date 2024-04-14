package service

import (
	"ai-script-generator/config"
	"ai-script-generator/model"
	"ai-script-generator/views"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"io/ioutil"
	"net/http"
	"os"
)

type ScriptService struct {
}

func (s ScriptService) GenerateNewScript(req views.NewScriptRequest) (string, error) {

	account := model.Account{}

	err := config.DB.Preload("Profile").Where("uid = ?", req.AccountId).Find(&account).Error

	if err != nil {
		log.Error(err)
		return "", errors.New("could not find the account")
	}

	prompt := fmt.Sprintf("I speak %d words per second, Generate a voice over script with timeline for a %s video, duration is %d seconds, topic is %s",
		account.Profile.Wps, req.Platform, req.Duration, req.Topic)

	res, err := getChatCompletionResult(prompt)

	if err != nil {
		log.Error(err)
		return "", err
	}

	return res, nil
}

func getChatCompletionResult(prompt string) (string, error) {
	messages := make([]message, 0)

	messages = append(messages, message{Role: "user", Content: prompt})

	body := request{Model: "gpt-3.5-turbo", Messages: messages}

	bodyBytes, _ := json.Marshal(body)

	r, _ := http.NewRequest("POST", os.Getenv("CHAT_COMPLETION_API_URL"), bytes.NewReader(bodyBytes))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_AI_KEY"))
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", errors.New("could not generate script")
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", errors.New("could not generate script")
	}

	output := response{}

	err = json.Unmarshal(responseBody, &output)
	if err != nil {
		return "", errors.New("could not generate script")
	}

	if len(output.Choices) == 0 {
		return "", errors.New("could not generate script")
	}

	return output.Choices[0].Message.Content, nil
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type response struct {
	Id                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
