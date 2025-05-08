package putest

import (
	"testing"

	"github.com/taikedz/pathctl/pathctl"
)

func Test_ParsePathFile(t *testing.T) {
	lines := []string{
		"%bin=/tmp/bin",
		"%config=/tmp/etc",
		"/tmp/bin",
		"/tmp/alt/bin",
	}
	paths, pconfig, err := pathctl.ParsePathFile(lines)

	CheckEqual(t, "/tmp/bin", pconfig.bin)
	CheckEqual(t, "/tmp/etc", pconfig.config)
	CheckEqual(t, "", pconfig.lib)
}