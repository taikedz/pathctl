use std::env;
use clap::Parser;


/// Add or show .PATH paths
#[derive(Parser)]
#[command(author, version)]
struct Args {
    /// Base action
    action: String,

    /// Target for base action
    #[arg(default_value_t = String::from(""))]
    target: String,
}


pub struct Action {
    // Because this struct is inside a module, if it is to be used outside,
    //  we also need to declare which properties are also externally accessible
    pub label: String,
    pub target: String,
}


pub fn get_arguments(valid_args:&Vec<&str>) -> Action {
    let action:Action;
    let args : Vec<String> = env::args().collect();

    if args.len() <= 1 {
        eprintln!("Specify an action: {}", String::from(valid_args.join(", ")));
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

