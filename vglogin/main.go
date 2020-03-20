package main

import (
	"vgproj/vglogin/private"

	"github.com/panlibin/virgo"
)

func main() {
	virgo.Launch(private.NewServer())
}
