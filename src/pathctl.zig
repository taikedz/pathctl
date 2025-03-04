// Learning program
// I'm trying to figure out how to handle returning arbitrary-length arrays of strings from functions
// This involves double-depths of allocation and deallocation, and a faster dive into pointers than I expected

// before any paths, a `!~/.config/pathctlrc` can be specified to configure a configs location
// bin, config, logs, libs, data -- commands to print these locations, from config rc file
// PATHSPEC=~/.PATH
// PATHSPEC_LOCAL=./.PATH
// On queries:
// when `~/.PATH` is not defined, or when config key is not defined, return error
// `-d` flag uses built-in defaults if a value is not found from definition/config
// `-D` flag uses built-in defaults, always.

const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpa = std.heap.GeneralPurposeAllocator(.{}){};

pub fn main() !void {
    const alloc = gpa.allocator();
    defer _ = gpa.deinit();

    const names = try getNames(alloc);
    defer destroyNames(alloc, names);

    for(names.*) |name| {
        try stdout.print("Hello {s}\n", .{name});
    }
}


