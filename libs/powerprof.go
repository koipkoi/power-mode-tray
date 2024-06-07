package libs

import (
	"syscall"
	"unsafe"
)

var (
	PowerModeBestPerformance = syscall.GUID{
		Data1: 0xded574b5,
		Data2: 0x45a0,
		Data3: 0x4f42,
		Data4: [8]byte{0x87, 0x37, 0x46, 0x34, 0x5c, 0x09, 0xc2, 0x38},
	}

	PowerModeBetterPerformance = syscall.GUID{
		Data1: 0x00000000,
		Data2: 0x0000,
		Data3: 0x0000,
		Data4: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	PowerModeBetterBattery = syscall.GUID{
		Data1: 0x961cc777,
		Data2: 0x2547,
		Data3: 0x4f9d,
		Data4: [8]byte{0x81, 0x74, 0x7d, 0x86, 0x18, 0x1b, 0x8a, 0x7a},
	}

	PowerModeBatterySaver = syscall.GUID{
		Data1: 0x3af9B8d9,
		Data2: 0x7c97,
		Data3: 0x431d,
		Data4: [8]byte{0xad, 0x78, 0x34, 0xa8, 0xbf, 0xea, 0x43, 0x9f},
	}
)

var (
	powrprof                      *syscall.LazyDLL
	setActiveOverlaySchemeProc    *syscall.LazyProc
	getEffectiveOverlaySchemeProc *syscall.LazyProc
)

func init() {
	powrprof = syscall.NewLazyDLL("powrprof.dll")
	if powrprof == nil {
		return
	}

	setActiveOverlaySchemeProc = powrprof.NewProc("PowerSetActiveOverlayScheme")
	if setActiveOverlaySchemeProc == nil {
		return
	}

	getEffectiveOverlaySchemeProc = powrprof.NewProc("PowerGetEffectiveOverlayScheme")
	if getEffectiveOverlaySchemeProc == nil {
		return
	}
}

func PowerSetActiveOverlayScheme(guid *syscall.GUID) uint32 {
	ret, _, _ := setActiveOverlaySchemeProc.Call(uintptr(unsafe.Pointer(guid)))
	return uint32(ret)
}

func PowerGetEffectiveOverlayScheme() (*syscall.GUID, uint32) {
	var guid syscall.GUID
	ret, _, _ := getEffectiveOverlaySchemeProc.Call(uintptr(unsafe.Pointer(&guid)))
	return &guid, uint32(ret)
}
