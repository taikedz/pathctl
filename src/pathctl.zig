// Learning program
// I'm trying to figure out how to handle returning arbitrary-length arrays of strings from functions
// This involves double-depths of allocation and deallocation, and a faster dive into pointers than I expected

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


