libraries {
    appimage {
        source = "girie"
        destination = 'girie-${VERSION}.appimage'
    }
    dependency_check
    dependency_track {
        project = "girie"
        version = "master"
    }
    git {
        repo_url = "https://github.com/livelace/girie.git"
    }
    go {
        options = "github.com/livelace/girie/cmd/girie"
    }
    harbor_replicate {
        policy = "girie"
    }
    harbor_scan {
        artifact = "data/girie:latest"
        severity = "medium"
    }
    k8s_build {
        image = "harbor-core.k8s-2.livelace.ru/dev/gobuild:latest"
        privileged = true
    }
    kaniko {
        destination = "data/girie:latest"
    }
    mattermost
    nexus {
        source = 'girie-${VERSION}.appimage'
        destination = 'dists-internal/girie/girie-${VERSION}'
    }
    sonarqube
    version
}
