# pathctl

Utility to manage and query bin and install paths for shell scripts and makefiles:

* use `export PATH="$(pathctl)"` (or your shell's equivalent) in `~/.profile` or `~/.bashrc` or such, to populate `PATH`
* use `pathctl bin` in a Makefile or other install script, to get the user's preferred default bin install path
    * same for other paths of interest like `lib`, `config`, `logs`, `data`
* use a command to add to `PATH` instead of appending to `.profile` or `.bashrc` (which is duplicate-prone)

## Motivation

Pathctl aims to provide a single utility via which to manage PATH entries - no need to tell users to manually edit their shell rc files, or try to auto-determine whether to append a new export to those files.

Pathctl aims to provide a query tool that any install script can invoke to discover the user's preferred default location, instead of guessing an unsuitable default.

For example if you are developing an installable application: some users like their configurations in `~/.config`, some like it in `~/.local/config`, and some use `~/.local/etc`. Let the user specify this in a well-known file, and honour it. Avoid producing your own `~/.thingrc` - please!

## Querying the pathfile

When writing an installer, use `pathctl`/`~/.PATH` to resolve locations.

If you are using a shell script, makefile, or the language supports easily captuing subprocesse output, you can determine the install binary and config directories by calling the `pathctl` command:

```sh
# Bash script:

BIN_DIR="$(pathctl get bin)"
CONFIG_DIR="$(pathctl get config)"
```

If you are using an expressive language, you can query the file directly - for example in python:

```python
import os

def loadConfigs(*config_keys) -> List[str]:
    """ Load ~/.PATH config, return values of corresponding keys, in-order.

    If a config key is not defined, corresponding value will be None
    """
    with open(os.path.expanduser("~/.PATH")) as fh:
        lines = [ line.trim() for line in fh.readlines() if line.startswith("%") ]
        configs = dict( [line[1:].split("=", 1) for line in lines[::-1]] )

    return [configs.get(k) for k in config_keys]


bin_dir, config_dir = loadConfigs("bin", "config")

```

## Example pathfile content

```
# Definitions
%bin=~/.local/bin
%config=~/.local/etc
%data=~/.local/data

# Paths
/bin
/usr/bin
/usr/local/bin
~/.local/bin

# Path appended by another installer ...
~/.cargo/bin
```

## Building this project

TBC

## License

The code is conveyed to you under the terms of the GNU _Lesser_ General Public License. It remains open-source, free-software itself, but does not require any project including it to be open source or free-software.

## Spec

I have set myself out a set of requirements to keep me focused:

Path File:

* The pathfile at `~/.PATH` contains user's paths, one per line. Lines may be empty. Lines starting with '#' are comments and are to be skipped
* A config key can be declared before any path lines, with a line beginning `%` followed by a "keyname=path" string. Where several declarations exist, the _first_ occurrence is used.
* Append to the path file via usual means - e.g. `echo ~/.local/bin >> ~/.PATH`

Query:

* `pathctl` may be run without arguments, to print a path-notation of paths, each separated by colon '`:`' . This is the "default action."
* the pathfile will assume precedence for later entries - therefore the resulting path string will list paths in _reverse_ order.
* `pathctl` will not include the same canonical path twice
* `pathctl` resolves user expansion (`~` , `~user`) but _NOT_ environment variables
* `pathctl get` may receive a config key to look up in the path file. If not found, return status 10
* `pathctl default` returns a built-in default from `~/.local/<config key>` if the config key is not defined, otherwise returns the defined key

Info:

* `pathctl info` will print version and copyright notice
* `pathctl version` will print the version string for pathctl alone

