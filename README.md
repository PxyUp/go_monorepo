# Go-monorepo with Bazel
Hello every this is real-life example of monorepository with Bazel and go-modules

# Setup

1) Install Bazel - https://docs.bazel.build/versions/master/install.html
2) Install Go Modules
```shell script
go mod download
```

# Build
All applications inside packages folder:
1) Main App
2) Second app

## Main App

Build:
```shell script
bazel build //packages/main_app:main_app
```

Bin will be here:

` bazel-bin/packages/main_app/darwin_amd64_stripped/main_app`

Test:
```shell script
bazel test //packages/main_app/...:all
```
## Second App

Build:
```shell script
bazel build //packages/second_app:second_app
```

Bin will be here:

`bazel-bin/packages/second_app/darwin_amd64_stripped/second_app`

Test:
```shell script
bazel test //packages/second_app/...:all
```

# All application

Test:
```shell script
bazel test //packages/...:all
```