package main

import (
	"vgproj/vgmaster/private"

	"github.com/panlibin/virgo"
)

func main() {
	virgo.Launch(private.NewServer())
}
