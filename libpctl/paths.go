package libpctl

import (
	"os"
	"errors"
	"strings"
	"os/user"
	"path/filepath"
)

func HomePath(target string) (string, ErrorExit) {
	curuser, err := user.Current()
	if err != nil {
		return "", ErrorAction{ERR_SYSTEM, err.Error()}
	}
	return filepath.Join(curuser.HomeDir, target[2:]), nil
}

func BestPath(target string) string {
	// Try to resolve the most precise path
	// If a path cannot be absolute-ed, or resolved, just return the last
	//   known path string
	if strings.Index(target, "~/") == 0 {
		path, err := HomePath(target)
		if err != nil {
			err.Exit()
		}
		target = path
	} else {
		path, err := filepath.Abs(target)
		if err != nil {
			// Typically will not happen, even on non-existent paths...??
			// Just return the entry
			return path
		}
		target = path
	}

	if ! pathExists(target) {
		// The path location listed in .PATH does not (yet) exist
		// We do not need it to exist, only to resolve it as far as possible
		return target
	}

	end_target, err := filepath.EvalSymlinks(target)
	if err != nil {
		// Typically if the symlink cannot be resolved due to a destination not existing
		// we do not care...
		return target
	}
	return end_target
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