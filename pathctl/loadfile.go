package pathctl

import (
	"fmt"
	"strings"
	"os/user"
	"path/filepath"
)

type PathConfig struct {
	bin string
	log string
	config string
	data string
	lib string
}

func LoadPathFile() ([]string, *PathConfig, ErrorExit) {
	curuser, err := user.Current()
	if err != nil {
		return nil, nil, ErrorAction{ERR_SYSTEM, err.Error()}
	}
	path := filepath.Join(curuser.HomeDir, ".PATH")
	lines, err := ReadLines(path)
	if err != nil {
		return nil, nil, ErrorAction{ERR_PATHFILE_FAIL, fmt.Sprintf("ERROR: %v", err)}
	}

	return parsePathFile(lines)
}

func parsePathFile(lines []string) ([]string, *PathConfig, ErrorExit) {

	config := PathConfig{}
	var paths []string

	in_head := true

	for lineno, line := range lines {
		if isBlankOrCommentLine(line) {
			continue
		}

		section, path := extractTokens(line)
		if path == "" {
			return nil, nil, ErrorAction{ERR_NOPATH, fmt.Sprintf("Unspecified path on line %d of ~/.PATH", lineno)}
		}
		if strings.Index(path, ":") >= 0 {
			return nil, nil, ErrorAction{ERR_INVALID_PATH, fmt.Sprintf("Path cannot contain ':' , got '%s'", path)}
		}
		if section != "" {
			if !in_head {
				return nil, nil, ErrorAction{ERR_HEADS_BEYOND_HEADS, "Cannot specify a section beyond top of .PATH file"}
			}

			switch(section) {
			case "bin":
				config.bin = path
			case "config":
				config.config = path
			case "data":
				config.data = path
			case "log":
				config.log = path
			case "lib":
				config.lib = path
			default:
				return nil, nil, ErrorAction{ERR_BAD_SECTION, fmt.Sprintf("Unknown section: %s", section)}
			}
		} else {
			in_head = false
			paths = append(paths, path)
		}
	}

	return paths, &config, nil
}

func isBlankOrCommentLine(line string) bool {
	line = strings.Trim(line, " \t")
	return line == "" || strings.Index(line, "#") == 0
}

func extractTokens(line string) (string, string) {
	if strings.Index(line,"%") == 0 && strings.Index(line,"=") > 1 {
		tokens := strings.SplitN(line, "=", 2)
		return tokens[0][1:], tokens[1]
	} else {
		return "", line
	}
}