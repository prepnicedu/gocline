load("@rules_go//go:def.bzl", "go_binary", "go_library")
load("//bazel/oci:go_image.bzl", "go_image")

go_library(
    name = "bff-svc_lib",
    srcs = ["main.go"],
    importpath = "gocline.com/ws/services/bff-svc",
    visibility = ["//visibility:private"],
    deps = [
        "//gen/go/user/v1:user",
        "//gen/go/user/v1/userv1connect",
        "//pkg/core/config",
        "//pkg/core/grpc-client",
        "//pkg/core/http-server",
        "//services/bff-svc/service",
    ],
)

go_binary(
    name = "bff-svc",
    embed = [":bff-svc_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "image",
    binary = "bff-svc",
)
