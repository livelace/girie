def APP_NAME = "girie"
def APP_REPO = "https://github.com/livelace/girie.git"
def APP_VERSION = env.VERSION + '-${GIT_COMMIT_SHORT}'
def IMAGE_TAG = env.VERSION == "master" ? "latest" : env.VERSION

libraries {
    appimage {
        source = "${APP_NAME}"
        destination = "${APP_NAME}-${APP_VERSION}.appimage"
    }
    //dependency_check
    //dependency_track {
    //    project = "${APP_NAME}"
    //    version = env.VERSION
    //}
    dracut
    git {
        repo_url = "${APP_REPO}"
        repo_branch = env.VERSION
    }
    go {
        options = "-ldflags=\"-X 'github.com/livelace/girie/pkg/girie/core.APP_VERSION=${APP_VERSION}'\" " +
            "-tags dynamic github.com/livelace/girie/cmd/girie"
        test = false
    }
    //harbor_replicate {
    //    policy = "${APP_NAME}"
    //}
    k8s_build {
        image = "registry.livelace.ru/dev/girie:${IMAGE_TAG}"
    }
    kaniko {
        destination = "infra/${APP_NAME}:${IMAGE_TAG}"
    }
    mattermost
    nexus {
        source = "${APP_NAME}-${APP_VERSION}.appimage"
        destination = "raw/${APP_NAME}/${APP_NAME}-${APP_VERSION}.appimage"
    }
    //sonarqube
}
