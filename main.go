package main

import (
	"github.com/taikedz/pathctl/pathctl"
	"os"
	"fmt"
	"strings"
)

// This go project can be used as a library. These are the exported functions.

/* Obtain the path information from the user's .PATH file

Returns:

* []string : The list of search path locations
* PathConfig : holds the declared preferred locations specified by the user, if any.
* error : an error, if any

*/
func GetUserPaths() ([]string, *pathctl.PathConfig, error) {
	return loadPathFile()
}

/* Determine whether the specified path is in the supplied string list of paths

Both the target path and the individual list paths will be tentatively resolved.
Relative paths are resolved to absolute paths, symlinks are resolved to their original locations.
If the symlink cannot be resolved, the last resolved path is used.

Args:

* target (string) : the path to search for
* paths ([]string) : the list of paths in whcih to search for the target

Returns:

* bool : whether the target was found amongst the list of paths
*/
func PathsHave(target string, paths []string) bool {
	target = pathctl.BestPath(target)
	for _, item := range paths {
		item = pathctl.BestPath(item)
		if target == item {
			return true
		}
	}
	return false
}

// ========================================
// Non-library functions. Place exported functions above.

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		pathctl.PrintHelp()
		os.Exit(0)
	}

	switch len(os.Args) {
	case 1:
		paths, _ := pathsAndConfigs()
		fmt.Print(strings.Join(paths, ":"))
	case 2:
		switch os.Args[1] {
		case "version":
			fmt.Printf("%s v%s\n", pathctl.NAME, pathctl.VERSION)
		default:
			_, config := pathsAndConfigs()
			printRequestedSection(os.Args[1], config)
		}
	case 3:
		paths, _ := pathsAndConfigs()

		switch os.Args[1] {
		case "has":
			if PathsHave(os.Args[2], paths) {
				os.Exit(0)
			} else {
				os.Exit(pathctl.ERR_NO)
			}
		default:
			pathctl.JustFail("Unsupported subcommand")
		}
	default:
		pathctl.JustFail("Too many arguments")
	}
}

func loadPathFile() ([]string, *pathctl.PathConfig, pathctl.ErrorExit) {
	path, err := pathctl.HomePath("~/.PATH")
	if err != nil {
		return nil, nil, pathctl.NewErrorAction(pathctl.ERR_SYSTEM, err.Error() )
	}
	lines, rerr := pathctl.ReadLines(path)
	if rerr != nil {
		return nil, nil, pathctl.NewErrorAction(pathctl.ERR_PATHFILE_FAIL, fmt.Sprintf("Pathctl: %v", rerr) )
	}

	return pathctl.ParsePathFile(lines)
}

func pathsAndConfigs() ([]string, *pathctl.PathConfig) {
	paths, config, err := loadPathFile()
	if err != nil {
		err.Exit()
	}
	return paths, config
}

func printRequestedSection(section string, config *pathctl.PathConfig) {
	switch section {
	case "bin":
		sectionValue(section, config.Bin, "/usr/local/bin")
	case "lib":
		sectionValue(section, config.Lib, "/usr/local/lib")
	case "log":
		sectionValue(section, config.Log, "/var/log")
	case "data":
		sectionValue(section, config.Data, "/usr/local/share")
	case "config":
		sectionValue(section, config.Config, "/etc")
	default:
		pathctl.JustFail(fmt.Sprintf("Section '%s' invalid", section))
	}
}

func sectionValue(section, value, root_default string) {
	if value == "" {
		if pathctl.IsRootUser() {
			fmt.Print(root_default)
			return
		}
		pathctl.JustFail(fmt.Sprintf("Section '%s' not defined in ~/.PATH", section))
	}

	fmt.Print(value)
}