load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "second",
    srcs = ["main.go"],
    importpath = "go_monogrepo/packages/second",
    deps = [
        "//packages/second/router:router"
    ],
)