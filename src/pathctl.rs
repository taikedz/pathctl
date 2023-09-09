mod util;
mod version;


fn main() {
    // Advanced: specify which ones have which lengths required
    let valid_args : Vec<&str> = vec!["load", "add", "version"];

    let action = util::args::get_arguments(&valid_args);

    if action.label == "load" {
        util::display::display_path();
    }
    else if action.label == "version" {
        println!("{}", version::VERSION);
    }
    else if action.label == "add" {
        util::append::append(&action.target);
    }
    else {
        eprintln!("Invalid action: {}", action.label);
    }
}


