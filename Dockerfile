# syntax = docker/dockerfile:experimental
# NOTE: Must be run in the context of the repo's root directory

####################################
## (1) Setup the build environment
FROM dockerproxy.com/library/golang:1.21-bullseye AS build-env

# Build the app binary in /app
RUN mkdir /app
WORKDIR /app

ARG TARGET
ARG COMMIT
ARG VERSION

ENV GOPRIVATE=

RUN go env -w GOPROXY=https://goproxy.cn

COPY . .

####################################
## (3) Build the production app binary
FROM build-env as build-production
WORKDIR /app

ARG GOARCH=amd64

# TAGS can be overriden to modify the go build tags (e.g. build without netgo)
# ARG TAGS="relic,netgo"

# Keep Go's build cache between builds.
# https://github.com/golang/go/issues/27719#issuecomment-514747274
# --tags "${TAGS}"
RUN --mount=type=cache,sharing=locked,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=secret,id=git_creds,dst=/root/.netrc \
    CGO_ENABLED=1 GOOS=linux go build  -ldflags "-extldflags -static \
    -X  'main.Version=${VERSION}'" \
    -o ./bin/app ${TARGET} 

RUN chmod a+x /app/bin/app

## (4) Add the statically linked production binary to a distroless image
FROM gcr.dockerproxy.com/distroless/base-debian11 as production

COPY --from=build-production /app/bin/app /bin/app

ENTRYPOINT ["/bin/app"]

####################################
## (3) Build the debug app binary
# --tags "relic,netgo"
FROM build-env as build-debug
WORKDIR /app
ARG GOARCH=amd64
RUN --mount=type=ssh \
    --mount=type=cache,sharing=locked,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=1 GOOS=linux go build  -ldflags "-extldflags -static \
    -X  'main.Version=${VERSION}'" \
    -gcflags="all=-N -l" -o ./bin/app ${TARGET}

RUN chmod a+x /app/bin/app

## (4) Add the statically linked debug binary to a distroless image configured for debugging
FROM dockerproxy.com/library/golang:1.21-bullseye as debug

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY --from=build-debug /app/bin/app /bin/app

ENTRYPOINT ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/bin/app", "--"]
