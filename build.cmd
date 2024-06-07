@echo off

set TARGET=PowerModeTray.exe

@go build -o %TARGET% -ldflags="-H windowsgui"
@tools\go-winres patch --no-backup --in res/winres.json %TARGET%

@pause
