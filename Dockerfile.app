## Dockerfile for a Kotlin application that is build using Maven.
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


FROM eclipse-temurin:21-jdk-jammy AS build
LABEL maintainer="sebastian@sommerfeld.io"

ARG MVN_VERSION=3.9.8
WORKDIR /opt
RUN wget --progress=dot:giga https://dlcdn.apache.org/maven/maven-3/${MVN_VERSION}/binaries/apache-maven-${MVN_VERSION}-bin.tar.gz \
    && tar xzvf apache-maven-${MVN_VERSION}-bin.tar.gz \
    && echo PATH="/opt/apache-maven-${MVN_VERSION}/bin:$PATH" >> ~/.profile \
    && echo export PATH >> ~/.profile

COPY components/app /workspace/source2adoc/components/app
WORKDIR /workspace/source2adoc/components/app
RUN ./mvnw clean verify


FROM eclipse-temurin:21-jre-jammy AS run
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build /workspace/source2adoc/components/app/target/source2adoc.jar /opt/source2adoc.jar
ENTRYPOINT ["java", "-jar", "/opt/source2adoc.jar"]
