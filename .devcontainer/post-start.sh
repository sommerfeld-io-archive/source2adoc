#!/bin/bash
## Start Service after the devcontainer is started.
##
## @see .devcontainer/devcontainer.json

echo "[INFO] Starting portainer"
docker compose --file .devcontainer/ops/docker-compose.yml --env-file .devcontainer/ops/.env up -d

# echo "[INFO] Init Go app"
# (
#     cd components/app || exit

#     go mod download
#     go mod tidy
#     go fmt ./...
#     go vet ./...
# )
