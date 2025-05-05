package pathctl

import (
	"os"
	"fmt"
)

type ErrorExit interface {
	Error() string
	Exit()
}

type ErrorAction struct {
	code int
	info string
}

func (e ErrorAction) Error() string {
	return e.info
}

func (e ErrorAction) Exit() {
	fmt.Printf("%s\n", e.info)
	os.Exit(e.code)
}

const ERR_SYSTEM int = 1

const ERR_PATHFILE_FAIL int = 10
const ERR_BAD_SECTION int = 11
const ERR_NOPATH int = 12
const ERR_BAD_USER_SECTION int = 13
const ERR_HEADS_BEYOND_HEADS int = 14
const ERR_INVALID_PATH int = 15
const ERR_FAIL_FIND int = 16