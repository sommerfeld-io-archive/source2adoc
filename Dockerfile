## This Dockerfile serves as the central piece for the source2adoc application, representing both
## both local development and the creation of publicly available images on DockerHub.
##
## As the only Dockerfile relevant to the source2adoc application, it represents the sole artifact
## published from the application. This means that any Docker image produced by the application,
## whether for development or production, stems directly from this Dockerfile.
##
## In essence, this Dockerfile encapsulates the entire lifecycle of the source2adoc application's
## containerized environment, from development to deployment, making it a crucial and central
## component of the project.
##
## [source, bash]
## ....
## docker build -t local/source2adoc:dev .
## docker run --rm --volume "$(pwd):$(pwd)" --workdir "$(pwd)" local/source2adoc:dev list --lang=yml
## ....
##
## @see docker-compose.yml



## Named `build`, this stage is responsible for unit-testing and building the Go application.
FROM golang:1.22.4-alpine3.19 AS build
LABEL maintainer="sebastian@sommerfeld.io"

ARG COMMIT_SHA=UNSPECIFIED

COPY /components/app /components/app
COPY /metadata/NEXT_VERSION /components/app/internal/metadata/VERSION
WORKDIR /components/app

RUN go mod download \
    && go mod tidy \
    && test ./... \
    && go build .


## Named `run`, this stage is responsible for running the Go application in a minimal Alpine Linux.
FROM alpine:3.20.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"
LABEL org.opencontainers.image.title=source2adoc \
      org.opencontainers.image.description="Convert inline documentation into AsciiDoc files, tailored for seamless integration with Antora. " \
      org.opencontainers.image.authors="source2adoc open source project" \
      org.opencontainers.image.url="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.documentation="https://source2adoc.sommerfeld.io" \
      org.opencontainers.image.source="https://github.com/sommerfeld-io/source2adoc" \
      org.opencontainers.image.vendor="source2adoc open source project" \
      org.opencontainers.image.licenses="GNU GENERAL PUBLIC LICENSE (Version 3, 29 June 2007)"

ARG USER=source2adoc
RUN adduser -D "$USER"

COPY --from=build /components/app/source2adoc /usr/bin/source2adoc

USER $USER
ENTRYPOINT ["/usr/bin/source2adoc"]
