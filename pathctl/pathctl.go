package pathctl

import (
	"os"
	"fmt"
	"strings"
	"os/user"
)

func Main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		PrintHelp()
		os.Exit(0)
	}

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
			JustFail("Unsupported subcommand")
		}
	default:
		JustFail("Too many arguments")
	}
}

func pathsAndConfigs() ([]string, *PathConfig) {
	paths, config, err := LoadPathFile()
	if err != nil {
		err.Exit()
	}
	return paths, config
}

func printRequestedSection(section string, config *PathConfig) {
	switch section {
	case "bin":
		sectionValue(section, config.bin, "/usr/local/bin")
	case "lib":
		sectionValue(section, config.lib, "/usr/local/lib")
	case "log":
		sectionValue(section, config.log, "/var/log")
	case "data":
		sectionValue(section, config.data, "/usr/local/share")
	case "config":
		sectionValue(section, config.config, "/etc")
	default:
		JustFail(fmt.Sprintf("Cannot lookup section '%s'", section))
	}
}

func sectionValue(section, value, root_default string) {
	if value == "" {
		if IsRootUser() {
			fmt.Print(root_default)
			return
		}
		JustFail(fmt.Sprintf("%s undefined", section))
	}

	fmt.Print(value)
}

func IsRootUser() bool {
	u, e := user.Current()
	if e != nil {
		JustFail("Fatal - Could not get current user!")
	}
	return u.Uid == "0" // posix only, but this is a posix tool, so OK
}
