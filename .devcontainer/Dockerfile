# This Dockerfile sets up an environment for Kotlin development within a devcontainer.
# It is based on the link:https://hub.docker.com/r/sommerfeldio/devcontainer-ubuntu[sommerfeldio/devcontainer-ubuntu:rc] image.
#
# * The Java version of choice is Amazon Corretto 22.
# * The Maven version of choice is 3.9.6
#
# The `USER` argument is set to 'vscode'.
#
# === Prerequisites
#
# Having Visual Studio Code (VSCode) and the Dev Container plugin installed are
# essential prerequisites for this development environment.

FROM sommerfeldio/devcontainer-ubuntu:rc
LABEL maintainer="sebastian@sommerfeld.io"

ARG USER=vscode

# Install Java
ARG JAVA_VERSION=21
RUN mkdir /usr/lib/jvm \
    && curl -L -o /tmp/corretto.tar.gz https://corretto.aws/downloads/latest/amazon-corretto-${JAVA_VERSION}-x64-linux-jdk.tar.gz \
    && tar -xzf /tmp/corretto.tar.gz -C /usr/lib/jvm \
    && rm /tmp/corretto.tar.gz

# Install Maven
ARG MAVEN_VERSION=3.9.6
ARG MAVEN_BASE_URL=https://apache.osuosl.org/maven/maven-3/${MAVEN_VERSION}/binaries

RUN curl -fsSL -o /tmp/apache-maven.tar.gz ${MAVEN_BASE_URL}/apache-maven-${MAVEN_VERSION}-bin.tar.gz \
    && tar -xzf /tmp/apache-maven.tar.gz -C /usr/share \
    && rm /tmp/apache-maven.tar.gz \
    && ln -s /usr/share/apache-maven-${MAVEN_VERSION} /usr/share/maven \
    && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn

USER $USER

RUN echo "alias ls='ls -a --color=auto'" >> /home/$USER/.bashrc \
    && echo "alias ll='ls -alF --color=auto'" >> /home/$USER/.bashrc

ENV JAVA_HOME=/usr/lib/jvm/amazon-corretto-21.0.3.9.1-linux-x64
ENV PATH=$PATH:$JAVA_HOME/bin

RUN mkdir -p /home/$USER/.m2
ENV MAVEN_HOME=/usr/share/maven
ENV PATH=$PATH:$MAVEN_HOME/bin
