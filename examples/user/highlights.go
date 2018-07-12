// +build ignore

package user

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)
	user.Sync()

	hlts, err := user.Highlights()
	e.CheckErr(err)

	for _, h := range hlts {
		// getting images URL
		for _, item := range h.Items {
			if len(item.Images.Versions) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Versions[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}

	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
