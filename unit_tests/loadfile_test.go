package putest

import (
	"testing"

	"github.com/taikedz/pathctl/libpctl"
	"github.com/taikedz/gocheck"
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

	if paths, pconfig, err := libpctl.ParsePathFile(lines); err != nil {
		t.Errorf("Failed to parse: %v", err)
	} else {
		gocheck.Equal(t, "/tmp/bin", pconfig.Bin)
		gocheck.Equal(t, "/tmp/etc", pconfig.Config)
		gocheck.Equal(t, "", pconfig.Lib)

		gocheck.EqualArr(t, []string{"/tmp/bin", "/tmp/alt/bin"}, paths)
	}
}
