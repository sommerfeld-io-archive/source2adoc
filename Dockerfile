FROM eclipse-temurin:22.0.1_8-jdk-jammy AS builder
LABEL maintainer="sebastian@sommerfeld.io"

COPY /components/app /components/app
WORKDIR /components/app
ARG MVN_OPTS=--no-transfer-progress -Dstyle.color=always
RUN ./mvnw "$MVN_OPTS" dependency:go-offline
RUN ls -alF
RUN ./mvnw "$MVN_OPTS" clean verify


FROM eclipse-temurin:22.0.1_8-jre-jammy
LABEL maintainer="sebastian@sommerfeld.io"
LABEL org.opencontainers.image.source="https://github.com/sommerfeld-io/source2adoc"
LABEL org.opencontainers.image.description "Streamline the process of generating AsciiDoc documentation from inline comments within source code files."

ARG USER=source2adoc
RUN addgroup ${USER} && adduser --ingroup ${USER} --disabled-password ${USER}
USER ${USER}

COPY --from=builder /components/app/target/source2adoc.jar /opt/source2adoc/source2adoc.jar
ENTRYPOINT ["java", "-jar", "/opt/source2adoc/source2adoc.jar"]
