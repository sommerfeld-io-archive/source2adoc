## Dockerfile for a Go application.
##
## == How to use
## Build the Docker image using the following command: `docker build --no-cache  -t local/source2adoc:dev -f Dockerfile.app .`
##
## Use `docker run --rm local/source2adoc:dev` to run the application based on the local build.
##
## Use `docker run --rm sommerfeldio/source2adoc:rc` to run the most recent release candidate from
## DockerHub.
##
## @see docker-compose.yml


FROM golang:1.22.5-alpine3.19 AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY /components/app /workspaces/source2adoc/components/app
WORKDIR /workspaces/source2adoc/components/app

RUN pwd && ls -alF \
    && go mod download \
    && go mod tidy \
    && go test -coverprofile=go-code-coverage.out ./... \
    && go build .


FROM alpine:3.20.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"
LABEL org.opencontainers.image.title=source2adoc \
      org.opencontainers.image.description="Streamline the process of generating documentation from inline comments within source code files." \
      org.opencontainers.image.authors="source2adoc open source project" \
      org.opencontainers.image.url="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.documentation="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.source="https://github.com/sommerfeld-io/source2adoc" \
      org.opencontainers.image.vendor="source2adoc open source project" \
      org.opencontainers.image.licenses="MIT License"

ARG USER=source2adoc
RUN adduser -D "$USER"

COPY --from=build /workspaces/source2adoc/components/app/source2adoc /usr/bin/source2adoc

USER $USER
ENTRYPOINT ["/usr/bin/source2adoc"]
