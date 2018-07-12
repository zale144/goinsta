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

	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)
	err = media.Items[0].Comments.Enable()
	e.CheckErr(err)

	media.Sync()
	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
