use std::fs::File;
use std::path::Path;

const PATHFILE_PATH :&str = "~/.PATH";

fn main() {
    let path_str = format!("{}", shellexpand::tilde(PATHFILE_PATH));
    let path = Path::new(&path_str);

    if(!path.exists()) {
        eprintln!("Could not find '{}'", path_str);
        std::process::exit(1);
    }

    let file = match File::open(&path) {
        Err(e) => panic!("File open error: {path_str}: {e}"),
        Ok(file) => file,
    };
}
