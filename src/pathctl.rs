use std::fs::read_to_string;
use std::path::Path;

const PATHFILE_PATH :&str = "~/.PATH";

fn main() {
    display_path();
}

fn display_path() {
    let data:Vec<String> = load_lines(PATHFILE_PATH);
    let resolved_searchpath = data.join(":");

    println!("{resolved_searchpath}");
}

fn load_lines(raw_path:&str) -> Vec<String> {
    read_lines(raw_path)
        .iter() // Explicitly iterate
        .map(|line| line.trim()) // trim() returns a &str , referencing the iterated line
        .filter(|line| is_valid_line(line)) // line is the &str
        .map(String::from) // because the current function owns the object, and we now only have a &str, we create a new object for returning
        .collect::<Vec<String>>() // and finally create a vector, rather than a map . Note the "::" and the doubled '<'
}


fn is_valid_line(line:&str) -> bool {
    return line.len() > 0 && &line[0..1] != "#";
}


fn read_lines(raw_path: &str) -> Vec<String> {
    let path_str = format!("{}", shellexpand::tilde(raw_path));
    let path = Path::new(&path_str);

    if !path.exists() {
        eprintln!("Could not find '{}'", path_str);
        std::process::exit(1);
    }

    // https://doc.rust-lang.org/rust-by-example/std_misc/file/read_lines.html
    read_to_string(path_str) 
        .unwrap()           // panic on possible file-reading errors
        .lines()            // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()          // gather them together into a vector
}
