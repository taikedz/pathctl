# pathctl

Based on the quick python scirpt I wrote as an idea: <https://gist.github.com/taikedz/4cda9e4650ad10fe827c1224816e0269>

This is predominantly a learning project. I am at the beginning of my journey in rust, so this code is likely to be pretty janky, and potentially over-experimental.

## Building this project

You probably want to check the basics in [_The Book_](https://doc.rust-lang.org/book/) but in summary

* With rust and cargo installed (`sudo apt install cargo` usually will do it)
* run `cargo build`
* execute with `./target/debug/pathctl`

This _should_ be a fully finalized copy of the program, as set out by the specification below.

In the `Cargo.toml` I explicitly chose to strip the debug symbols form the build - this reduces the file's size from 12MB to 400K . An impressive difference...

## Next steps


Look into unit testing libraries and practices.


## Requirements

I had set myself out a set of requirements to keep me focused:

* the program binary is called `pathctl`
* `~/.PATH` file contains user's paths, one per line. Lines may be empty. Lines starting with '#' are comments and are to be skipped
* command must take the argument `load` to print a path-notation of paths, each separated by colon '`:`'
* command must take the arguments `add DIR_PATH` which will add the path to the end of `~/.PATH`
* command may take the argument `version` which will print the version string alone
* any other argument form causes help to be printed, and exits with error code
* any error should cause print out of error message, without printing stack traces or other debugging noise

