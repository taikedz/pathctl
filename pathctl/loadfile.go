package pathctl

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type PathConfig struct {
	bin string
	logs string
	config string
	data string
}

func loadPathFile() string[], *PathConfig, ErrorExit {
	lines, err := readlines("~/.PATH")
	if err != nil {
		return nil, nil, ErrorAction{ERR_PATHFILE_FAIL, fmt.Sprintf("ERROR: %v", err)}
	}

	return parsePathFile(lines)
}

func parsePathfile(lines []string) string[], *PathConfig, ErrorExit {

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
		if strings.Index(path, ":") {
			return nil, nil, ErrorAction{ERR_INVALID_PATH, fmt.Sprintf("Path cannot contain ':' , got '%s'", path)}
		}
		if section != nil {
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
			case "logs":
				config.logs = path
			case "lib":
				config.lib = path
			default:
				return nil, nil, ErrorAction{ERR_BAD_SECTION, fmt.Sprintf("Unknown section: %s", section)}
			}
		}
		else {
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

func extractTokens(line string) string, string {
	if line.Index("%") == 0 && line.Index("=") > 1 {
		tokens := strings.SplitN(line, "=", 2)
		return tokens[0][1:], tokens[1]
	} else {
		return nil, line
	}
}