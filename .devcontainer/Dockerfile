# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.224.3/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.21-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:1.1.6-1-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# install gopls, dlv
RUN echo "Getting development tools"
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

# vscode-go dependencies 
# via: https://github.com/microsoft/vscode-go/blob/master/.travis.yml
RUN echo "Getting dependencies for the golang.go plugin "
RUN go install -v github.com/acroca/go-symbols@latest
RUN go install -v github.com/cweill/gotests/...@latest
RUN go install -v github.com/davidrjenni/reftools/cmd/fillstruct@latest
RUN go install -v github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install -v github.com/mdempsky/gocode@latest
RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v github.com/rogpeppe/godef@latest
RUN go install -v github.com/sqs/goreturns@latest
RUN go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install -v github.com/zmb3/gogetdoc@latest
RUN go install -v golang.org/x/lint/golint@latest
RUN go install -v golang.org/x/tools/cmd/gorename@latest
# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
