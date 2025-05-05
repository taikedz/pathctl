package pathctl

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"os/user"
)

const NAME string = "Pathctl"
const VERSION string = "0.0.2"

func Main() {
	switch len(os.Args) {
	case 1:
		paths, _ := pathsAndConfigs()
		fmt.Print(strings.Join(paths, ":"))
	case 2:
		switch os.Args[1] {
		case "version":
			fmt.Printf("%s v%s\n", NAME, VERSION)
		default:
			_, config := pathsAndConfigs()
			printRequestedSection(os.Args[1], config)
		}
	case 3:
		paths, _ := pathsAndConfigs()

		switch os.Args[1] {
		case "has":
			if pathsHave(os.Args[2], paths) {
				os.Exit(0)
			} else {
				os.Exit(ERR_NO)
			}
		default:
			ErrorAction{ERR_CMD, "Unsupported subcommand"}.Exit()
		}
	default:
		print("Oops!\n")
	}
}

func pathExists(target string) bool {
	_, err := os.Stat(target)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		// we cannot know whether the item exists
		// https://stackoverflow.com/a/12518877/2703818
		ErrorAction{ERR_SYSTEM, err.Error()}.Exit()
	}
	return false
}

func pathsAndConfigs() ([]string, *PathConfig) {
	paths, config, err := LoadPathFile()
	if err != nil {
		err.Exit()
	}
	return paths, config
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
		if IsRootUser() {
			fmt.Print(defaultval)
			return true
		}
		return false
	}

	fmt.Print(value)
	return true
}

func IsRootUser() bool {
	u, e := user.Current()
	if e != nil {
		ErrorAction{ERR_SYSTEM, "Fatal - Could not get current user!"}.Exit()
	}
	return u.Uid == "0" // posix only, but this is a posix tool, so OK
}
