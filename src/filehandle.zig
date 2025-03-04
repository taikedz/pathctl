const std = @import("std");

const PATHFILE="~/.PATH";

// Load lines from a file, skip empty lines and lines starting with '#'
pub fn loadPaths(alloc:std.mem.Allocator) std.ArrayList([]const u8) {
    // implementation : can we take just a pointer to `[]u8`, and assign onto it? no allocation?
}

pub fn loadConfigKey(config_key:[]const u8) []const u8 { // alloc?
}

// Whether the lines contain the pathline
fn contains(pathline:[]const u8, lines:std.ArrayList([]const u8) ) bool {
}

fn canonicalPath(path:[]const u8) []const u8 { // allocation?
}

