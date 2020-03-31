package main

import (
	"vgproj/vgrecharge/private"

	"github.com/panlibin/virgo"
)

func main() {
	virgo.Launch(private.NewServer())
}
