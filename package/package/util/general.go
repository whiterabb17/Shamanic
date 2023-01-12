package util

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//RemoveDuplicates returns a new slice with duplicate elements removed.
func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, i := range slice {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			list = append(list, i)
		}
	}

	return list
}

//CopyFile attempts to copy a file from src to dst.
//Attributes are not preserved.
//Environment variables in paths are not supported.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

//RunPowershell executes a PowerShell command.
//Returns an error if the command fails or PowerShell cannot run.
func RunPowershell(command string) error {
	return RunPowershellInternal(command, false)
}

func RunPowershellInternal(command string, mScope bool) error {
	cmd := exec.Command("powershell", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if mScope {
		cmd.Dir = os.ExpandEnv("$temp")
	}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Command returned an error: " + string(out))
	}
	return err
}

//CheckSingle checks for a lock file and exits if one is found.
func CheckSingle() {
	err := os.Remove(os.Args[0] + ":" + Lock)
	if err != nil && !os.IsNotExist(err) {
		log.Println("An instance is already running")
		os.Exit(0)
	}
	os.OpenFile(os.Args[0]+":"+Lock, os.O_CREATE|os.O_EXCL, 0600)
}
