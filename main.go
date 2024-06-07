package main

import (
	_ "embed"
	"power-mode-tray/libs"
	"syscall"

	"github.com/getlantern/systray"
)

//go:embed res/icon.ico
var appIcon []byte

const appName = "PowerModeTray v1.0.0.0"

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(appIcon)
	systray.SetTitle(appName)
	systray.SetTooltip(appName)

	current, _ := libs.PowerGetEffectiveOverlayScheme()
	menuBestPerformance := systray.AddMenuItemCheckbox("최고 성능", "최고 성능", *current == libs.PowerModeBestPerformance)
	menuBetterPerformance := systray.AddMenuItemCheckbox("균형 잡힌", "균형 잡힌", *current == libs.PowerModeBetterPerformance)
	menuBetterBattery := systray.AddMenuItemCheckbox("최고의 전원 효율성", "최고의 전원 효율성", *current == libs.PowerModeBetterBattery)

	systray.AddSeparator()

	menuClose := systray.AddMenuItem("끝내기", "트레이를 종료합니다.")

	menuOptions := map[syscall.GUID]systray.MenuItem{
		libs.PowerModeBestPerformance:   *menuBestPerformance,
		libs.PowerModeBetterPerformance: *menuBetterPerformance,
		libs.PowerModeBetterBattery:     *menuBetterBattery,
	}

	go func() {
		for {
			select {
			case <-menuBestPerformance.ClickedCh:
				libs.PowerSetActiveOverlayScheme(&libs.PowerModeBestPerformance)
			case <-menuBetterPerformance.ClickedCh:
				libs.PowerSetActiveOverlayScheme(&libs.PowerModeBetterPerformance)
			case <-menuBetterBattery.ClickedCh:
				libs.PowerSetActiveOverlayScheme(&libs.PowerModeBetterBattery)
			case <-menuClose.ClickedCh:
				systray.Quit()
			}

			currentScheme, _ := libs.PowerGetEffectiveOverlayScheme()
			for guid, menuItem := range menuOptions {
				if guid == *currentScheme {
					menuItem.Check()
				} else {
					menuItem.Uncheck()
				}
			}
		}
	}()
}

func onExit() {
	// NOOP
}
