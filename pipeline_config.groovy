jte {
    pipeline_template = 'k8s_build.groovy'
}

libraries {
    git {
        repo_url = 'https://github.com/livelace/girie.git'
    }
    go {
        options = 'github.com/livelace/girie/cmd/girie'
    }
    kaniko {
        destination = 'data/girie:latest'
    }
    mattermost
    nexus {
        source = 'girie'
        destination = 'dists-internal/girie/girie-${VERSION}'
    }
    version
}

keywords {
    build_image = 'harbor-core.k8s-2.livelace.ru/dev/gobuild:latest'
}
