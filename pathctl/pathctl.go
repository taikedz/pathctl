package pathctl

import (
	"os"
	"fmt"
	"strings"
	"os/user"
)

const NAME string = "Pathctl"
const VERSION string = "0.0.1"

func Main() {
	switch len(os.Args) {
	case 1:
		paths, _, err := LoadPathFile()
		if err != nil {
			err.Exit()
		}
		fmt.Print(strings.Join(paths, ":"))
	case 2:
		if os.Args[1] == "version" {
			fmt.Printf("%s v%s\n", NAME, VERSION)
		} else {
			_, config, err := LoadPathFile()
			if err != nil {
				err.Exit()
			}
			printRequestedSection(os.Args[1], config)
		}
	default:
		print("Oops!\n")
	}
}


func printRequestedSection(section string, config *PathConfig) {
	printed := false

	switch section {
	case "bin":
		printed = sectionValue(config.bin, "/usr/local/bin")
	case "lib":
		printed = sectionValue(config.lib, "/usr/local/lib")
	case "log":
		printed = sectionValue(config.log, "/var/log")
	case "data":
		printed = sectionValue(config.data, "/usr/local/share")
	case "config":
		printed = sectionValue(config.config, "/etc")
	default:
		ErrorAction{ERR_BAD_USER_SECTION, fmt.Sprintf("Cannot lookup section %s", section)}.Exit()
	}

	if !printed {
		ErrorAction{1, fmt.Sprintf("%s undefined", section)}.Exit()
	}
}

func sectionValue(value, defaultval string) bool {
	if value == "" {
		// if root, print defaultval
		u, e := user.Current()
		if e != nil {
			ErrorAction{ERR_SYSTEM, "Fatal - Could not get current user!"}.Exit()
		}
		if u.Uid == "0" { // posix only, but this is a posix tool, so OK
			fmt.Print(defaultval)
			return true
		}
		return false
	}

	fmt.Print(value)
	return true
}