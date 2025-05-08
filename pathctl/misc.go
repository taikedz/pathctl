package pathctl

import (
	"os/user"
)

func IsRootUser() bool {
	u, e := user.Current()
	if e != nil {
		JustFail("Fatal - Could not get current user!")
	}
	return u.Uid == "0" // posix only, but this is a posix tool, so OK
}