@echo off

set dir=%cd%

echo begin building...
del /a /f /q %dir%\vggame
cd %dir%
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" -o vggame main.go
echo build finish