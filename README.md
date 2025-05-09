# pathctl

Utility to manage and query bin and install paths:

* use `export PATH="$(pathctl):$PATH"` (or your shell's equivalent) in `~/.profile` to populate `PATH`
* use `pathctl bin` in a Makefile or other install script, to get the user's preferred default bin install path
    * for example `bindir="$(pathctl bin)" || exit 1; cp build/bin/myprogram "$bindir"`
    * same for other paths of interest like `lib`, `config`, `log`, `data`

## Motivation

The problem: different developers favour different locations for placing files, and whilst there are general conventions, there are indeed several. Binaries placed in arbitrary folders directly under the home folder, configurations placed in `~/.*rc` files, some configs placed directly in home, some under `~/.config`, some like go just creating a `~/go/bin` and `~/go/pkg` ... The home folder gets littered, and it is never clear where files are going to end up.

Pathctl aims to provide an example utility/technique of a single utility via which to manage PATH entries and common locations - no need to tell users to manually edit their shell rc files, or try to auto-determine whether to append a new `export PATH=...` to those files.

Pathctl aims to provide a query tool that any install script can invoke to discover the user's preferred default location, instead of providing an unsuitable default like `/usr/local/bin` or `~/.local/bin`.

The plain `pathctl` command prints a `PATH` search path string from all locations known to Pathctl.

## Building this project

You need the go toolchain to build this utility.

Run `./build.sh`, a `bin/pathctl` will be built. Alternatively run

```sh
go build -o bin/pathctl main.go
```

## As a library

For go projects, you can use this as a library. Import `github.com/taikedz/pathctl` and use one of the two helper functions:

* `GetUserPaths() ([]string, *libpctl.PathConfig, error)` - to load the plain `.PATH` data
* `PathsHave(target string, paths []string) bool` - to check for path existence, using resolved paths

Remember to `go mod tidy` before first run after import.

## Installation

Create a `~/.PATH` file with the following contents (adjust paths as needed, ensure the preferred path exists):

```
# Define a preferred bin location
%bin=~/.local/bin

# Specify bin locations for export.
~/.local/bin

# Add any others from your shell rc file and remove/comment them out from there
```

Edit your shell rc file and add `export PATH="$(pathctl):$PATH"`

Run `./install-example.sh` to perform the install of the command.

## License

The code is conveyed to you under the terms of the GNU _Lesser_ General Public License. It remains open-source itself, but does not require any project including it to be open source.
