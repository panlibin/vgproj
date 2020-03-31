@echo off

set dir=%cd%

echo begin building...
del /a /f /q %dir%\vgmaster.exe
SET CGO_ENABLED=0
go build -ldflags "-s -w" -o vgmaster.exe main.go

echo build finish
