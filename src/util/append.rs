use std::fs::OpenOptions;
use std::fs::File;
use std::io::Write;
use super::path;


pub fn append(value: &str) {
    let mut fh = open_file(path::get_pathfile_path().as_str());
    let data = format!("{}\n", value);

    // Using the Write module, we can only operate on bytes...
    match fh.write_all(&data.as_bytes()) {
        Err(e) => {
            eprintln!("Failed write: {e}");
            std::process::exit(1);
        }
        Ok(_) => {
            // nothing?
        }
    }
}


fn open_file(filename:&str) -> File {
    // OpenOptions::new() creates a reference. We must capture it in a variable to continue using it in the function
    // we can't use OpenOptions::new().write(true).append(true) as the last doesn't return the temporary itself back out
    let mut opt = OpenOptions::new();
    opt .write(true)
        .append(true);

    match opt.open(filename) {
        Err(e) => {
            println!("Error opening file: {}", e);
            std::process::exit(1);
        }
        Ok(fh) => { return fh; }
    }
}
