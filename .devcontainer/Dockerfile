ARG GOLANG_VERSION=1.20.6-bullseye
FROM --platform=${BUILDPLATFORM} golang:${GOLANG_VERSION}
# gopls - go language server (https://pkg.go.dev/golang.org/x/tools/gopls)
RUN go install -v golang.org/x/tools/gopls@latest
# dlv -  debugger for golang (https://github.com/go-delve/delve)
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
# staticcheck - linter for golang (https://staticcheck.io/)
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
# buildifier - bazel build file formatter
go install github.com/bazelbuild/buildtools/buildifier@latest
# gazelle - bazel build file generator 
RUN go install -v github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
# bazelisk - bazel version manager
RUN go install github.com/bazelbuild/bazelisk@latest
