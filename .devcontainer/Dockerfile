##
# This Dockerfile sets up an environment for Go development within a devcontainer.
# It is based on the link:https://hub.docker.com/r/sommerfeldio/devcontainer-go[sommerfeldio/devcontainer-go:rc] image.
#
# The `USER` argument is set to 'vscode'.
##

FROM sommerfeldio/devcontainer-go:rc
LABEL maintainer="sebastian@sommerfeld.io"

ARG USER=vscode

USER $USER

RUN echo "alias ls='ls -a --color=auto'" >> /home/$USER/.bashrc \
    && echo "alias ll='ls -alF --color=auto'" >> /home/$USER/.bashrc
