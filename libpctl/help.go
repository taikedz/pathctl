package pathctl

import (
	"os"
	"text/template"
)

const AUTHOR string = "Tai Kedzierski"
const NAME string = "Pathctl"
const VERSION string = "0.0.3"
const BINNAME string = "pathctl"

const helptext string = `{{.name}} v{{.version}} (C) {{.author}}

Helper tool to set and inform the user's PATH entries and preferred locations.

Commands
========

{{.cmd}}
	Print PATH string, joining paths from ~/.PATH

{{.cmd}} bin
	Print the user's preferred path for installing executables

{{.cmd}} config
	Print the user's preferred path for installing configuration files and folders

{{.cmd}} lib
	Print the user's preferred path for installing lib files

{{.cmd}} data
	Print the user's preferred path for installing data files

{{.cmd}} log
	Print the user's preferred path for writing log files


.PATH file
==========

Add a ~/.PATH file to your home directory. Define preferred locations by adding their section declarations to the top of the file,
followed by executables' paths one per line, like

	# Blank lines and comments are supported

	# Section declarations:
	%bin=~/.local/bin
	%config=~/.local/etc

	# Paths list:
	~/.local/bin
	~/go/bin

This will cause '{{.cmd}}' to print '/home/user/.local/bin:/home/user/go/bin',
and '{{.cmd}} config' will print '/home/user/.local/etc'

Preference sections must be defined before paths.
`

func PrintHelp() {
	data := make(map[string]string)
	data["version"] = VERSION
	data["name"] = NAME
	data["cmd"] = BINNAME
	data["author"] = AUTHOR

	templ, err := template.New("Pathctl Help").Parse(helptext)
	AbortIfError(err, ERR_SYSTEM)

	err = templ.Execute(os.Stdout, data)
	AbortIfError(err, ERR_SYSTEM)
}