const std = @import("std");

pub fn build(b:*std.Build) void {
    const exe = b.addExecutable(.{
        .name = "pathctl",
        .root_source_file = b.path("src/pathctl.zig"),
        .target = b.standardTargetOptions(.{}),
        .optimize = b.standardOptimizeOption(.{}),
    });

    // FIXME - see if we even need this?
    // Find the dependency named "zg" as declared in build.zg.zon
    const dep_zg = b.dependency("zg", .{});

    // Get the `CaseData` module from the `zg` dependency
    //   and register as importable module "zg_CaseData"
    // In project code, use @import("zg_CaseData")
    exe.root_module.addImport("zg_CaseData", dep_zg.module("CaseData"));

    const dep_lr = b.dependency("lr", .{});
    exe.root_module.addImport("linereader", dep_lr.module("zig-linereader"));

    b.installArtifact(exe);
}

