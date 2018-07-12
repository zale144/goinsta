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

	fmt.Println("Deleting", os.Args[2])
	err = media.Items[0].Delete()
	e.CheckErr(err)

	err = media.Sync()
	if err != nil {
		fmt.Println("Deleted!")
	} else {
		fmt.Println("error deleting...")
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
