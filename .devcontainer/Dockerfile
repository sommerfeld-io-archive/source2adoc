## This Dockerfile is used to build the development container for the source2adoc project. It
## contains the necessary instructions to create an environment with all the dependencies required
## for development.
##
## The development container is based on the official Java 21 container from Microsoft and comes
## with Apache Maven as well.
##
## == Prerequisites
## This container is intended to be used with Visual Studio Code and the Dev Container plugin.
## So having Visual Studio Code (VSCode) and the Dev Container plugin installed are essential
## prerequisites for this development environment. Docker is also mandatory.

FROM mcr.microsoft.com/vscode/devcontainers/java:21-bullseye
LABEL maintainer="sebastian@sommerfeld.io"

ARG USER=vscode

RUN apt-get update \
    && apt-get install -y --no-install-recommends python3-pip=20.3.4-4+deb11u1 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && pip install --no-cache-dir pre-commit==3.3.3

WORKDIR /opt

ARG MVN_VERSION=3.9.8
RUN wget --progress=dot:giga https://dlcdn.apache.org/maven/maven-3/${MVN_VERSION}/binaries/apache-maven-${MVN_VERSION}-bin.tar.gz \
    && tar xzvf apache-maven-${MVN_VERSION}-bin.tar.gz \
    && echo PATH="/opt/apache-maven-${MVN_VERSION}/bin:\$PATH" >> /home/${USER}/.profile

# ARG KOTLIN_VERSION=2.0.0
# RUN wget --progress=dot:giga https://github.com/JetBrains/kotlin/releases/download/v${KOTLIN_VERSION}/kotlin-compiler-${KOTLIN_VERSION}.zip \
#     && unzip kotlin-compiler-${KOTLIN_VERSION}.zip \
#     && echo PATH="/opt/kotlinc/bin:\$PATH" >> /home/${USER}/.profile \
#     && echo export PATH >> /home/${USER}/.profile

USER $USER

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN echo "alias ls='ls -a --color=auto'" >> /home/${USER}/.bashrc \
    && echo "alias ll='ls -alF --color=auto'" >> /home/${USER}/.bashrc
