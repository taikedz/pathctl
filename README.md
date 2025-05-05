# pathctl

Utility to manage and query bin and install paths:

* use `export PATH="$(pathctl)"` (or your shell's equivalent) in `~/.profile` to populate `PATH`
* use `pathctl bin` in a Makefile or other install script, to get the user's preferred default bin install path
    * same for other paths of interest like `lib`, `config`, `log`, `data`
* use a command to add to `PATH` instead of appending to `.profile` or `.bashrc` (which is duplicate-prone)

## Motivation

Path-ctl aims to provide a single utility via which to manage PATH entries - no need to tell users to manually edit their shell rc files, or try to auto-determine whether to append a new export to those files.

Path-ctl aims to provide a query tool that any install script can invoke to discover the user's preferred default location, instead of providing an unsuitable default like `/usr/local/bin` or `~/.local/bin`.

The plain `pathctl` command prints a `PATH` search path string from all locations known to Path-ctl.

## Building this project

You need the go toolchain to build this utility.

Run `./build.sh`, a `bin/pathctl` will be built.

## License

The code is conveyed to you under the terms of the GNU _Lesser_ General Public License. It remains open-source itself, but does not require any project including it to be open source.

## Requirements

These are the project requirements.

* the program binary is called `pathctl`
* `~/.PATH` file contains user's paths, one per line. Lines may be empty. Lines starting with '#' are comments and are to be skipped
* command run without arguments prints a path-notation of paths, each separated by colon '`:`'
* command may take the argument `version` which will print the pathctl binary's version string alone
* command may take the agument `help` which will print the help and exit with zero status code
* command may take a path-like string; if the path is in the `.PATH` paths, exit zero, else exit non-zero
* subcommands `bin`, `config`, `lib`, `log`, `data` described under subcommands section below
* any error should cause print out of error message, without printing stack traces or other debugging noise

Preference subcommands as follows. If a section is not defined, nothing is printed, and the command exits with a non-zero status code. If the section is not defined, and `pathctl` was run as root, a default is provided unless overridden by an entry in `/root/.PATH`

* `pathctl config` - prints the user's preferred config directory; root default is `/etc`
* `pathctl log` - prints the user's preferred logs directory; root default is `/var/logs`
* `pathctl bin` - prints the user's preferred `bin` directory ; root default is `/usr/local/bin`
* `pathctl lib` - prints the user's preferred libs directory; root default is `/usr/local/lib`
* `pathctl data` - prints the user's preferred data directory; root default is `/usr/local/share`