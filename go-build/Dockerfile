ARG GO="1.21"

FROM golang:alpine

# https://github.com/boumenot/gocover-cobertura
ARG GOCOVERCOBERTURA="1.1.0"
# https://github.com/mvdan/gofumpt
ARG GOFUMPT="0.1.1"
# https://github.com/golang/tools
ARG GOIMPORTS="0.1.3"
# https://github.com/jstemmer/go-junit-report
ARG GOJUNITREPORT="0.9.1"
# https://github.com/golangci/golangci-lint/releases
ARG GOLANGCI_LINT="1.41.1"
# https://github.com/mcubik/goverreport
ARG GOVERREPORT="1.0.0"
# https://github.com/hexdigest/gowrap/releases
ARG GOWRAP="1.2.1"
# https://github.com/santhosh-tekuri/jsonschema

LABEL name="golang-build"
LABEL maintainer="Developer"

ARG BIN_DIR="/usr/local/bin"

# Install from apk
RUN echo -e "\e[32m==> go version\e[0m" && \
    go version && \
    echo -e "\e[32m==> alpine linux\e[0m" && \
    cat /etc/os-release

RUN \
    #
    # Go install
    #
    #   Formatting tools
    #
    echo -e "\e[32m==> gofumpt ${GOFUMPT}\e[0m" && \
    go install mvdan.cc/gofumpt@v${GOFUMPT} && \
    gofumpt -version && \
    #
    echo -e "\e[32m==> goimports ${GOIMPORTS}\e[0m" && \
    go install golang.org/x/tools/cmd/goimports@v${GOIMPORTS} && \
    which goimports && \
    #
    #   Testing / coverage tools
    #
    echo -e "\e[32m==> gocover-cobertura ${GOCOVERCOBERTURA}\e[0m" && \
    go install github.com/boumenot/gocover-cobertura@v${GOCOVERCOBERTURA} && \
    gocover-cobertura -h && \
    #
    echo -e "\e[32m==> go-junit-report ${GOJUNITREPORT}\e[0m" && \
    go install github.com/jstemmer/go-junit-report@v${GOJUNITREPORT} && \
    go-junit-report -h && \
    #
    echo -e "\e[32m==> goverreport ${GOVERREPORT}\e[0m" && \
    go install github.com/mcubik/goverreport@v${GOVERREPORT} && \
    goverreport -h \


ENV CGO_ENABLED=0

WORKDIR /project
ENTRYPOINT ["sh"]
