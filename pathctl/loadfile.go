package pathctl

import (
	"fmt"
	"strings"
)

type PathConfig struct {
	Bin string
	Log string
	Config string
	Data string
	Lib string
}

func ParsePathFile(lines []string) ([]string, *PathConfig, ErrorExit) {

	config := PathConfig{}
	var paths []string

	in_head := true

	for lineno, line := range lines {
		line = strings.Trim(line, " \t")
		if line == "" || strings.Index(line, "#") == 0 {
			continue
		}

		section, path := extractTokens(line)
		if path == "" {
			return nil, nil, ErrorAction{ERR_NOPATH, fmt.Sprintf("Unspecified path on line %d of ~/.PATH", lineno)}
		}
		if strings.Index(path, ":") >= 0 {
			return nil, nil, ErrorAction{ERR_INVALID_PATH, fmt.Sprintf("Path cannot contain ':' , got '%s'", path)}
		}

		path = BestPath(path)

		if section != "" {
			if !in_head {
				return nil, nil, ErrorAction{ERR_HEADS_BEYOND_HEADS, fmt.Sprintf("Line %d: Cannot specify a section beyond top of .PATH file: '%s'", lineno, line)}
			}

			switch(section) {
			case "bin":
				config.Bin = path
			case "config":
				config.Config = path
			case "data":
				config.Data = path
			case "log":
				config.Log = path
			case "lib":
				config.Lib = path
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

func extractTokens(line string) (string, string) {
	if strings.Index(line,"%") == 0 && strings.Index(line,"=") > 1 {
		tokens := strings.SplitN(line, "=", 2)
		return tokens[0][1:], tokens[1]
	} else {
		return "", line
	}
}