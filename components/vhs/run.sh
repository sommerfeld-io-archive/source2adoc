#!/bin/bash
## This script is used to run the VHS tapes in the components/vhs directory to capture bash
## commands and render an animated gif for the terminal session.
##
## The script installs VHS and all dependencies.


readonly SRC_DIR="src"
readonly OUTPUT_DIR="docs"


## Install VHS and all dependencies
function install() {
    echo "[INFO] Install dependencies"
    sudo apt-get update
    sudo apt-get install -y --no-install-recommends ffmpeg=7:4.3.7-0+deb11u1 \
    sudo apt-get install -y --no-install-recommends chromium-sandbox=120.0.6099.224-1~deb11u1 \
    sudo apt-get clean \
    sudo rm -f /var/lib/apt/lists/*

    if [ ! -f /usr/bin/ttyd ]; then
        echo "[INFO] Install ttyd"
        curl -L https://github.com/tsl0922/ttyd/releases/download/1.7.7/ttyd.x86_64 -o /usr/bin/ttyd
        sudo chmod +x /usr/bin/ttyd
    fi

    (
        echo "[INFO] Install VHS"
        go install github.com/charmbracelet/vhs@latest
    )
}


## Prepare testdata for the screen recording.
function prepare() {
    echo "[INFO] Prepare testdata"
    mkdir "$SRC_DIR"
    cp -a ../../testdata/common/good/script.sh "$SRC_DIR"
    cp -a ../../testdata/common/good/docker/Dockerfile "$SRC_DIR"
    cp -a ../../testdata/common/good/yaml "$SRC_DIR"
    cp -a ../../testdata/common/good/Makefile "$SRC_DIR"
    cp -a ../../testdata/common/good/Vagrantfile "$SRC_DIR"
}

## Cleanup after the screen recording.
function cleanup() {
    echo "[INFO] Cleanup"
    rm -rf "$SRC_DIR"
    rm -rf "$OUTPUT_DIR"
}


## Run all VHS tapes
function run() {
    prepare

    for tape in *.tape
    do
        echo "[INFO] Running tape $tape"
        vhs "$tape"
    done

    cleanup
}

install
run
