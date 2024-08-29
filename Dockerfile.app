## The Dockerfile.app is used to build the application image. Part of the build process is to run
## unit tests and acceptance tests. Each task is executed in a separate stage. The final image
## contains the application binary and the necessary runtime environment and configuration.
##
## [ditaa, target="dockerfile-app"]
## ....
## +-----------------------+
## | Dockerfile.app        |
## |                       |
## | +-------------------+ |
## | | unit test + build | |
## | +-------------------+ |
## | | acceptance test   | |
## | +-------------------+ |
## | | run               | |
## | +-------------------+ |
## +-----------------------+
## ....
##
## == How to use
## Remember to mount volumes, when running an actual command other than `--help`.
##
## [source, bash]
## ....
## docker build -t local/source2adoc:dev -f Dockerfile.app .
## docker run --rm local/source2adoc:dev --help
## ....
##
## To run releases or release candidates, use the respective tags.
##
## [source, bash]
## ....
## docker run --rm sommerfeldio/source2adoc:rc
## docker run --rm sommerfeldio/source2adoc:latest
## ....
##
## @see docker-compose.yml


## The build stage is used to compile the application and run the unit tests. The unit tests are
## executed with the `go test` command as part of the build process.
FROM golang:1.23.0-alpine3.19 AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY testdata /workspaces/source2adoc/testdata
COPY components/app /workspaces/source2adoc/components/app

WORKDIR /workspaces/source2adoc/components/app
RUN go mod download \
    && go mod tidy \
    && go vet ./... \
    && go test -coverprofile=go-code-coverage.out ./... \
    && go build .


## The acceptance-test stage is used to run the acceptance tests against the binary created in the
## build stage.
FROM golang:1.23.0-alpine3.19 AS acceptance-test
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build /workspaces/source2adoc/components/app/source2adoc /workspaces/source2adoc/components/app/source2adoc
COPY testdata /workspaces/source2adoc/testdata
COPY components/test-acceptance /workspaces/source2adoc/components/test-acceptance

WORKDIR /workspaces/source2adoc/components/test-acceptance
RUN go mod download \
    && go mod tidy \
    && go vet ./...

WORKDIR /workspaces/source2adoc/components/test-acceptance/testhelper
RUN go test ./...

WORKDIR /workspaces/source2adoc/components/test-acceptance
RUN go test


## The run stage is used to run the application in a minimal runtime environment. The binary does
## not expect a dedicated runtime environment. A simple Linux environment is sufficient.
FROM alpine:3.20.2 AS run
LABEL maintainer="sebastian@sommerfeld.io"
LABEL org.opencontainers.image.title=source2adoc \
      org.opencontainers.image.description="Streamline the process of generating documentation from inline comments within source code files." \
      org.opencontainers.image.authors="source2adoc open source project" \
      org.opencontainers.image.url="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.documentation="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.source="https://github.com/sommerfeld-io/source2adoc" \
      org.opencontainers.image.vendor="source2adoc open source project" \
      org.opencontainers.image.licenses="MIT License"

COPY config/etc/login.defs /etc/login.defs
RUN chmod og-r /etc/shadow \
    && chmod 0444 /etc/login.defs

ARG USER=source2adoc
RUN adduser -D "$USER"

COPY --from=acceptance-test /workspaces/source2adoc/components/app/source2adoc /usr/bin/source2adoc

RUN chown -hR "$USER:$USER" /usr/bin/source2adoc \
    && chmod 0700 /usr/bin/source2adoc

USER $USER
ENTRYPOINT ["/usr/bin/source2adoc"]
