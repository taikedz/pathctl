/*
This is a demonstration of how to organise code into other files.

The `util/` folder holds module files themselves, and a `util.rs` file serves as the
unifying layer. We can then import this into the main script using `mod util`.

Note that this file is responsible for declaring the existence of the submodules.
We declare here `pub mod path` at this level, otherwise sibling modules like `display`
cannot find `path`.

*/

pub mod path;
pub mod display;
pub mod append;
