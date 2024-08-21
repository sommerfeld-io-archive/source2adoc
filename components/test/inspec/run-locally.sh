#!/bin/bash
## (Temporary) File to test InSpec profiles before turning them into a GitHub Actions workflow.
##
## This script does not accept any arguments.

(
    cd /workspaces/source2adoc || exit

    echo "[INFO] Building images ================================================================="
    docker build -t local/source2adoc:dev -f Dockerfile.app .
    docker build -t local/source2adoc-docs:dev -f Dockerfile.docs .

    readonly INSPEC_PROFILE_PATH="/workspaces/source2adoc/components/test/inspec"

    readonly IMAGES=(
        "source2adoc"
        "source2adoc-docs"
    )

    for image in "${IMAGES[@]}"
    do
        echo "[INFO] Starting container in background ============================================"
        container=$(docker run -d --entrypoint tail "local/$image:dev" -f /dev/null)
        
        echo "[INFO] Run InSpec tests against image =============================================="
        docker run --rm \
            --volume /var/run/docker.sock:/var/run/docker.sock \
            --volume "$(pwd):$(pwd)" \
            --workdir "$(pwd)" \
            chef/inspec:5.22.55 exec "$INSPEC_PROFILE_PATH/$image" --target "docker://$container" --chef-license=accept
        
        echo "[INFO] Stopping container =========================================================="
        docker stop --time 0 "$container"

        echo "[INFO] Removing container =========================================================="
        docker rm "$container"
    done
)
