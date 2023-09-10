use super::path;

use std::fs::read_to_string;
use std::path::Path;

pub fn display_path() {
    let path_str = path::get_pathfile_path();
    let data:Vec<String> = load_lines(&path_str.as_str());
    let resolved_searchpath = data.join(":");

    println!("{resolved_searchpath}");
}


fn load_lines(path_str:&str) -> Vec<String> {
    read_lines(path_str)
        .iter() // Explicitly iterate

        // trim() returns a &str , referencing the iterated line
        // which cannot be returned (owned and remaining in the function)
        // so we produce a new String
        .map(|line| String::from(line.trim()))

        .filter(|line| is_valid_line(&line))

        // and finally create a vector, from the map .
        // Note the "::" and the doubled '<' for generic-in-generic
        .collect::<Vec<String>>()
}


fn is_valid_line(line:&String) -> bool {
    return line.len() > 0 && &line[0..1] != "#";
}


fn read_lines(path_str: &str) -> Vec<String> {
    let path = Path::new(&path_str);

    if !path.exists() {
        eprintln!("Could not find '{}'", path_str);
        std::process::exit(1);
    }

    // https://doc.rust-lang.org/rust-by-example/std_misc/file/read_lines.html
    match read_to_string(path_str) {
        Err(e) => {
            eprintln!("Error reading file '{}': {}", path_str, e);
            std::process::exit(1);
        }
        Ok(data) => {
            data.lines()            // split the string into an iterator of string slices
                .map(String::from)  // make each slice into an owned string
                .collect()          // gather them together into a vector
        }
    }
}
