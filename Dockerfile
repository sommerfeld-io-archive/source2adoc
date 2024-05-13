FROM golang:1.22.3-alpine3.19 AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY /components/app /components/app
WORKDIR /components/app

RUN ls -alF \
    && go mod download \
    && go mod tidy \
    && go test ./... \
    && go build .


FROM alpine:3.19.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"

ARG USER=source2adoc
RUN adduser -D $USER

COPY --from=build /components/app/source2adoc /usr/bin/source2adoc

USER $USER
ENTRYPOINT ["/usr/bin/source2adoc"]
