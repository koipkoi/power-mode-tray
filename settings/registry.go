package settings

import (
	"fmt"
	"power-mode-tray/libs"

	"golang.org/x/sys/windows/registry"
)

const registryKey = `Software\Microsoft\Windows\CurrentVersion\Run`
const autoStartKey = "PowerModeTray"

func AutoStartEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}
	defer k.Close()

	s, _, err := k.GetStringValue(autoStartKey)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	currentAppPath := libs.GetModuleFileName()
	return s == currentAppPath
}

func AutoStartEnable() {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.CREATE_SUB_KEY|registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer k.Close()

	currentAppPath := libs.GetModuleFileName()
	err2 := k.SetStringValue(autoStartKey, currentAppPath)
	if err2 != nil {
		fmt.Printf("%s\n", err2)
		return
	}
}

func AutoStartDisable() {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.CREATE_SUB_KEY|registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer k.Close()

	err2 := k.DeleteValue(autoStartKey)
	if err2 != nil {
		fmt.Printf("%s", err2)
		return
	}
}
