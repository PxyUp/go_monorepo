load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test")
go_binary(
    name = "main_app",
    srcs = ["main.go"],
    importpath = "go_monogrepo/packages/main",
    deps = [
        "//packages/main_app/router:router"
    ],
)
go_test(
    name = "mainTest",
    srcs = ["main_test.go"],
)