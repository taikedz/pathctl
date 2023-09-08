mod util;

const PATHFILE_PATH :&str = "~/.PATH";

fn main() {
    util::display::display_path(PATHFILE_PATH);
}

