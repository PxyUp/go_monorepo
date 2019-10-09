load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "main_app",
    srcs = ["main.go"],
    importpath = "go_monogrepo/packages/main",
    deps = [
        "//packages/main/router:router"
    ],
)