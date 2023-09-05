use std::fs::read_to_string;
use std::path::Path;

const PATHFILE_PATH :&str = "~/.PATH";

fn main() {
    let data:Vec<String> = read_lines(PATHFILE_PATH);
    for line in data.iter() {
        println!("{line}");
    }
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
