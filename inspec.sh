#!/bin/bash
## Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt
## ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo
## dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor
## sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor
## invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et
## justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est.


# docker run --name source2adoc -d --entrypoint tail sommerfeldio/source2adoc:rc -f /dev/null
# docker run --name source2adoc -d --entrypoint tail sommerfeldio/source2adoc-docs:rc -f /dev/null

docker build -t local/source2adoc:dev -f Dockerfile.app --no-cache .
container=$(docker run -d --entrypoint tail local/source2adoc:dev -f /dev/null)
export container

# docker build -t local/source2adoc-docs:dev -f Dockerfile.docs --no-cache .
# docker run --name source2adoc_docs -d --entrypoint tail local/source2adoc-docs:dev -f /dev/null

(
    cd /tmp || exit
    
    rm -rf linux-baseline
    git clone https://github.com/dev-sec/linux-baseline

    readonly exclude="/^((?!os-14).)*$/"

    docker run --rm \
        --volume /var/run/docker.sock:/var/run/docker.sock \
        --volume /workspaces/source2adoc/.inspec.yml:/workspaces/source2adoc/.inspec.yml:ro \
        --volume "$(pwd):$(pwd)" \
        --workdir "$(pwd)" \
        chef/inspec:5.22.55 exec linux-baseline --target "docker://$container" --controls "$exclude" --chef-license=accept
    
    docker stop --time 0 "$container"
    docker rm "$container"

    # docker run --rm \
    #     --volume /var/run/docker.sock:/var/run/docker.sock \
    #     --volume /workspaces/source2adoc/.inspec.yml:/workspaces/source2adoc/.inspec.yml:ro \
    #     --volume "$(pwd):$(pwd)" \
    #     --workdir "$(pwd)" \
    #     chef/inspec:5.22.55 exec linux-baseline --target docker://source2adoc_docs --controls "$exclude" --chef-license=accept
    
    # docker stop --time 0 source2adoc_docs
    # docker rm source2adoc_docs
)
