package putest

import (
	"testing"

	"github.com/taikedz/pathctl/pathctl"
)

func Test_ParsePathFile(t *testing.T) {
	lines := []string{
		"%bin=/tmp/bin",
		" %config=/tmp/etc",
		"",
		"  # ignored",
		"/tmp/bin",
		"/tmp/alt/bin",
	}

	if paths, pconfig, err := pathctl.ParsePathFile(lines); err != nil {
		t.Errorf("Failed to parse: %v", err)
	} else {
		CheckEqual(t, "/tmp/bin", pconfig.Bin)
		CheckEqual(t, "/tmp/etc", pconfig.Config)
		CheckEqual(t, "", pconfig.Lib)

		CheckEqualArr(t, []string{"/tmp/bin", "/tmp/alt/bin"}, paths)
	}
}
