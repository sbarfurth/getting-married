ARG GO_VERSION="1.24"

FROM mcr.microsoft.com/devcontainers/go:${GO_VERSION}

ARG PROTOC_VERSION="30.2"

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
    rm protoc-${PROTOC_VERSION}-linux-x86_64.zip
