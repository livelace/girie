libraries {
    git {
        repo_url = "https://github.com/livelace/girie.git"
    }
    go {
        options = "github.com/livelace/girie/cmd/girie"
    }
    k8s {
        image = "harbor-core.k8s-2.livelace.ru/dev/gobuild:latest"
    }
    kaniko {
        destination = "data/girie:latest"
    }
    mattermost
    nexus {
        source = "girie"
        destination = 'dists-internal/girie/girie-${VERSION}'
    }
    version
}
