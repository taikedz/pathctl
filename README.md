# pathctl

Utility to manage and query bin and install paths:

* use `export PATH="$(pathctl)"` (or your shell's equivalent) in `~/.profile` to populate `PATH`
* use `pathctl bin` in a Makefile or other install script, to get the user's preferred default bin install path
    * same for other paths of interest like `lib`, `config`, `logs`, `data`
* use a command to add to `PATH` instead of appending to `.profile` or `.bashrc` (which is duplicate-prone)

## Motivation

Path-ctl aims to provide a single utility via which to manage PATH entries - no need to tell users to manually edit their shell rc files, or try to auto-determine whether to append a new export to those files.

Path-ctl aims to provide a query tool that any install script can invoke to discover the user's preferred default location, instead of providing an unsuitable default like `/usr/local/bin` or `~/.local/bin`.

## Building this project

TBC

## License

The code is conveyed to you under the terms of the GNU _Lesser_ General Public License. It remains open-source itself, but does not require any project including it to be open source.

## Requirements

I had set myself out a set of requirements to keep me focused:

* the program binary is called `pathctl`
* `~/.PATH` file contains user's paths, one per line. Lines may be empty. Lines starting with '#' are comments and are to be skipped
* command must take the argument `load` to print a path-notation of paths, each separated by colon '`:`'
* command must take the arguments `add DIR_PATH` which will add the path to the end of `~/.PATH`
* command may take the argument `version` which will print the version string alone
* any other argument form causes help to be printed, and exits with error code
* any error should cause print out of error message, without printing stack traces or other debugging noise

