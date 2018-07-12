// +build ignore

package media

import (
	"fmt"
	"os"
	"time"

	e "gopkg.in/ahmdrz/goinsta.v1/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media, err := inst.GetMedia(os.Args[2])
	e.CheckErr(err)

	fmt.Printf("Comments: %d\n", media.Items[0].CommentCount)
	err = media.Items[0].Comments.DelMine(0)
	e.CheckErr(err)

	fmt.Println("wait 5 seconds...")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println()

	media.Sync()
	fmt.Printf("After calling: Comments: %d\n", media.Items[0].CommentCount)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
