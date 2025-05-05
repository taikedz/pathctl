package pathctl

import (
	"os"
	"fmt"
	"os/user"
)

const NAME string = "Pathctl"
const VERSION string = "0.0.1"

func Main() {
	switch len(os.Args) {
	case 0:
		paths, _, err := loadPathfile()
		if err != nil {
			err.Exit()
		}
		print(strings.Join(paths, ":"))
	case 1:
		if os.Args[1] == "version" {
			fmt.Printf("%s v%s", NAME, VERSION)
		} else {
			_, config, err = loadPathfile()
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
		printed = sectionValue(config.bin)
	case "libs":
		printed = sectionValue(config.libs)
	case "log":
		printed = sectionValue(config.log)
	case "data":
		printed = sectionValue(config.data)
	case "config":
		printed = sectionValue(config.config)
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
			print(defaultval)
			return true
		}
		return false
	}

	print(value)
	return true
}