##
# This Dockerfile serves as the central piece for the source2adoc application, representing both
# both local development and the creation of publicly available images on DockerHub.
#
# * Named `build`, this stage is responsible for unit-testing and building the Go application.
# * Named `run`, this stage is responsible for running the Go application in a minimal Alpine Linux.
#
# As the only Dockerfile relevant to the source2adoc application, it represents the sole artifact
# published from the application. This means that any Docker images produced by the application,
# whether for development or production, stem directly from this Dockerfile.
#
# In essence, this Dockerfile encapsulates the entire lifecycle of the source2adoc application's
# containerized environment, from development to deployment, making it a crucial and central
# component of the project.
#
# [source, bash]
# ....
# docker build --build-arg VERSION=dev --build-arg COMMIT_SHA=dev -t local/source2adoc:dev .
# ....
#
# @see docker-compose.yml
##


##
# Build stage
##
FROM golang:1.22.3-alpine3.19 AS build
LABEL maintainer="sebastian@sommerfeld.io"

ARG VERSION=UNSPECIFIED
ARG COMMIT_SHA=UNSPECIFIED

COPY /components/app /components/app
COPY /components/testdata /components/testdata
WORKDIR /components/app

RUN go mod download \
    && go mod tidy \
    && echo "${VERSION}" > internal/metadata/VERSION \
    && echo "${COMMIT_SHA}" > internal/metadata/COMMIT_SHA \
    && test ./... \
    && go build .


##
# Run stage
##
FROM alpine:3.19.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"
LABEL version=$VERSION
LABEL org.opencontainers.image.title=source2adoc \
      org.opencontainers.image.description="Convert inline documentation into AsciiDoc files, tailored for seamless integration with Antora. " \
      org.opencontainers.image.authors="source2adoc open source project" \
      org.opencontainers.image.url="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.documentation="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.source="https://github.com/sommerfeld-io/source2adoc" \
      org.opencontainers.image.vendor="source2adoc open source project" \
      org.opencontainers.image.licenses="... tbd ... "

ARG USER=source2adoc
RUN adduser -D "$USER"

COPY --from=build /components/app/source2adoc /usr/bin/source2adoc

USER $USER
ENTRYPOINT ["/usr/bin/source2adoc"]
