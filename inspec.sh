#!/bin/bash
## Temporary file to test InSpec profiles before turning them into a GitHub Actions workflow.


docker build -t local/source2adoc-docs:dev -f Dockerfile.docs --no-cache .
container=$(docker run -d --entrypoint tail local/source2adoc-docs:dev -f /dev/null)
export container

(
    cd /tmp || exit
    
    rm -rf apache-baseline
    git clone https://github.com/dev-sec/apache-baseline

    readonly exclude="/^((?!os-14).)*$/"

    docker run --rm \
        --volume /var/run/docker.sock:/var/run/docker.sock \
        --volume "$(pwd):$(pwd)" \
        --workdir "$(pwd)" \
        chef/inspec:5.22.55 exec apache-baseline --target "docker://$container" --controls "$exclude" --chef-license=accept
    
    docker stop --time 0 "$container"
    docker rm "$container"
)
