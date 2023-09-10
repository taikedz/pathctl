// We use shellexpand in this file, but don't declare its `use`
//  This is because it is declared in Cargo as an auto-scoped item ?

pub const PATHFILE_PATH:&str = "~/.PATH";


pub fn get_pathfile_path() -> String {
    format!("{}", shellexpand::tilde(PATHFILE_PATH))
}
