mod util;
mod version;

use std::env;

struct Action {
    label: String,
    target: String,
}


fn main() {
    let action = get_arguments();

    if action.label == "load" {
        util::display::display_path();
    }
    else if action.label == "version" {
        println!("{}", version::VERSION);
    }
    else if action.label == "add" {
        util::append::append(&action.target);
    }
    // more options
}


fn get_arguments() -> Action {
    let action:Action;
    let args : Vec<String> = env::args().collect();

    if args.len() <= 1 {
        eprintln!("Specify an action: load, add, version");
        std::process::exit(1);
    }
    else if args.len() <= 2 {
        action = Action{ label: String::from(&args[1]), target: String::from("") };
    }
    else if args.len() <= 3 {
        action = Action{ label: String::from(&args[1]), target: String::from(&args[2]) };
    }
    else {
        eprintln!("Incorrect arguments count");
        std::process::exit(2);
    }

    // Validate on enum

    return action;
}
