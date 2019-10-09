load("@io_bazel_rules_go//go:def.bzl", "go_library")
go_library(
    name = "router",
    srcs = ["router.go"],
    importpath = "go_monogrepo/packages/main_app/router",
    visibility = ["//visibility:public"],
    deps = [
         "//packages/shared/handlers/health:health",
          "@com_github_gin_gonic_gin//:go_default_library"
    ]
)