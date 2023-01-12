package install

import (
	"os"
	"path"

	"github.com/whiterabb17/shaman/package/util"
	"golang.org/x/sys/windows/registry"
)

const (
	runKey = "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
)

// TryRegistryInstall attempts to create a Run entry under the appropriate root key.
func TryRegistryInstall() error {
	var root registry.Key
	if util.RunningAsAdmin() {
		root = registry.LOCAL_MACHINE
	} else {
		root = registry.CURRENT_USER
	}
	run, err := registry.OpenKey(root, runKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer run.Close()
	return run.SetStringValue(util.Registry, path.Join(Info.Base, util.Binary))
}

// UninstallRegistry removes Run entries created by a registry install.
// Call with nil for a full uninstall.
func UninstallRegistry(root interface{}) error {
	if root == nil {
		if util.RunningAsAdmin() {
			err := UninstallRegistry(registry.LOCAL_MACHINE)
			if err != nil {
				return err
			}
		}
		root = registry.CURRENT_USER
	}

	run, err := registry.OpenKey(root.(registry.Key), runKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = run.DeleteValue(util.Registry)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
