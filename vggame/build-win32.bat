@echo off

set dir=%cd%

echo begin building...
del /a /f /q %dir%\vggame.exe
SET CGO_ENABLED=0
go build -ldflags "-s -w" -o vggame.exe main.go

echo build finish
