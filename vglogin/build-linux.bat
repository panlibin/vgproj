@echo off

set dir=%cd%

echo begin building...
del /a /f /q %dir%\vglogin
cd %dir%
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" -o vglogin main.go
echo build finish