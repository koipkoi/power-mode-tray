package libs

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func GetModuleFileName() string {
	buf := make([]uint16, 1024)
	ret, err := windows.GetModuleFileName(0, &buf[0], 1024)
	if ret == 0 || err != nil {
		return ""
	}
	return syscall.UTF16ToString(buf)
}
