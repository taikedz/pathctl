/*
This is a demonstration of how to organise code into other files.

The `util/` folder holds _submodules_ to the `util.rs` module.

We can import this using `mod util` in a file sibling to `util.rs`

Note that this file is responsible for declaring the _existence_ of the submodules.

If this module does not declare its submodules, then the submodules will not
 find eachother even as siblings.
*/

pub mod path;
pub mod args;
pub mod display;
pub mod append;

// we can specify functions/structs etc at this level for the `util` module
// the rest as declared above are from _submodules_ of `util`.
