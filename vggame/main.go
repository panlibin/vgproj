package main

import (
	"vgproj/vggame/private"

	"github.com/panlibin/virgo"
)

func main() {
	virgo.Launch(private.NewServer())
}
