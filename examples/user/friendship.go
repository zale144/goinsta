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

	// At this context you can use:
	// user.FriendShip()
	// user.Sync(true)
	err = user.Sync(true)
	e.CheckErr(err)

	fmt.Println("Following:", user.Friendship.Following)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
