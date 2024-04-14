package service

import (
	"ai-script-generator/config"
	"ai-script-generator/model"
	"ai-script-generator/views"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/charmbracelet/log"
	"github.com/oklog/ulid/v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) CreateUser(req views.UserCreateRequest) (views.UserResponse, error) {
	profile := model.Profile{
		Uid:   ulid.Make().String(),
		Wps:   0,
		Name:  req.Name,
		Email: req.Email,
	}

	account := model.Account{
		Uid:     ulid.Make().String(),
		Profile: profile,
	}

	err := config.DB.Create(&account).Error

	if err != nil {
		log.Error(err)
		return views.UserResponse{}, err
	}

	return views.NewUserResponse(account), nil
}

func (u UserService) AnalyzeVoice(filename string, accountId string) error {
	file, err := os.Open(filename)

	if err != nil {
		log.Error(err)
		return errors.New("failed to open audio file")
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", file.Name())
	io.Copy(part, file)

	writer.WriteField("model", "whisper-1")
	writer.WriteField("response_format", "verbose_json")

	writer.Close()

	r, _ := http.NewRequest("POST", os.Getenv("TRANSCRIPTION_API_URL"), body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_AI_KEY"))
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Error(err)
		return errors.New("error in reading transcription")
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return errors.New("error in reading transcription")
	}

	response := views.TranscriptionResponse{}

	err = json.Unmarshal(responseBody, &response)

	log.Info("Response: ", ">>", response)

	account := model.Account{}

	err = config.DB.Preload("Profile").Where("uid = ?", accountId).Find(&account).Error

	if err != nil {
		log.Error(err)
		return err
	}

	totalWords := len(strings.Split(response.Text, " "))

	account.Profile.Wps = totalWords / int(response.Duration)
	err = config.DB.Save(&account.Profile).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
