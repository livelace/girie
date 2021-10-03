def APP_NAME = "girie"
def APP_REPO = "https://github.com/livelace/girie.git"
def APP_VERSION = env.VERSION + '-${GIT_COMMIT_SHORT}'
def IMAGE_TAG = env.VERSION == "master" ? "latest" : env.VERSION

libraries {
    appimage {
        source = "${APP_NAME}"
        destination = "${APP_NAME}-${APP_VERSION}.appimage"
    }
    dependency_check
    dependency_track {
        project = "${APP_NAME}"
        version = "env.VERSION"
    }
    git {
        repo_url = "${APP_REPO}"
        //repo_branch = env.VERSION
    }
    go {
        options = "github.com/livelace/girie/cmd/girie"
    }
    harbor_replicate {
        policy = "${APP_NAME}"
    }
    k8s_build {
        image = "harbor-core.k8s-2.livelace.ru/dev/gobuild:latest"
        privileged = true
    }
    kaniko {
        destination = "data/${APP_NAME}:${IMAGE_TAG}"
    }
    mattermost
    nexus {
        source = "${APP_NAME}-${APP_VERSION}.appimage"
        destination = "dists-internal/${APP_NAME}/${APP_NAME}-${APP_VERSION}.appimage"
    }
    sonarqube
}
