@echo off

set dir=%cd%

echo begin building...
del /a /f /q %dir%\dummycli
cd %dir%
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" -o dummycli main.go
echo build finish