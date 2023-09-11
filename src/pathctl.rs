// Since this is the main script, `util` and `version` can live sidecar to `pathctl.rs`
//  If they lived in another file, say `mods.rs`, the declarations would cause these to
//  need to live in `mods/util.rs` and `mods/version.rs`
mod util;
mod version;

use util::*;

fn main() {
    // Advanced: specify which ones have which lengths required
    let valid_args : Vec<&str> = vec!["load", "add", "version"];

    let action = args::get_arguments(&valid_args);

    if action.label == "load" {
        display::display_path();
    }
    else if action.label == "version" {
        println!("{}", version::VERSION);
    }
    else if action.label == "add" {
        append::append(&action.target);
    }
    else {
        eprintln!("Invalid action: {}", action.label);
    }
}


