// +build ignore

package media

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media, err := inst.GetMedia(os.Args[2])
	e.CheckErr(err)

	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)
	media.Items[0].Like()

	media.Sync()
	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
