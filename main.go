package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/taikedz/pathctl/libpctl"
)

// This go project can be used as a library. These are the exported functions.

/* Obtain the path information from the user's .PATH file

Returns:

* []string : The list of search path locations
* PathConfig : holds the declared preferred locations specified by the user, if any.
* error : an error, if any

*/
func GetUserPaths() ([]string, *libpctl.PathConfig, error) {
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
	target = libpctl.BestPath(target)
	for _, item := range paths {
		item = libpctl.BestPath(item)
		if target == item {
			return true
		}
	}
	return false
}

// ========================================
// Non-library functions. Place exported functions above.

func justFail(message string) {
	libpctl.NewErrorAction(libpctl.ERR_CMD, fmt.Sprintf("%s. Try 'help' command.", message) ).Exit()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		libpctl.PrintHelp()
		os.Exit(0)
	}

	switch len(os.Args) {
	case 1:
		paths, _ := pathsAndConfigs()
		fmt.Print(strings.Join(paths, ":"))
	case 2:
		switch os.Args[1] {
		case "version":
			fmt.Printf("%s v%s\n", libpctl.NAME, libpctl.VERSION)
		default:
			_, config := pathsAndConfigs()
			printRequestedSection(os.Args[1], config)
		}
	case 3:
		paths, config := pathsAndConfigs()

		switch os.Args[1] {
		case "has":
			if PathsHave(os.Args[2], paths) {
				os.Exit(0)
			} else {
				os.Exit(libpctl.ERR_NO)
			}
		case "install":
			addBinFile(os.Args[2], config.Bin)
		default:
			justFail("Unsupported subcommand")
		}
	default:
		justFail("Too many arguments")
	}
}

func loadPathFile() ([]string, *libpctl.PathConfig, libpctl.ErrorExit) {
	path, err := libpctl.HomePath("~/.PATH")
	if err != nil {
		return nil, nil, libpctl.NewErrorAction(libpctl.ERR_SYSTEM, err.Error() )
	}
	lines, rerr := libpctl.ReadLines(path)
	if rerr != nil {
		return nil, nil, libpctl.NewErrorAction(libpctl.ERR_PATHFILE_FAIL, fmt.Sprintf("Pathctl: %v", rerr) )
	}

	return libpctl.ParsePathFile(lines)
}

func pathsAndConfigs() ([]string, *libpctl.PathConfig) {
	paths, config, err := loadPathFile()
	if err != nil {
		err.Exit()
	}
	return paths, config
}

func printRequestedSection(section string, config *libpctl.PathConfig) { // FIXME: RESOLVE IN PathConfig !!
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
		justFail(fmt.Sprintf("Section '%s' invalid", section))
	}
}

func sectionValue(section, value, root_default string) {
	if value == "" {
		if os.Getuid() == 0 { // POSIX only, and since this is a posix tool, fine as it is.
			fmt.Print(root_default)
			return
		}
		justFail(fmt.Sprintf("Section '%s' not defined in ~/.PATH", section))
	}

	fmt.Print(value)
}

func addBinFile(filepath string, destdir string) {
	if ! libpctl.PathExists(filepath) {
		libpctl.NewErrorAction(
			libpctl.ERR_FAIL_FIND,
			fmt.Sprintf("Could not find file '%s'", filepath),
		).Exit()
	}

	destpath := path.Join(BestPath(destdir), filepath.File(filepath).Name()) // TODO
	libpctl.FileCopy(filepath, destpath)
}
